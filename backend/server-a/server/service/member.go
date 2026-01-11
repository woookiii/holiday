package service

import (
	"errors"
	"log"
	"log/slog"
	"server-a/server/dto"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateMemberByEmail(email, password string) (map[string]string, error) {
	i, err := s.repository.EmailExists(email)
	if err != nil {
		return nil, err
	}
	if i {
		log.Printf("this email already exist")
		return nil, errors.New("this email already exist")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("fail to hash password: %v", err)
		return nil, err
	}
	password = string(hashedPassword)
	id := gocql.TimeUUID()
	err = s.repository.SaveEmailMember(id, email, password)
	if err != nil {
		return nil, err
	}
	return map[string]string{"id": id.String()}, nil
}

func (s *Service) LoginWithEmail(email, password string) (resp dto.EmailLoginResp, refreshToken string, err error) {
	emailVerified, phoneNumberVerified, id, dbPassword, role, err :=
		s.repository.FindLoginInfoByEmail(email)
	if err != nil {
		return resp, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		slog.Info("invalid password",
			"err", err,
		)
		return resp, "", err
	}
	if !emailVerified {
		resp.EmailVerified = false
		resp.PhoneNumberVerified = false

		return resp, "", nil
	}
	if !phoneNumberVerified {
		sid := gocql.TimeUUID()
		err = s.repository.SaveEmailBySessionId(sid, email)
		//resp := dto.EmailLoginResp{
		//	EmailVerified:       true,
		//	PhoneNumberVerified: false,
		//	SessionId:           sid.String(),
		//}
		resp.EmailVerified = true
		resp.PhoneNumberVerified = false
		resp.SessionId = sid.String()
		return resp, "", nil
	}
	at, rt, err := s.createLoginTokens(id.String(), role)
	if err != nil {
		return resp, "", nil
	}
	err = s.repository.SaveRefreshTokenById(id, rt)
	if err != nil {
		return resp, "", err
	}
	resp.EmailVerified = true
	resp.PhoneNumberVerified = false
	resp.AccessToken = at
	return resp, rt, nil
}

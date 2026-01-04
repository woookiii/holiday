package service

import (
	"errors"
	"log"
	"log/slog"
	"server-a/server/dto"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateMember(req *dto.MemberSaveReq) (map[string]string, error) {
	i, err := s.repository.EmailExists(req.Email)
	if err != nil {
		return nil, err
	}
	if i {
		log.Printf("this email already exist")
		return nil, errors.New("this email already exist")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("fail to hash password: %v", err)
		return nil, err
	}
	req.Password = string(hashedPassword)
	secret, err := generateUserSecret()
	if err != nil {
		return nil, err
	}
	id := gocql.TimeUUID()
	err = s.repository.SaveMember(req, id, secret)
	if err != nil {
		return nil, err
	}
	return map[string]string{"id": id.String()}, nil
}

func (s *Service) LoginWithEmail(req *dto.MemberLoginReq) (*dto.EmailLoginResp /*refreshToken*/, string, error) {
	m, err := s.repository.FindLoginInfoByEmail(req.Email)
	if err != nil {
		slog.Info("fail to login with email", req.Email, err)
		return nil, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(req.Password))
	if err != nil {
		slog.Info("invalid password", req.Password, err)
		return nil, "", err
	}
	if !m.IsEmailVerified {
		resp := dto.EmailLoginResp{}
		return &resp, "", nil
	}
	if !m.IsPhoneNumberVerified {
		sid := gocql.TimeUUID()
		err = s.repository.SaveEmailBySessionId(sid, req.Email)
		resp := dto.EmailLoginResp{
			IsEmailVerified: m.IsEmailVerified,
			SessionId:       sid.String(),
		}
		return &resp, "", nil
	}
	at, err := createToken(m.Id.String(), m.Role, s.secretKeyAT, s.aTExp)
	if err != nil {
		slog.Error("fail to create access token", err)
		return nil, "", err
	}
	rt, err := createToken(m.Id.String(), m.Role, s.secretKeyRT, s.rTExp)
	if err != nil {
		slog.Error("fail to create refresh token", err)
		return nil, "", err
	}
	err = s.repository.SaveRefreshToken(m.Id, rt)
	if err != nil {
		return nil, "", err
	}
	resp := dto.EmailLoginResp{
		IsEmailVerified:       m.IsEmailVerified,
		IsPhoneNumberVerified: m.IsPhoneNumberVerified,
		AccessToken:           at,
	}
	return &resp, rt, nil
}

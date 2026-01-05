package service

import (
	"errors"
	"log"
	"log/slog"
	"server-a/server/constant"
	"server-a/server/dto"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateMemberByEmail(req *dto.MemberSaveReq) (map[string]string, error) {
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
	id := gocql.TimeUUID()
	err = s.repository.SaveEmailMember(req, id)
	if err != nil {
		return nil, err
	}
	return map[string]string{"id": id.String()}, nil
}

func (s *Service) LoginWithEmail(req *dto.MemberLoginReq) (*dto.EmailLoginResp /*refreshToken*/, string, error) {
	m, err := s.repository.FindLoginInfoByEmail(req.Email)
	if err != nil {
		return nil, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(req.Password))
	if err != nil {
		slog.Info("invalid password", req.Password, err)
		return nil, "", err
	}
	if !m.EmailVerified {
		resp := dto.EmailLoginResp{
			EmailVerified:       false,
			PhoneNumberVerified: false,
		}
		return &resp, "", nil
	}
	if !m.PhoneNumberVerified {
		sid := gocql.TimeUUID()
		err = s.repository.SaveEmailBySessionId(sid, req.Email)
		resp := dto.EmailLoginResp{
			EmailVerified:       true,
			PhoneNumberVerified: false,
			SessionId:           sid.String(),
		}
		return &resp, "", nil
	}
	at, err := createToken(m.Id.String(), m.Role, s.secretKeyAT, constant.ACCESS_TOKEN_TTL)
	if err != nil {
		slog.Error("fail to create access token", err)
		return nil, "", err
	}
	rt, err := createToken(m.Id.String(), m.Role, s.secretKeyRT, constant.REFRESH_TOKEN_TTL)
	if err != nil {
		slog.Error("fail to create refresh token", err)
		return nil, "", err
	}
	err = s.repository.SaveRefreshTokenById(m.Id, rt)
	if err != nil {
		return nil, "", err
	}
	resp := dto.EmailLoginResp{
		EmailVerified:       true,
		PhoneNumberVerified: true,
		AccessToken:         at,
	}
	return &resp, rt, nil
}

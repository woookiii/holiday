package service

import (
	"errors"
	"log"
	"server-a/server/dto"
	"server-a/server/entity"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateMember(req *dto.MemberSaveReq) (*entity.Member, error) {
	i, err := s.repository.EmailExists(req.Email)
	if err != nil {
		log.Printf("fail to create member: %v", err)
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
	m, err := s.repository.SaveMember(req, secret)
	if err != nil {
		log.Printf("fail to create member: %v", err)
		return nil, err
	}

	return m, nil
}

func (s *Service) Login(req *dto.MemberLoginReq) (*dto.Token, error) {
	m, err := s.repository.FindIdPasswordRoleByEmail(req.Email)
	if err != nil {
		log.Printf("fail to login - email: %v, err: %v", req.Email, err)
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(req.Password))
	if err != nil {
		log.Printf("invalid password: %v, err: %v", req.Password, err)
		return nil, err
	}
	at, err := createToken(m.Id.String(), m.Role, s.secretKeyAT, s.aTExp)
	if err != nil {
		log.Printf("fail to create access token: %v", err)
		return nil, err
	}
	rt, err := createToken(m.Id.String(), m.Role, s.secretKeyRT, s.rTExp)
	if err != nil {
		log.Printf("fail to create refresh token: %v", err)
		return nil, err
	}
	err = s.repository.SaveRefreshToken(m.Id, rt)
	t := &dto.Token{
		AccessToken:  at,
		RefreshToken: rt,
	}
	return t, nil
}

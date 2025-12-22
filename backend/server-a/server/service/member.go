package service

import (
	"errors"
	"log"
	"server-a/server/dto"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateMember(req *dto.MemberSaveReq) error {
	i, err := s.repository.EmailExists(req.Email)
	if err != nil {
		log.Printf("fail to create member: %v", err)
		return err
	}
	if i == true {
		log.Printf("this email already exist")
		return errors.New("this email already exist")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("fail to hash password: %v", err)
		return err
	}
	req.Password = string(hashedPassword)
	err = s.repository.SaveMember(req)
	if err != nil {
		log.Printf("fail to create member: %v", err)
		return err
	}
	return nil
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
	sAT := s.secretKeyAT
	sRT := s.secretKeyRT
	at, err := createToken(m.Id.String(), m.Role, sAT, s.aTExp)
	if err != nil {
		log.Printf("fail to create access token: %v", err)
		return nil, err
	}
	rt, err := createToken(m.Id.String(), m.Role, sRT, s.rTExp)
	if err != nil {
		log.Printf("fail to create refresh token: %v", err)
		return nil, err
	}
	t := &dto.Token{
		AccessToken:  at,
		RefreshToken: rt,
	}
	return t, nil
}

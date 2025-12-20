package service

import (
	"log"
	"server-a/server/dto"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateMember(req *dto.MemberSaveReq) error {
	i, err := s.repository.FindByEmail(req.Email)
	if err != nil {
		log.Printf("fail to create member: %v", err)
		return err
	}
	if i == true {
		log.Printf("this email already exist")
		return nil
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

func (s *Service) Login(dto *dto.MemberLoginReq) (*dto.Token, error) {

	p, err := s.repository.FindPasswordByEmail(dto.Email)
	if err != nil {
		log.Printf("fail to login - email: %v, err: %v", dto.Email, dto.Email, err)
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(p), []byte(dto.Password))
	if err != nil {
		log.Printf("invalid password: %v, err: %v", dto.Password, err)
	}

}

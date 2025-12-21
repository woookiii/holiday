package service

import (
	"log"
	"os"
	"server-a/server/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	secretKeyAt := []byte(os.Getenv("SECRET_KEY_AT"))
	secretKeyRt := []byte(os.Getenv("SECRET_KEY_RT"))
	at, err := createToken(m.Id.String(), m.Role, secretKeyAt, 10)
	if err != nil {
		log.Printf("fail to create access token: %v", err)
		return nil, err
	}
	rt, err := createToken(m.Id.String(), m.Role, secretKeyRt, 100000)
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

func createToken(id, role string, secretKey []byte, expirationMinutes int64) (string, error) {
	claims := jwt.MapClaims{
		"sub":  id,
		"role": role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Minute * time.Duration(expirationMinutes)).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		log.Printf("fail to make token")
		return "", err
	}
	return token, nil
}

package service

import (
	"errors"
	"fmt"
	"log"
	"server-a/server/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) GenerateAccessToken(refreshToken string) (*dto.Token, error) {
	token, err := jwt.Parse(refreshToken, s.keyFunc)
	if err != nil {
		log.Printf("fail to parse token: %v", err)
		return nil, err
	}
	if !token.Valid {
		log.Printf("invalid token: %v", token)
		return nil, errors.New("invalid token")
	}
	id, err := token.Claims.GetSubject()
	if err != nil {
		i
	}
	return token, nil
}

func (s *Service) keyFunc(token *jwt.Token) (any, error) {
	if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
		return nil, fmt.Errorf("unexpected signing method: %s", token.Method.Alg())
	}
	return s.secretKeyRT, nil
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

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
	rt, err := jwt.Parse(refreshToken, s.keyFunc)
	if err != nil {
		log.Printf("fail to parse token: %v", err)
		return nil, err
	}
	if !rt.Valid {
		log.Printf("invalid token: %v", rt)
		return nil, errors.New("invalid token")
	}
	id, err := rt.Claims.GetSubject()
	if err != nil {
		log.Printf("fail to get subject from claim: %v", err)
		return nil, err
	}
	tokenInDB, err := s.repository.FindTokenById(id)
	if err != nil {
		log.Printf("fail to find token: %v", err)
		return nil, err
	}
	if rt.Raw != tokenInDB.RefreshToken {
		log.Printf("token is not same with DB: %v", rt.Raw)
		return nil, fmt.Errorf("invalid token: %v", rt.Raw)
	}
	role, ok := rt.Claims.(jwt.MapClaims)["role"].(string)
	if !ok {
		log.Printf("fail to get role - token: %v", rt.Raw)
		return nil, fmt.Errorf("fail to get role - token: %v", rt.Raw)
	}
	at, err := createToken(id, role, s.secretKeyAT, s.aTExp)
	if err != nil {
		return nil, err
	}

	return &dto.Token{AccessToken: at}, nil
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

package service

import (
	"os"
	"server-a/server/repository"
)

type Service struct {
	repository  *repository.Repository
	secretKeyAT []byte
	secretKeyRT []byte
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		repository:  r,
		secretKeyAT: []byte(os.Getenv("SECRET_KEY_AT")),
		secretKeyRT: []byte(os.Getenv("SECRET_KEY_RT")),
	}
}

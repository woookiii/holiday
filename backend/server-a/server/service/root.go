package service

import (
	"os"
	"server-a/config"
	"server-a/server/repository"
)

type Service struct {
	repository  *repository.Repository
	secretKeyAT []byte
	secretKeyRT []byte
	aTExp       int64
	rTExp       int64
}

func NewService(cfg *config.Config, r *repository.Repository) *Service {
	return &Service{
		repository:  r,
		secretKeyAT: []byte(os.Getenv("SECRET_KEY_AT")),
		secretKeyRT: []byte(os.Getenv("SECRET_KEY_RT")),
		aTExp:       cfg.Exp.ATExp,
		rTExp:       cfg.Exp.RtExp,
	}
}

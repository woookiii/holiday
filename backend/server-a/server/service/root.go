package service

import (
	"server-a/server/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{r}
}

package server

import (
	"server-a/config"
	"server-a/server/network"
	"server-a/server/repository"
	"server-a/server/service"
)

type Server struct {
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	server := &Server{cfg}

	r := repository.NewRepository(cfg)

	s := service.NewService(cfg, r)

	n := network.NewNetwork(cfg, s)

	go func() {
		n.Start()
	}()

	return server
}

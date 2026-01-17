package server

import (
	"server-a/config"
	"server-a/server/kafka/producer"
	"server-a/server/logger"
	"server-a/server/network"
	"server-a/server/repository"
	"server-a/server/service"
)

type Server struct {
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	server := &Server{cfg}

	kp := producer.NewKafkaProducer(cfg)

	logger.SetLogger(kp)

	r := repository.NewRepository(cfg)

	s := service.NewService(cfg, r)

	n := network.NewNetwork(cfg, s)

	n.Start()

	return server
}

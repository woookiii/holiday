package src

import (
	"server-a/config"
	"server-a/src/kafka/producer"
	"server-a/src/network"
	"server-a/src/repository"
	"server-a/src/service"
)

type Server struct {
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	server := &Server{cfg}

	kp := producer.NewKafkaProducer(cfg)

	r := repository.NewRepository(cfg)

	s := service.NewService(cfg, r)

	n := network.NewNetwork(cfg, s)

	n.Start()

	return server
}

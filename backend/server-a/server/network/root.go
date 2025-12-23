package network

import (
	"server-a/config"
	"server-a/server/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	service *service.Service
	engine  *gin.Engine
	port    string
	rtExp   int64
}

func NewNetwork(cfg *config.Config, s *service.Service) *Network {
	n := &Network{
		service: s,
		engine:  gin.New(),
		port:    cfg.Info.Port,
		rtExp:   cfg.Exp.RtExp,
	}

	setGin(n.engine)
	memberRouter(n)
	tokenRouter(n)

	return n
}

func (n *Network) Start() error {
	return n.engine.Run(n.port)
}

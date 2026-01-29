package network

import (
	"github.com/gin-gonic/gin"
)

type HTTPMethod int

const (
	GET HTTPMethod = iota
	POST
	DELETE
	PUT
)

func setGin(engine *gin.Engine) {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
}

func (n *Network) Router(httpMethod HTTPMethod, path string, handler ...gin.HandlerFunc) {
	e := n.engine.Group("/auth")

	switch httpMethod {
	case GET:
		e.GET(path, handler...)
	case POST:
		e.POST(path, handler...)
	case PUT:
		e.PUT(path, handler...)
	case DELETE:
		e.DELETE(path, handler...)

	default:
		panic("This HTTP method is not registered")
	}
}

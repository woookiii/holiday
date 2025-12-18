package network

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type HTTPMethod int

const (
	GET HTTPMethod = iota
	POST
	DELETE
	PUT
)

type header struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type response struct {
	*header
	Result any `json:"result"`
}

func res(ctx *gin.Context, statusCode int, result any, messages ...string) {
	ctx.JSON(statusCode, &response{
		header: &header{StatusCode: statusCode, Message: strings.Join(messages, ",")},
		Result: result,
	})
}

func setGin(engine *gin.Engine) {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	//TODO: cors if needed
}

func (n *Network) Router(httpMethod HTTPMethod, path string, handler gin.HandlerFunc) {
	e := n.engine

	switch httpMethod {
	case GET:
		e.GET(path, handler)
	case POST:
		e.POST(path, handler)
	case PUT:
		e.PUT(path, handler)
	case DELETE:
		e.DELETE(path, handler)

	default:
		panic("This HTTP method is not registered")
	}
}

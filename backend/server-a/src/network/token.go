package network

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func tokenRouter(n *Network) {
	n.Router(POST, "/refresh-token", n.refreshToken)
}

func (n *Network) refreshToken(c *gin.Context) {
	rt, err := c.Request.Cookie("refresh_token")
	if err != nil {
		res(c, http.StatusUnauthorized, err.Error())
		return
	}
	result, err := n.service.GenerateAccessToken(rt.Value)
	if err != nil {
		res(c, http.StatusUnauthorized, err.Error())
		return
	}
	res(c, http.StatusOK, result)

}

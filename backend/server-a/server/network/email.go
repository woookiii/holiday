package network

import (
	"net/http"
	"server-a/server/dto"

	"github.com/gin-gonic/gin"
)

func emailRouter(n *Network) {
	n.Router(GET, "/email/check/:email", n.checkEmail)
	n.Router(POST, "/email/send-code/:email", n.sendCode)
	n.Router(POST, "/email/validate-code", n.validateCode)
}

func (n *Network) checkEmail(c *gin.Context) {
	email := c.Param("email")
	ok, err := n.service.IsEmailUsable(email)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	res(c, http.StatusOK, ok)
}

func (n *Network) sendCode(c *gin.Context) {
	email := c.Param("email")
	err := n.service.SendEmailVerifyCode(email)
	if err != nil {
		res(c, http.StatusInternalServerError, err.Error())
		return
	}
	res(c, http.StatusOK, "ok")
}

func (n *Network) validateCode(c *gin.Context) {
	var req dto.EmailValidateReq
	err := c.ShouldBindJSON(req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	err = n.service.ValidateEmailVerifyCode(&req)
	if err != nil {
		res(c, http.StatusUnauthorized, err.Error())
		return
	}
	res(c, http.StatusOK, "ok")
}

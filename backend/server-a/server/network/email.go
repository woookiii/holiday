package network

import (
	"net/http"
	"server-a/server/dto"

	"github.com/gin-gonic/gin"
)

func emailRouter(n *Network) {
	n.Router(GET, "/email/check", n.checkEmail)
	n.Router(POST, "/email/otp/send", n.sendEmailOTP)
	n.Router(POST, "/email/otp/verify", n.verifyEmailOTP)
}

func (n *Network) checkEmail(c *gin.Context) {
	var req dto.EmailReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	ok, err := n.service.IsEmailUsable(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	res(c, http.StatusOK, ok)
}

func (n *Network) sendEmailOTP(c *gin.Context) {
	var req dto.EmailReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := n.service.SendEmailOTP(&req)
	if err != nil {
		res(c, http.StatusInternalServerError, err.Error())
		return
	}
	res(c, http.StatusOK, result)
}

func (n *Network) verifyEmailOTP(c *gin.Context) {
	var req dto.OTPVerifyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := n.service.VerifyEmailOTP(&req)
	if err != nil {
		res(c, http.StatusUnauthorized, err.Error())
		return
	}
	res(c, http.StatusOK, result)
}

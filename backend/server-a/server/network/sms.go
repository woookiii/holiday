package network

import (
	"net/http"
	"server-a/server/dto"

	"github.com/gin-gonic/gin"
)

func smsRouter(n *Network) {
	n.Router(POST, "/auth/sms/otp/send", n.sendSMSOTP)
	n.Router(POST, "/auth/sms/otp/verify", n.verifySMSOTP)
}

func (n *Network) sendSMSOTP(c *gin.Context) {
	var req dto.SMSOTPSendReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := n.service.SendSMSOTP(&req)
	if err != nil {
		res(c, http.StatusInternalServerError, err.Error())
		return
	}
	res(c, http.StatusOK, result)
}

func (n *Network) verifySMSOTP(c *gin.Context) {
	var req dto.OTPVerifyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	err = n.service.VerifySMSOTP(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	res(c, http.StatusOK, "ok")
}

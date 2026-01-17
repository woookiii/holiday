package network

import (
	"net/http"
	"server-a/server/constant"
	"server-a/server/dto"

	"github.com/gin-gonic/gin"
)

func smsRouter(n *Network) {
	n.Router(POST, "/sms/otp/send", n.sendSMSOTP)
	n.Router(POST, "/sms/otp/verify", n.verifySMSOTP)
}

func (n *Network) sendSMSOTP(c *gin.Context) {
	var req dto.SMSOTPSendReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := n.service.SendSMSOTP(req.PhoneNumber)
	if err != nil {
		res(c, http.StatusInternalServerError, err.Error())
		return
	}
	res(c, http.StatusOK, result)
}

func (n *Network) verifySMSOTP(c *gin.Context) {
	var req dto.SMSOTPVerifyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	result, rt, err := n.service.VerifySMSOTP(req.SessionId, req.OTP, req.VerificationId)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	if rt != "" {
		c.SetCookie("refresh_token",
			rt,
			constant.RefreshTokenTTL,
			"",
			"",
			false,
			true,
		)
	}
	res(c, http.StatusOK, result)
}

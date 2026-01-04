package network

import (
	"net/http"
	"server-a/server/dto"

	"github.com/gin-gonic/gin"
)

func smsRouter(n *Network) {
	n.Router(POST, "/auth/sms/otp/send", n.sendSmsOtp)
	n.Router(POST, "/auth/sms/otp/verify", n.verifySmsOtp)
}

func (n *Network) sendSmsOtp(c *gin.Context) {
	var req dto.SmsOTPSendReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := n.service.SendSmsOtp(&req)
	if err != nil {
		res(c, http.StatusInternalServerError, err.Error())
		return
	}
	res(c, http.StatusOK, result)
}

func (n *Network) verifySmsOtp(c *gin.Context) {
	var req dto.OTPVerifyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	err = n.service.VerifySmsOtp(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	res(c, http.StatusOK, "ok")
}

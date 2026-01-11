package network

import (
	"net/http"
	"server-a/server/constant"
	"server-a/server/dto"

	"github.com/gin-gonic/gin"
)

func emailRouter(n *Network) {
	n.Router(POST, "/email/create", n.createMemberByEmail)
	n.Router(POST, "/email/login", n.loginWithEmail)
	n.Router(GET, "/email/check", n.checkEmail)
	n.Router(POST, "/email/otp/send", n.sendEmailOTP)
	n.Router(POST, "/email/otp/verify", n.verifyEmailOTP)
}

func (n *Network) createMemberByEmail(c *gin.Context) {
	var req dto.MemberSaveReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	result, err := n.service.CreateMemberByEmail(req.Email, req.Password)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	res(c, http.StatusOK, result)
}

func (n *Network) loginWithEmail(c *gin.Context) {
	var req dto.MemberLoginReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusUnauthorized, err.Error())
		return
	}
	result, rt, err := n.service.LoginWithEmail(req.Email, req.Password)
	if err != nil {
		res(c, http.StatusUnauthorized, err.Error())
		return
	}
	if rt != "" {
		c.SetCookie("refresh_token",
			rt,
			constant.REFRESH_TOKEN_TTL,
			"",
			"",
			false,
			true,
		)
	}
	res(c, http.StatusOK, result)
}

func (n *Network) checkEmail(c *gin.Context) {
	var req dto.EmailReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	ok, err := n.service.IsEmailUsable(req.Email)
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
	result, err := n.service.SendEmailOTP(req.Email)
	if err != nil {
		res(c, http.StatusInternalServerError, err.Error())
		return
	}
	res(c, http.StatusOK, result)
}

func (n *Network) verifyEmailOTP(c *gin.Context) {
	var req dto.EmailOTPVerifyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := n.service.VerifyEmailOTP(req.OTP, req.VerificationId)
	if err != nil {
		res(c, http.StatusUnauthorized, err.Error())
		return
	}
	res(c, http.StatusOK, result)
}

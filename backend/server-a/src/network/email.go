package network

import (
	"net/http"
	"server-a/src/constant"
	"server-a/src/dto"

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
	var req dto.EmailMemberSaveReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	result, err := n.service.CreateMemberByEmail(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	res(c, http.StatusOK, result)
}

func (n *Network) loginWithEmail(c *gin.Context) {
	var req dto.EmailMemberLoginReq

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
			constant.RefreshTokenTTL,
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
	ctx := c.Request.Context()
	ok, err := n.service.IsEmailUsable(ctx, req.Email)
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
	result, err := n.service.SendEmailOTP(c.Request.Context(), req.Email)
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

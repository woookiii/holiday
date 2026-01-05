package network

import (
	"net/http"
	"server-a/server/dto"

	"github.com/gin-gonic/gin"
)

func memberRouter(n *Network) {
	n.Router(POST, "/member/email/create", n.createMemberByEmail)
	n.Router(POST, "/member/email/login", n.loginWithEmail)
}

func (n *Network) createMemberByEmail(c *gin.Context) {
	var req dto.MemberSaveReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	result, err := n.service.CreateMemberByEmail(&req)
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
	result, rt, err := n.service.LoginWithEmail(&req)
	if err != nil {
		res(c, http.StatusUnauthorized, err.Error())
		return
	}
	if rt != "" {
		c.SetCookie("refresh_token",
			rt,
			int(n.rtExp),
			"",
			"",
			false,
			true,
		)
	}
	res(c, http.StatusOK, result)
}

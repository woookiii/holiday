package network

import (
	"net/http"
	"server-a/server/dto"

	"github.com/gin-gonic/gin"
)

func memberRouter(n *Network) {
	n.Router(POST, "/member/create", n.createMember)
	n.Router(POST, "/member/login", n.login)
}

func (n *Network) createMember(c *gin.Context) {
	var req dto.MemberSaveReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	m, err := n.service.CreateMember(&req)
	if err != nil {
		res(c, http.StatusBadRequest, err.Error())
		return
	}
	res(c, http.StatusOK, m.Id.String())
}

func (n *Network) login(c *gin.Context) {
	var req dto.MemberLoginReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusUnauthorized, err.Error())
		return
	}
	t, err := n.service.Login(&req)
	if err != nil {
		res(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.SetCookie("refresh-token",
		t.RefreshToken,
		int(n.rtExp),
		"",
		"",
		false,
		true,
	)
	res(c, http.StatusOK, map[string]any{"accessToken": t.AccessToken})
}

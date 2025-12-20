package network

import (
	"net/http"
	"server-a/server/dto"

	"github.com/gin-gonic/gin"
)

func memberRouter(n *Network) {

	n.Router(POST, "/member/create", n.createMember)
	n.Router(POST, "/member/login", n.login)
	n.Router(POST, "/member/refresh-token", n.refreshToken)
}

func (n *Network) createMember(c *gin.Context) {
	var req dto.MemberSaveReq

	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err)
	} else if err = n.service.CreateMember(&req); err != nil {
		res(c, http.StatusInternalServerError, err)
	} else {
		res(c, http.StatusOK, "Success create member")
	}
}

func (n *Network) login(c *gin.Context) {
	var req dto.MemberLoginReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res(c, http.StatusUnauthorized, err)
	}
	t, err := n.service.Login(&req)
	if err != nil {
		res(c, http.StatusUnauthorized, err)
	}
	c.SetCookie("refresh-token",
		t.RefreshToken,
		1000000000000,
		"",
		"",
		true,
		true,
	)
	res(c, http.StatusOK, t.AccessToken)

}

func (n *Network) refreshToken(c *gin.Context) {
	//rt, err := c.Request.Cookie("refresh-token")
	//if err != nil {
	//	res(c, http.StatusUnauthorized, err)
	//}
}

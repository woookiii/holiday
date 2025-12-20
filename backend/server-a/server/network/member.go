package network

import (
	"net/http"
	"server-a/server/dto"

	"github.com/gin-gonic/gin"
)

func memberRouter(n *Network) {

	n.Router(POST, "/member/create", n.createMember)
}

func (n *Network) createMember(c *gin.Context) {
	var req dto.MemberSaveReq

	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else if err = n.service.CreateMember(&req); err != nil {
		res(c, http.StatusInternalServerError, err.Error())
	} else {
		res(c, http.StatusOK, "Success create member")
	}
}

package network

func memberRouter(n *Network) {

	//n.Router(POST, "/member/create", n.createMember)
}

//func (n *Network) createMember(ctx *gin.Context) {
//	var req dto.MemberSaveReq
//
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		res(ctx, http.StatusUnprocessableEntity, err.Error())
//	} else if err = n.service.CreateMember(&req); err != nil {
//		res(ctx, http.StatusInternalServerError, err.Error())
//	} else {
//		res(ctx, http.StatusOK, "Success create member")
//	}
//}

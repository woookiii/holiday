package network

import (
	"log"
	"net/http"
	"server-a/server/dto"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func memberRouter(n *Network) {

	//n.Router(POST, "/member/create", n.createMember)
}

func (s *Service) CreateMember(req *dto.MemberSaveReq) error {
	if member, _ := s.repository.FindMemberByEmail(req.Email); member != nil {
		log.Println("This email already exist")
		return nil
	} else if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost); err != nil {
		return err
	} else if registeredMember, err := s.repository.CreateMember(
		req.Name,
		req.Email,
		string(hashedPassword),
	); err != nil {
		log.Println("Failed to create member", "Member name", req.Name, "err", err.Error())
		return err
	} else {
		log.Println("Success create new member", "Member name", req.Name)


		return nil
	}

}
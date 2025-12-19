package service

import (
	"log"
	"server-a/server/entity"
)

func (s *Service) CreateMember(member *entity.Member) error {

	err := s.repository.SaveMember(member)
	if err != nil {
		log.Printf("Failed to save member to redis member id: %v, err : %v", member.Id, err)
		return err
	}

	log.Printf("Success to save new member to redis member id: %v", member.Id)

	return nil
}

package repository

import (
	"log"
	"server-a/server/entity"
)

func (r *Repository) SaveEmailValidationCode(email, code string) error {
	err := r.session.Query(
		"INSERT INTO member_by_email (email, code) VALUES (?, ?) USING TTL ?",
		email, code, 300,
	).Exec()
	if err != nil {
		log.Printf("fail to save email validatation code: %v", err)
		return err
	}
	return nil
}

func (r *Repository) FindCodeByEmail(email string) (*entity.Member, error) {
	var m entity.Member
	err := r.session.Query(
		"SELECT code FROM member_by_email WHERE email = ?",
		email,
	).Scan(&m.Code)
	if err != nil {
		log.Printf("fail to select code by email: %v", err)
		return nil, err
	}
	return &m, nil
}

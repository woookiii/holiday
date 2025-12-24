package repository

import "log"

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

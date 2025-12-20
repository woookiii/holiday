package repository

import (
	"errors"
	"log"
	"server-a/server/dto"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

func (r *Repository) SaveMember(req *dto.MemberSaveReq) error {
	err := r.session.Batch(gocql.LoggedBatch).
		Query("INSERT INTO member_email (email) VALUES (?);", req.Email).
		Query(""+
			"INSERT INTO member_by_id (id, name, email, password, role, created_time) VALUES (?, ?, ?, ?, ?, ?)",
			gocql.TimeUUID(), req.Name, req.Email, req.Password, "user", time.Now(),
		).
		Exec()
	if err != nil {
		log.Printf("fail to save member: %v", err)
		return err
	}
	return nil
}

func (r *Repository) IsEmailAlreadyUsed(email string) (bool, error) {

	err := r.session.Query(
		"SELECT * FROM member_email WHERE email = ?",
		email,
	).Exec()
	if errors.Is(err, gocql.ErrNotFound) {
		return false, nil
	}
	if err != nil {
		log.Printf("fail to select email: %v", err)
		return true, err
	}
	return true, nil
}

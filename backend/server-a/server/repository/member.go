package repository

import (
	"log"
	"server-a/server/dto"
	"server-a/server/entity"
	"time"

	"github.com/apache/cassandra-gocql-driver/v2"
)

func (r *Repository) SaveMember(req *dto.MemberSaveReq, id gocql.UUID, secret string) error {
	err := r.session.Batch(gocql.LoggedBatch).
		Query(
			"INSERT INTO member_by_email (is_email_verified, is_phone_number_verified, id, email, password, secret, role, created_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
			false, false, id, req.Email, req.Password, secret, "user", time.Now(),
		).
		Query(
			"INSERT INTO member_by_id (is_email_verified, is_phone_number_verified, id, email, password, secret, role, created_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			false, false, id, req.Email, req.Password, secret, "user", time.Now(),
		).
		Exec()
	if err != nil {
		log.Printf("fail to save member: %v", err)
		return err
	}
	return err
}

func (r *Repository) EmailExists(email string) (bool, error) {
	var c int64
	err := r.session.Query(
		"SELECT COUNT(1) FROM member_by_email WHERE email = ?",
		email,
	).Scan(&c)
	if c == 0 {
		return false, nil
	}
	if err != nil {
		log.Printf("fail to find by email: %v", err)
		return true, err
	}
	return true, nil
}

func (r *Repository) FindLoginInfoByEmail(email string) (*entity.Member, error) {
	var m entity.Member
	err := r.session.Query(
		"SELECT is_email_verified, is_phone_number_verified, id, password, role FROM member_by_email WHERE email = ?",
		email,
	).Scan(&m.VerificationId, &m.Id, &m.Password, &m.Role)
	if err != nil {
		log.Printf("fail to find by email: %v", err)
		return nil, err
	}
	return &m, nil
}

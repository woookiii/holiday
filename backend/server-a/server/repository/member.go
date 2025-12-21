package repository

import (
	"log"
	"server-a/server/dto"
	"server-a/server/entity"
	"time"

	"github.com/apache/cassandra-gocql-driver/v2"
)

func (r *Repository) SaveMember(req *dto.MemberSaveReq) error {
	id := gocql.TimeUUID()
	err := r.session.Batch(gocql.LoggedBatch).
		Query(
			"INSERT INTO member_by_email (id, name, email, password, role, created_time) VALUES (?, ?, ?, ?, ?, ?);",
			id, req.Name, req.Email, req.Password, "user", time.Now(),
		).
		Query(
			"INSERT INTO member_by_id (id, name, email, password, role, created_time) VALUES (?, ?, ?, ?, ?, ?)",
			id, req.Name, req.Email, req.Password, "user", time.Now(),
		).
		Exec()
	if err != nil {
		log.Printf("fail to save member: %v", err)
		return err
	}
	return nil
}

func (r *Repository) EmailExists(email string) (bool, error) {
	var c int64
	err := r.session.Query(
		"SELECT COUNT(*) FROM member_by_email WHERE email = ?",
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

func (r *Repository) FindIdPasswordRoleByEmail(email string) (*entity.Member, error) {
	var m entity.Member
	err := r.session.Query(
		"SELECT id, password, role FROM member_by_email WHERE email = ?",
		email,
	).Scan(&m.Id, &m.Password, &m.Role)
	if err != nil {
		log.Printf("fail to find by email: %v", err)
		return nil, err
	}
	return &m, nil
}

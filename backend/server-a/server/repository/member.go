package repository

import (
	"log/slog"
	"server-a/server/constant"
	"time"

	"github.com/apache/cassandra-gocql-driver/v2"
)

func (r *Repository) SaveEmailMember(id gocql.UUID, email, password string) error {
	err := r.session.Batch(gocql.LoggedBatch).
		Query(
			"INSERT INTO member_by_email (email_verified, phone_number_verified, id, email, password, role, created_time) VALUES (?, ?, ?, ?, ?, ?, ?);",
			false, false, id, email, password, constant.ROLE_USER, time.Now()).
		Query(
			"INSERT INTO member_by_id (email_verified, phone_number_verified, id, email, role, created_time) VALUES (?, ?, ?, ?, ?, ?)",
			false, false, id, email, constant.ROLE_USER, time.Now()).
		Exec()
	if err != nil {
		slog.Error("fail to save member",
			"err", err,
			"id", id.String(),
		)
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
		slog.Error("fail to check email existence",
			"err", err,
			"email", email,
		)
		return true, err
	}
	return true, nil
}

func (r *Repository) FindLoginInfoByEmail(email string) (emailVerified, phoneNumberVerified bool, id gocql.UUID, password, role string, err error) {
	err = r.session.Query(
		"SELECT email_verified, phone_number_verified, id, password, role FROM member_by_email WHERE email = ?",
		email,
	).Scan(&emailVerified, &phoneNumberVerified, &id, &password, &role)
	if err != nil {
		slog.Info("fail to find by email",
			"err", err,
			"email", email,
		)
		return false, false, gocql.UUID{}, "", "", err
	}
	return emailVerified, phoneNumberVerified, id, password, role, nil
}

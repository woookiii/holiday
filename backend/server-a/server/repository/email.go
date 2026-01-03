package repository

import (
	"log"
	"log/slog"
	"server-a/server/entity"

	"github.com/apache/cassandra-gocql-driver/v2"
)

func (r *Repository) SaveEmailAndOtpByVerificationId(verificationId gocql.UUID, email, otp string) error {
	err := r.session.Query(
		"INSERT INTO member_by_verification_id (verification_id, email, otp) VALUES (?, ?, ?) USING TTL ?",
		verificationId, email, otp, 300,
	).Exec()
	if err != nil {
		slog.Error("fail to save email otp by verificationId: %v", err)
		return err
	}
	return nil
}

func (r *Repository) FindMemberByVerificationId(verificationId string) (*entity.Member, error) {
	var m entity.Member
	err := r.session.Query(
		"SELECT email, otp FROM member_by_verification_id WHERE verification_id = ?",
		verificationId,
	).Scan(&m.Email, &m.OTP)
	if err != nil {
		log.Printf("fail to select code by email: %v", err)
		return nil, err
	}
	return &m, nil
}

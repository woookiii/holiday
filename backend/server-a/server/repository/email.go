package repository

import (
	"log/slog"
	"server-a/server/entity"

	"github.com/apache/cassandra-gocql-driver/v2"
)

const AUTH_ID_TTL = 500

func (r *Repository) SaveEmailAndOtpByVerificationId(verificationId gocql.UUID, email, otp string) error {
	err := r.session.Query(
		"INSERT INTO member_by_verification_id (verification_id, email, otp) VALUES (?, ?, ?) USING TTL ?",
		verificationId, email, otp, AUTH_ID_TTL,
	).Exec()
	if err != nil {
		slog.Error("fail to save email otp by verificationId: %v", err)
		return err
	}
	return nil
}

func (r *Repository) FindEmailAndOtpByVerificationId(verificationId gocql.UUID) (*entity.Member, error) {
	var m entity.Member
	err := r.session.Query(
		"SELECT email, otp FROM member_by_verification_id WHERE verification_id = ?",
		verificationId,
	).Scan(&m.Email, &m.OTP)
	if err != nil {
		slog.Info("fail to select email and otp by verification_id", verificationId, err)
		return nil, err
	}
	return &m, nil
}

func (r *Repository) MarkEmailVerified(email string) error {
	var m entity.Member
	err := r.session.Query(
		"SELECT id FROM member_by_email WHERE email = ?",
		email,
	).Scan(&m.Id)
	if err != nil {
		slog.Error("fail to select id by email", email, err)
		return err
	}
	err = r.session.Query(
		"UPDATE member_by_id SET is_email_verified = ? WHERE id = ?",
		true, m.Id,
	).Exec()
	if err != nil {
		slog.Error("fail to update is_email_verified at member_by_id", m.Id, email, err)
		return err
	}
	err = r.session.Query(
		"UPDATE member_by_email SET is_email_verified = ? WHERE email = ?",
		true, email,
	).Exec()
	if err != nil {
		slog.Error("fail to update is_email_verified at member_by_email", m.Id, email, err)
		return err
	}
	return nil
}

func (r *Repository) SaveEmailBySessionId(sessionId gocql.UUID, email string) error {
	err := r.session.Query(
		"INSERT INTO email_by_session_id (session_id, email) VALUES (?, ?) USING TTL ?",
		sessionId, email, AUTH_ID_TTL,
	).Exec()
	if err != nil {
		slog.Error("fail to insert email by session_id at email_by_session_id", sessionId, email, err)
	}
	return nil
}

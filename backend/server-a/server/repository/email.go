package repository

import (
	"log/slog"
	"server-a/server/constant"
	"server-a/server/entity"
	"time"

	"github.com/apache/cassandra-gocql-driver/v2"
)

func (r *Repository) SaveEmailAndOtpByVerificationId(verificationId gocql.UUID, email, otp string) error {
	err := r.session.Query(
		"INSERT INTO member_by_verification_id (verification_id, email, otp) VALUES (?, ?, ?) USING TTL ?",
		verificationId, email, otp, constant.AUTH_ID_TTL,
	).Exec()
	if err != nil {
		slog.Error("fail to save email otp by verificationId",
			"err", err,
		)
		return err
	}
	return nil
}

func (r *Repository) FindEmailAndOTPByVerificationId(verificationId gocql.UUID) (email string, otp string, err error) {
	err = r.session.Query(
		"SELECT email, otp FROM member_by_verification_id WHERE verification_id = ?",
		verificationId,
	).Scan(&email, &otp)
	if err != nil {
		slog.Info("fail to select email and otp by verification_id",
			"err", err,
			"verificationId", verificationId.String(),
		)
		return "", "", err
	}
	return email, otp, nil
}

func (r *Repository) MarkEmailVerified(email string) error {
	var m entity.Member
	err := r.session.Query(
		"SELECT id FROM member_by_email WHERE email = ?",
		email,
	).Scan(&m.Id)
	if err != nil {
		slog.Error("fail to select id by email",
			"err", err,
			"email", email,
		)
		return err
	}
	err = r.session.Query(
		"UPDATE member_by_id SET is_email_verified = ? WHERE id = ?",
		true, m.Id,
	).Exec()
	if err != nil {
		slog.Error("fail to update is_email_verified at member_by_id",
			"err", err,
			"id", m.Id.String(),
			"email", email,
		)
		return err
	}
	err = r.session.Query(
		"UPDATE member_by_email SET is_email_verified = ? WHERE email = ?",
		true, email,
	).Exec()
	if err != nil {
		slog.Error("fail to update is_email_verified at member_by_email",
			"err", err,
			"id", m.Id.String(),
			"email", email,
		)
		return err
	}
	return nil
}

func (r *Repository) SaveEmailBySessionId(sessionId gocql.UUID, email string) error {
	err := r.session.Query(
		"INSERT INTO member_by_session_id (session_id, email) VALUES (?, ?) USING TTL ?",
		sessionId, email, constant.AUTH_ID_TTL,
	).Exec()
	if err != nil {
		slog.Error("fail to insert email by session_id at email_by_session_id",
			"err", err,
			"sessionId", sessionId.String(),
			"email", email,
		)
	}
	return nil
}

func (r *Repository) FindEmailBySessionId(sessionId gocql.UUID) (email string, err error) {
	err = r.session.Query(
		"SELECT email FROM member_by_session_id WHERE session_id = ?",
		sessionId,
	).Scan(&email)
	if err != nil {
		slog.Info("fail to find email by sessionId, might be expired sessionId",
			"err", err,
			"sessionId", sessionId,
		)
	}
	return email, nil
}

func (r *Repository) FindMemberInfoByEmail(email string) (id gocql.UUID, role string, t time.Time, err error) {
	err = r.session.Query("SELECT id, role, created_time FROM member_by_email WHERE email = ?",
		email,
	).Scan(&id, &role, &t)
	if err != nil {
		slog.Error("fail to select id at member_by_email",
			"err", err,
			"email", email,
		)
		return gocql.UUID{}, "", time.Time{}, err
	}
	return id, role, t, nil
}

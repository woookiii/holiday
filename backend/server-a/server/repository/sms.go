package repository

import (
	"errors"
	"log/slog"
	"server-a/server/constant"
	"time"

	"github.com/apache/cassandra-gocql-driver/v2"
)

func (r *Repository) SavePhoneNumberByVerificationId(verificationId gocql.UUID, phoneNumber string) error {
	err := r.session.Query("INSERT INTO member_by_verification_id (phone_number, verification_id) values (?,?) USING TTL ?",
		phoneNumber, verificationId, constant.OTP_TTL,
	).Exec()
	if err != nil {
		slog.Error("fail to insert phone number with id",
			"err", err,
			"verificationId", verificationId,
			"phoneNumber", phoneNumber,
		)
		return err
	}
	return nil
}

func (r *Repository) FindPhoneNumberByVerificationId(verificationId gocql.UUID) (string, error) {
	var phoneNumber string
	err := r.session.Query(
		"SELECT phone_number FROM member_by_verification_id WHERE verification_id = ?",
		verificationId,
	).Scan(&phoneNumber)
	if err != nil {
		slog.Info("fail to find phone_number by verification_id",
			"err", err,
			"verificationId", verificationId,
		)
		return "", err
	}
	return phoneNumber, nil
}

func (r *Repository) SavePhoneNumberMember(phoneNumber string, id *gocql.UUID) error {
	err := r.session.Query(
		"SELECT id FROM member_by_phone_number WHERE phone_number = ?",
		phoneNumber,
	).Scan(id)
	if err != nil && !errors.Is(err, gocql.ErrNotFound) {
		slog.Error("fail to select id from member_by_phone_number",
			"err", err,
			"phoneNumber", phoneNumber,
		)
		return err
	}

	t := time.Now()
	err = r.session.Batch(gocql.LoggedBatch).
		Query(
			"INSERT INTO member_by_phone_number (phone_number_verified, id, phone_number, role, created_time) VALUES (?, ?, ?, ?, ?)",
			true, *id, phoneNumber, constant.ROLE_USER, t,
		).
		Query(
			"INSERT INTO member_by_id (phone_number_verified, id, phone_number, role, created_time) VALUES (?, ?, ?, ?, ?)",
			true, *id, phoneNumber, constant.ROLE_USER, t,
		).Exec()
	if err != nil {
		slog.Error("fail to insert member at member_by_id",
			"err", err,
			"phoneNumber", phoneNumber,
		)
		return err
	}
	return nil
}

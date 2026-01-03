package repository

import (
	"log/slog"

	"github.com/apache/cassandra-gocql-driver/v2"
)

const OTP_TTL = 300

func (r *Repository) SaveVerificationIdAndPhoneNumber(vid gocql.UUID, phoneNumber string) error {
	err := r.session.Query("INSERT INTO phone_number_by_id (phone_number, id) values (?,?) USING TTL ?",
		phoneNumber, vid, OTP_TTL,
	).Exec()
	if err != nil {
		slog.Error("fail to insert phone number with id", err)
		return err
	}
	return nil
}

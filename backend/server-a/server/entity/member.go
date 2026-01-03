package entity

import (
	"time"

	"github.com/apache/cassandra-gocql-driver/v2"
)

type Member struct {
	VerificationId        gocql.UUID `db:"verification_id"`
	OTP                   string     `db:"otp"`
	IsEmailVerified       bool       `db:"is_email_verified"`
	IsPhoneNumberVerified bool       `db:"is_phone_number_verified"`

	Id          gocql.UUID `db:"id"`
	Name        string     `db:"name"`
	PhoneNumber string     `db:"phone_number"`
	Email       string     `db:"email"`
	Password    string     `db:"password"`
	Secret      string     `db:"secret"`
	Role        string     `db:"role"`
	CreatedTime time.Time  `db:"created_time"`
	UpdatedTime time.Time  `db:"updated_time"`
	DeletedTime time.Time  `db:"deleted_time"`
}

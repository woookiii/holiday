package service

import (
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"net/smtp"
	"os"
	"server-a/server/dto"
	"strconv"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	_ "github.com/joho/godotenv/autoload"
)

func (s *Service) IsEmailUsable(req *dto.EmailReq) (bool, error) {
	i, err := s.repository.EmailExists(req.Email)
	if err != nil {
		return false, err
	}
	if i {
		log.Printf("email already exist")
		return false, nil
	}
	return true, nil
}

func (s *Service) SendEmailOTP(req *dto.EmailReq) (*dto.SendOTPResp, error) {
	otp := strconv.Itoa(rand.Intn(900000) + 100000)
	vid := gocql.TimeUUID()
	err := s.repository.SaveEmailAndOtpByVerificationId(vid, req.Email, otp)
	if err != nil {
		return nil, err
	}
	go func() {
		from := os.Getenv("FROM_EMAIL")
		auth := smtp.PlainAuth(
			"",
			from,
			os.Getenv("FROM_EMAIL_PASSWORD"),
			os.Getenv("FROM_EMAIL_SMTP"),
		)

		headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
		message := "Subject: Verify your email\n" + headers + "\n\n" + otp + "\ncode is valid for 5 minutes"

		err = smtp.SendMail(
			os.Getenv("SMTP_ADDR"),
			auth,
			from,
			[]string{req.Email},
			[]byte(message),
		)
		if err != nil {
			log.Printf("fail to send email: %v", err)
		}
	}()

	return &dto.SendOTPResp{VerificationId: vid.String()}, nil
}

func (s *Service) VerifyEmailOTP(req *dto.OTPVerifyReq) (*dto.VerifyEmailOTPResp, error) {
	vid, err := gocql.ParseUUID(req.VerificationId)
	if err != nil {
		slog.Info("fail to parse uuid from verificationId in req", err)
	}
	m, err := s.repository.FindEmailAndOtpByVerificationId(vid)
	if err != nil {
		return nil, err
	}
	if req.OTP != m.OTP {
		log.Printf(
			"code is not same with db code- received code: %v, db code: %v",
			req.OTP, m.OTP,
		)
		return nil, fmt.Errorf("your code %v is not valid", req.OTP)
	}

	err = s.repository.MarkEmailVerified(m.Email)
	if err != nil {
		return nil, err
	}

	sid := gocql.TimeUUID()
	err = s.repository.SaveEmailBySessionId(sid, m.Email)
	if err != nil {
		return nil, err
	}

	resp := dto.VerifyEmailOTPResp{
		SessionId: sid.String(),
	}
	return &resp, nil
}

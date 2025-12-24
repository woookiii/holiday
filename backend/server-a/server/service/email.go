package service

import (
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"server-a/server/dto"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func (s *Service) IsEmailUsable(email string) (bool, error) {
	i, err := s.repository.EmailExists(email)
	if err != nil {
		log.Printf("fail to check email: %v", err)
		return false, err
	}
	if i {
		log.Printf("email already exist")
		return false, nil
	}
	return true, nil
}

func (s *Service) SendEmailVerifyCode(to string) error {
	code := strconv.Itoa(rand.Intn(900000) + 100000)
	err := s.repository.SaveEmailValidationCode(to, code)
	if err != nil {
		log.Printf("fail to save email generation code: %v", err)
		return err
	}
	from := os.Getenv("FROM_EMAIL")
	auth := smtp.PlainAuth(
		"",
		from,
		os.Getenv("FROM_EMAIL_PASSWORD"),
		os.Getenv("FROM_EMAIL_SMTP"),
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	message := "Subject: Verify your email\n" + headers + "\n\n" + code + "\ncode is valid for 5 minutes"

	err = smtp.SendMail(
		os.Getenv("SMTP_ADDR"),
		auth,
		from,
		[]string{to},
		[]byte(message),
	)
	if err != nil {
		log.Printf("fail to send email: %v", err)
		return err
	}
	return nil
}

func (s *Service) ValidateEmailVerifyCode(req *dto.EmailValidateReq) error {

}

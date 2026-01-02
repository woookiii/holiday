package service

import (
	"os"
	"server-a/config"
	"server-a/server/repository"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twilio/twilio-go"
)

type Service struct {
	repository   *repository.Repository
	secretKeyAT  []byte
	secretKeyRT  []byte
	aTExp        int64
	rTExp        int64
	issuer       string
	twilioClient *twilio.RestClient
}

func NewService(cfg *config.Config, r *repository.Repository) *Service {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	apiKey := os.Getenv("TWILIO_API_KEY")
	apiSecret := os.Getenv("TWILIO_API_SECRET")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   apiKey,
		Password:   apiSecret,
		AccountSid: accountSid,
	})

	return &Service{
		repository:   r,
		secretKeyAT:  []byte(os.Getenv("SECRET_KEY_AT")),
		secretKeyRT:  []byte(os.Getenv("SECRET_KEY_RT")),
		aTExp:        cfg.Exp.ATExp,
		rTExp:        cfg.Exp.RtExp,
		issuer:       cfg.Info.Issuer,
		twilioClient: client,
	}
}

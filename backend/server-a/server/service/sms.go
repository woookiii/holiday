package service

import (
	"log/slog"
	"os"
	"server-a/server/dto"

	verify "github.com/twilio/twilio-go/rest/verify/v2"

	_ "github.com/joho/godotenv/autoload"
)

func (s *Service) SendSmsOtp(req *dto.SmsOtpSendReq) error {
	serviceSid := os.Getenv("TWILIO_SERVICE_SID")

	params := &verify.CreateVerificationParams{}
	params.SetTo(req.PhoneNumber)
	params.SetChannel("sms")

	resp, err := s.twilioClient.VerifyV2.CreateVerification(serviceSid, params)
	if err != nil {
		slog.Info("fail to send sms otp code", err)
		return err
	}
	slog.Info("success to send sms otp code", resp.Status)
	return nil
}

package service

import (
	"log/slog"
	"os"
	"server-a/server/dto"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	verify "github.com/twilio/twilio-go/rest/verify/v2"

	_ "github.com/joho/godotenv/autoload"
)

func (s *Service) SendSmsOtp(req *dto.SmsOTPSendReq) (*dto.SendOTPResp, error) {
	serviceSid := os.Getenv("TWILIO_SERVICE_SID")

	params := &verify.CreateVerificationParams{}
	params.SetTo(req.PhoneNumber)
	params.SetChannel("sms")

	resp, err := s.twilioClient.VerifyV2.CreateVerification(serviceSid, params)
	if err != nil {
		slog.Info("fail to send sms otp code", err)
		return nil, err
	}
	slog.Info("success to send sms otp code", resp.To, resp.Status)
	vid := gocql.TimeUUID()
	err = s.repository.SavePhoneNumberByVerificationId(vid, *resp.To)
	if err != nil {
		return nil, err
	}
	res := dto.SendOTPResp{VerificationId: vid.String()}
	return &res, nil
}

func (s *Service) VerifySmsOtp(d *dto.OTPVerifyReq) error {

}

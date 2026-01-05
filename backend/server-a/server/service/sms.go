package service

import (
	"fmt"
	"log/slog"
	"os"
	"server-a/server/constant"
	"server-a/server/dto"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	verify "github.com/twilio/twilio-go/rest/verify/v2"

	_ "github.com/joho/godotenv/autoload"
)

func (s *Service) SendSMSOTP(req *dto.SMSOTPSendReq) (*dto.SendOTPResp, error) {
	serviceSid := os.Getenv("TWILIO_SERVICE_SID")

	params := &verify.CreateVerificationParams{}
	params.SetTo(req.PhoneNumber)
	params.SetChannel("sms")

	resp, err := s.twilioClient.VerifyV2.CreateVerification(serviceSid, params)
	if err != nil {
		slog.Info("fail to send sms otp code", err)
		return nil, err
	}
	if resp.Status != nil {
		slog.Info("success to send sms otp code", "phoneNumber", resp.To, "status", *resp.Status)
	}
	vid := gocql.TimeUUID()
	err = s.repository.SavePhoneNumberByVerificationId(vid, *resp.To)
	if err != nil {
		return nil, err
	}
	res := dto.SendOTPResp{VerificationId: vid.String()}
	return &res, nil
}

func (s *Service) VerifySMSOTP(req *dto.OTPVerifyReq) (*dto.VerifySMSOTPResp, string /*refreshToken*/, error) {
	var sid gocql.UUID
	var email string
	if req.SessionId != nil {
		sid, err := gocql.ParseUUID(*req.SessionId)
		if err != nil {
			slog.Info("fail to parse SessionId from OTPVerifyReq",
				"err", err,
				"SessionId", *req.SessionId,
			)
			return nil, "", err
		}
		email, err := s.repository.FindEmailBySessionId(sid)
		if err != nil {
			return nil, "", err
		}
	}

	vid, err := gocql.ParseUUID(req.VerificationId)
	if err != nil {
		slog.Info("fail to parse verificationId from request",
			"err", err,
			"id", req.VerificationId,
		)
	}
	phoneNumber, err := s.repository.FindPhoneNumberByVerificationId(vid)
	if err != nil {
		return nil, "", err
	}
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(req.OTP)

	serviceSid := os.Getenv("TWILIO_SERVICE_SID")

	resp, err := s.twilioClient.VerifyV2.CreateVerificationCheck(serviceSid, params)
	if err != nil {
		slog.Error("fail to verify phone number otp", "err", err)
		return nil, "", err
	}
	if resp.Status == nil {
		slog.Error("status is nil pointer")
		return nil, "", fmt.Errorf("twilio response status is nil pointer: %v", req.VerificationId)
	}
	if *resp.Status != "approved" {
		slog.Info("otp is not correct", req)
		r := dto.VerifySMSOTPResp{
			PhoneNumberVerified: false,
		}
		return &r, "", nil
	}

	if req.SessionId == nil {
		id := gocql.TimeUUID()
		err = s.repository.SavePhoneNumberMember(phoneNumber, &id)
		if err != nil {
			return nil, "", err
		}

		at, err := createToken(id.String(), constant.ROLE_USER, s.secretKeyAT, constant.ACCESS_TOKEN_TTL)
		if err != nil {
			slog.Error("fail to create access token",
				"err", err,
				"id", id.String(),
			)
			return nil, "", err
		}
		rt, err := createToken(id.String(), constant.ROLE_USER, s.secretKeyRT, constant.REFRESH_TOKEN_TTL)
		r := dto.VerifySMSOTPResp{
			PhoneNumberVerified: true,
			AccessToken:         at,
		}
		return &r, rt, nil
	}

	//TODO: save member who verified email and phone number and give qr secret to them
	//since they are not oauth2 user, and this action need to manage oauth user's email
	//and non-ouath user's email as different
	s.repository.SaveEmail

	return nil
}

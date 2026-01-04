package dto

type SendOTPResp struct {
	VerificationId string `json:"verificationId"`
}
type OTPVerifyReq struct {
	SessionId      string `json:"sessionId"`
	VerificationId string `json:"verificationId"`
	OTP            string `json:"otp"`
}

type SmsOTPSendReq struct {
	PhoneNumber string `json:"phoneNumber"`
}

type VerifyEmailOTPResp struct {
	SessionId string `json:"session_id"`
}

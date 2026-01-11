package dto

type SendOTPResp struct {
	VerificationId string `json:"verificationId"`
}

type EmailOTPVerifyReq struct {
	VerificationId string `json:"verificationId"`
	OTP            string `json:"otp"`
}

type SMSOTPVerifyReq struct {
	SessionId      *string `json:"sessionId"` //nullable
	VerificationId string  `json:"verificationId"`
	OTP            string  `json:"otp"`
}

type SMSOTPSendReq struct {
	PhoneNumber string `json:"phoneNumber"`
}

type VerifyEmailOTPResp struct {
	EmailVerified bool   `json:"emailVerified"`
	SessionId     string `json:"session_id"`
}

type VerifySMSOTPResp struct {
	PhoneNumberVerified bool   `json:"phoneNumberVerified"`
	AccessToken         string `json:"accessToken"`
}

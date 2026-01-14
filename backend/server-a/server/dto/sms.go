package dto

type SMSOTPSendReq struct {
	PhoneNumber string `json:"phoneNumber"`
}
type SMSOTPVerifyReq struct {
	SessionId      *string `json:"sessionId"` //nullable
	VerificationId string  `json:"verificationId"`
	OTP            string  `json:"otp"`
}

type SMSOTPVerifyResp struct {
	PhoneNumberVerified bool   `json:"phoneNumberVerified"`
	AccessToken         string `json:"accessToken"`
}

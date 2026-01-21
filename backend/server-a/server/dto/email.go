package dto

type EmailReq struct {
	Email string `json:"email"`
}
type EmailMemberSaveReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailMemberLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailLoginResp struct {
	PhoneNumberVerified bool   `json:"phoneNumberVerified"`
	EmailVerified       bool   `json:"emailVerified"`
	Id                  string `json:"id"`
	SessionId           string `json:"sessionId"`
	AccessToken         string `json:"accessToken"`
}

type EmailOTPSendReq struct {
	Id string `json:"id"`
}

type EmailOTPVerifyReq struct {
	VerificationId string `json:"verificationId"`
	OTP            string `json:"otp"`
}
type EmailOTPVerifyResp struct {
	EmailVerified bool   `json:"emailVerified"`
	SessionId     string `json:"session_id"`
}

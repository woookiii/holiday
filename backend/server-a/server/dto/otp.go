package dto

type OtpResp struct {
	VerificationId string `json:"verificationId"`
}
type EmailOtpVerifyReq struct {
	VerificationId string `json:"verificationId"`
	OTP            string `json:"otp"`
}

type SmsOtpSendReq struct {
	PhoneNumber string `json:"phoneNumber"`
}

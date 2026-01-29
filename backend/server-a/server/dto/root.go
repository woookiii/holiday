package dto

type OTPSendResponse struct {
	VerificationId string `json:"verificationId"`
}

type TokenRefreshResponse struct {
	AccessToken string `json:"accessToken"`
}

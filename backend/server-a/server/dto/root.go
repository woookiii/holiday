package dto

type OTPSendResp struct {
	VerificationId string `json:"verificationId"`
}

type TokenRefreshResp struct {
	AccessToken string `json:"accessToken"`
}

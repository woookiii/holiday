package dto

type EmailReq struct {
	Email string `json:"email"`
}
type MemberSaveReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MemberLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailLoginResp struct {
	IsPhoneNumberVerified bool   `json:"isPhoneNumberVerified,omitempty"`
	IsEmailVerified       bool   `json:"isEmailVerified,omitempty"`
	SessionId             string `json:"session_id,omitempty"`
	AccessToken           string `json:"accessToken,omitempty"`
}

type TokenRefreshResp struct {
	AccessToken string `json:"accessToken"`
}

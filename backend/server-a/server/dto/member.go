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
	PhoneNumberVerified bool   `json:"phoneNumberVerified"`
	EmailVerified       bool   `json:"emailVerified"`
	SessionId           string `json:"session_id"`
	AccessToken         string `json:"accessToken"`
}

type TokenRefreshResp struct {
	AccessToken string `json:"accessToken"`
}

package dto

type MemberSaveReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MemberLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string
}

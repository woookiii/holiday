package dto

type EmailValidateReq struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

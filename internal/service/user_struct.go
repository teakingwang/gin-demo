package service

type CreateUser struct {
	Mobile     string `json:"mobile"`
	VerifyCode string `json:"verify_code"`
}

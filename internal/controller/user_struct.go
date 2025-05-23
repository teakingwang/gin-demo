package controller

type CreateUser struct {
	Mobile     string `json:"mobile" binding:"required,len=11"`
	VerifyCode string `json:"verify_code" binding:"required,len=4"`
}

type SendSms struct {
	Mobile string `json:"mobile" binding:"required,len=11"`
}

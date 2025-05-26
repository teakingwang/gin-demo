package controller

type CreateUserReq struct {
	Mobile     string `json:"mobile" binding:"required,len=11" validate:"mobile"`
	VerifyCode string `json:"verify_code" binding:"required,len=6"`
}

type CreateUserResp struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
	HasPwd bool   `json:"has_pwd"`
}

type SendSms struct {
	Mobile string `json:"mobile" binding:"required,len=11"  validate:"mobile"`
}

type SetPwd struct {
	Pwd string `json:"pwd" binding:"required,min=6,max=20"`
}

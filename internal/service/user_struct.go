package service

type CreateUser struct {
	Mobile     string
	VerifyCode string
}

type UserItem struct {
	UserID int64
	HasPwd bool
}

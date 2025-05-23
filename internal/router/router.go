package router

import (
	"github.com/gin-gonic/gin"
	"github.com/teakingwang/gin-demo/internal/app"
	"github.com/teakingwang/gin-demo/internal/controller"
)

func NewRouter(ctx *app.AppContext) *gin.Engine {
	r := gin.Default()

	userV1 := r.Group("/v1/user")
	{
		userController := controller.NewUserController(ctx)
		userV1.GET("/list", userController.GetUserList)
		userV1.POST("/register", userController.RegisterUser)
		userV1.POST("/sendsms", userController.SendSms)
	}

	return r
}

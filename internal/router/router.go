package router

import (
	"github.com/gin-gonic/gin"
	"github.com/teakingwang/gin-demo/internal/app"
	"github.com/teakingwang/gin-demo/internal/controller"
	"github.com/teakingwang/gin-demo/pkg/middleware"
)

func NewRouter(ctx *app.AppContext) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	registerUserRoutes(ctx, v1)

	return r
}

func registerUserRoutes(ctx *app.AppContext, v1 *gin.RouterGroup) {
	c := controller.NewUserController(ctx)
	user := v1.Group("/user")
	{
		user.POST("/loginsms", c.LoginSms)
		user.POST("/sendsms", c.SendSms)
	}

	authUser := v1.Group("/user").Use(middleware.JWTAuthMiddleware())
	{
		authUser.GET("/list", c.GetUserList)
	}
}

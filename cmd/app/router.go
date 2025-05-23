package app

import (
	"github.com/gin-gonic/gin"
	"github.com/teakingwang/gin-demo/internal/app"
	"github.com/teakingwang/gin-demo/internal/controller"
)

type Router struct {
	addr   string
	router *gin.Engine
}

func NewRouter(addr string) *Router {
	return &Router{
		addr:   addr,
		router: gin.Default(),
	}
}

func (r *Router) Config(ctx *app.AppContext) {
	r.router.MaxMultipartMemory = 8 << 20 // 8 MiB

	userV1 := r.router.Group("/v1/user")
	{
		userController := controller.NewUserController(ctx)
		userV1.GET("/list", userController.GetUserList)
		userV1.POST("/register", userController.RegisterUser)
		userV1.POST("/sendsms", userController.SendSms)
	}
}

func (r *Router) Run() {
	go r.router.Run(r.addr)
}

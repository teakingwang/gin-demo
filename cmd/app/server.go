package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/internal/app"
	"github.com/teakingwang/gin-demo/internal/router"
	"github.com/teakingwang/gin-demo/pkg/idgen"
	"net"
)

type Server struct{}

var server = newServer()

func newServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	// load config
	config.LoadConfig()
	// ctx
	ctx := app.NewAppContext()
	// idgen
	// 初始化 ID 生成器
	if err := idgen.Init(); err != nil {
		panic(fmt.Sprintf("failed to initialize idgen: %v", err))
	}

	// router
	r := router.NewRouter(ctx)
	go func() {
		if err := r.Run(net.JoinHostPort(config.Config.Server.Host, config.Config.Server.Port)); err != nil {
			ctx.Logger.Panic(fmt.Sprintf("failed to run Gin server: %v", err))
		}
	}()

	select {}
}

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "server",
		Long:         `The server is gin-demo demo`,
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("arg:", args)
			server.Run()
		},
	}

	cmd.Flags().StringP("config", "c", "config.yaml", "config file (default is ./resources/config.yaml)")
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		panic(err)
	}

	return cmd
}

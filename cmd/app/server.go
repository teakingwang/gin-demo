package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/pkg/db"
	"net"
)

type Server struct {
	router *Router
}

var server = newServer()

func newServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	// load config
	config.LoadConfig()
	// 初始化db
	gormDB, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	// 数据库迁移
	db.MigrateDB(gormDB)

	// router
	s.router = NewRouter(net.JoinHostPort(config.Config.Server.Host, config.Config.Server.Port))
	s.router.Config()
	s.router.Run()

	select {}
}

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "server",
		Long:         `The server is gin-demo demo`,
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			server.Run()
		},
		Args: cobra.ExactArgs(1),
	}

	cmd.Flags().StringP("config", "c", "config.yaml", "config file (default is ./resources/config.yaml)")
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		panic(err)
	}

	return cmd
}

package app

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teakingwang/gin-demo/cmd/cron"
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/internal/app"
	"github.com/teakingwang/gin-demo/internal/router"
	"github.com/teakingwang/gin-demo/pkg/idgen"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	cron *cron.CronServer
}

func newServer() *Server {
	cronServer := cron.NewCronServer()
	return &Server{
		cron: cronServer,
	}
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

	// 注册定时任务
	s.cron.RegisterTasks(ctx)
	s.cron.Start()
	defer s.cron.Stop() // 确保退出时停止任务

	// router
	r := router.NewRouter(ctx)
	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:    net.JoinHostPort(config.Config.Server.Host, config.Config.Server.Port),
		Handler: r,
	}

	go func() {
		if err := r.Run(srv.Addr); err != nil {
			ctx.Logger.Panic(fmt.Sprintf("failed to run Gin server: %v", err))
		}
	}()

	// 优雅退出：监听系统中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown signal received.")

	// 优雅关闭 HTTP 服务
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		ctx.Logger.Errorf("HTTP server shutdown error: %v", err)
	} else {
		ctx.Logger.Info("HTTP server shutdown complete.")
	}
}

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "server",
		Long:         `The server is gin-demo demo`,
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("arg:", args)
			s := newServer()
			s.Run()
		},
	}

	cmd.Flags().StringP("config", "c", "config.yaml", "config file (default is ./resources/config.yaml)")
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		panic(err)
	}

	return cmd
}

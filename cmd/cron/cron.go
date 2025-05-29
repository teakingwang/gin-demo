// app/cron.go
package cron

import (
	"context"
	"github.com/teakingwang/gin-demo/internal/app"
	"log"

	"github.com/robfig/cron/v3"
)

type CronServer struct {
	cron *cron.Cron
}

func NewCronServer() *CronServer {
	return &CronServer{
		cron: cron.New(cron.WithSeconds()), // 支持秒级调度
	}
}

// 注册定时任务
func (s *CronServer) RegisterTasks(ctx *app.AppContext) {
	// 示例任务：每30秒执行一次
	_, err := s.cron.AddFunc("*/30 * * * * *", func() {
		err := ctx.UserService.DoCleanupTask(context.Background())
		if err != nil {
			ctx.Logger.Error(err.Error())
			return
		}
	})
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
}

func (s *CronServer) Start() {
	log.Println("启动定时任务调度器...")
	s.cron.Start()
}

func (s *CronServer) Stop() {
	log.Println("停止定时任务调度器...")
	s.cron.Stop()
}

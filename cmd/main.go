package main

import (
	"github.com/sirupsen/logrus"
	"github.com/teakingwang/gin-demo/cmd/app"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	command := app.NewServerCommand()

	if err := command.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

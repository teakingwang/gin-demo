package main

import (
	"github.com/sirupsen/logrus"
	"github.com/teakingwang/gin-demo/cmd/app"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	command := app.NewServerCommand()
	if err := command.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

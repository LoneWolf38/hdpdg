package main

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
)

var log hclog.Logger

func InitLogger() {
	log = hclog.New(&hclog.LoggerOptions{
		Name:  "my-app",
		Level: hclog.LevelFromString("DEBUG"),
	})
}

func kafkaLogger(msg string, a ...interface{}) {
	log.Info(fmt.Sprintf(msg, a...))
}

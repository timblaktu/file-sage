package main

import (
	"fmt"
	"log"

	"github.com/timblaktu/wupdedup/logging"
	"golang.org/x/exp/slog"
)

func init() {
	logging.Init(slog.LevelInfo)
	slog.Debug("init: logging initialized")
	slog.Debug("init exiting..")
}

func main() {
	slog.Debug("main entered")

	c := loadConfig()
	contexts := loadStorageStrategyContexts(&c)
	slog.Debug(fmt.Sprintf("%v", contexts))
	for _, context := range contexts {
		slog.Debug(fmt.Sprintf("%v", context))
		context.scanTree()
	}

	log.Printf("main exiting..")
}

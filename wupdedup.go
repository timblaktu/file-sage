package main

import (
	"log"

	"github.com/timblaktu/wupdedup/db"
	"github.com/timblaktu/wupdedup/logging"
	"golang.org/x/exp/slog"
)

const (
	dbfile string = "wupdedup.bolt.db"
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

	d, err := db.Open(dbfile)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	for _, context := range contexts {
		b, err := d.Bucket([]byte(context.name))
		if err != nil {
			log.Fatal(err)
		}
		context.SetBucket(b)
	}

	for _, context := range contexts {
		context.scanTree()
	}

	slog.Debug("main exiting..")
}

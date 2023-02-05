package main

import (
	"log"

	"github.com/timblaktu/wupdedup/config"
	"github.com/timblaktu/wupdedup/db"
	"github.com/timblaktu/wupdedup/logging"
	"github.com/timblaktu/wupdedup/profiler"
	"golang.org/x/exp/slog"
)

const (
	dbfile string = "wupdedup.bolt.db"
)

func init() {
	slog.Debug("init: logging initialized")
	slog.Debug("init exiting..")
}

func main() {
	slog.Debug("main entered")

	c := config.LoadConfig()

	logging.Init(c.LogLevel)

	var p *profiler.Profiler
	if c.Profile.Specified() {
		p = profiler.New(c.Profile)
		p.Start()
	}

	contexts := loadStorageStrategyContexts(&c)

	d := db.Init(c.DBFile)
	defer d.Close()
	for _, context := range contexts {
		b, err := d.Bucket([]byte(context.name))
		if err != nil {
			log.Fatal(err)
		}
		context.SetBucket(b)
		context.scanTree()
	}

	if p != nil {
		p.Stop()
	}

	slog.Debug("main exiting..")
}

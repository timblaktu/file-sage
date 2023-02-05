package main

import (
	"github.com/timblaktu/wupdedup/config"

	"golang.org/x/exp/slog"
)

// Concrete type that implements StorageStrategy interface for SmugMug
type SmugmugStrategy struct {
	conf config.SmugMugConfig
}

func (s SmugmugStrategy) scanTree(c *StorageStrategyContext) {
	slog.Info("scanning Smugmug account", "url", s.conf.URL)
}

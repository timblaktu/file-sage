package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"golang.org/x/exp/slog"
)

// Concrete type that implements StorageStrategy interface for Local Storage
type LocalStrategy struct {
	conf LocalConfig
}

func (s LocalStrategy) scanTree(c *StorageStrategyContext) {
	slog.Info("scanning local tree..", "path", s.conf.RootPath)
	filepath.WalkDir(s.conf.RootPath, walkdirFunc)
	slog.Info("Done scanning local tree", "path", s.conf.RootPath)
}

func walkdirFunc(path string, d fs.DirEntry, err error) error {
	//slog.Info(fmt.Sprintf("walkdirFunc: visiting path %s, DirEntry %v, err %s\n", path, d, err))
	slog.Info(fmt.Sprintf("%s", path))
	return nil
}

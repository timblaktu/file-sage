package main

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/timblaktu/wupdedup/content"
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
	if d.IsDir() {
		slog.Debug("ignoring dir entry")
		return nil
	}
	ft, err := getType(path)
	if err != nil {
		return err
	}
	slog.Debug("visiting", "path", path, "ft", ft)
	// TODO: customize walkdirfunc to allow passing in context object which contains bucket obj
	// b.PutLocal()
	return nil
}

func getType(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("cannot open file", err, "path", path)
		return "", err
	}
	defer f.Close()
	buf := make([]byte, 512)
	_, err = f.Read(buf)
	if err != nil {
		slog.Error("cannot read 512 byte header from file", err, "path", path)
		return "", err
	}
	ft := content.GetType(buf)
	return ft, nil
}

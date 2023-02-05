package main

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/timblaktu/wupdedup/config"
	"github.com/timblaktu/wupdedup/content"
	"golang.org/x/exp/slog"
)

// Concrete type that implements StorageStrategy interface for Local Storage
type LocalStrategy struct {
	conf config.LocalConfig
}

// var wg sync.WaitGroup

func (s LocalStrategy) scanTree(c *StorageStrategyContext) {
	slog.Info("scanning local tree..", "path", s.conf.RootPath)
	// fileSystem := os.DirFS(s.conf.RootPath)
	// fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
	filepath.WalkDir(s.conf.RootPath, func(path string, d fs.DirEntry, err error) error {
		fi, _ := d.Info()
		slog.Debug("visiting", "path", path, "de.name", d.Name(),
			"de.isdir", d.IsDir(), "de.type", d.Type(), "fi.name", fi.Name(),
			"fi.size", fi.Size(), "fi.mode", fi.Mode().String(),
			"fi.modtime", fi.ModTime(), "fi.isdir", fi.IsDir(), "fi.sys", fi.Sys())
		c.nodeCount++
		// defer wg.Done()
		if d.IsDir() {
			slog.Debug("ignoring bc isdir")
			return nil
		}
		c.fileCount++
		fullPath := path
		// fullPath := filepath.Join(d.Name(), path)
		ft, err := getType(fullPath)
		if err != nil {
			return err
		}
		slog.Debug("visited", "fullPath", fullPath, "filetype", ft, "#nodes",
			c.nodeCount, "#files", c.fileCount, "DirEntry", d)
		return nil
	})
	slog.Info("Done scanning local tree", "path", s.conf.RootPath,
		"#nodes", c.nodeCount, "#files", c.fileCount)
}

// TODO: optimize this - faster way to get file mime type?
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

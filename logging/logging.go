package logging

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

func Init(lvl slog.Level) {
	// https://pkg.go.dev/golang.org/x/exp/slog#HandlerOptions
	opts := slog.HandlerOptions{
		AddSource: false,
		Level:     lvl,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// discard time attrs from log records
			// if a.Key == slog.TimeKey {
			// 	return slog.Attr{}
			// }
			return a
		},
	}
	th := opts.NewTextHandler(os.Stdout)

	// "bind" constant attr key/val to all log records handled
	// th := th.WithAttrs([]slog.Attr{slog.String("version", "v0.0.1-beta")})

	logger := slog.New(th).WithContext(context.Background())
	slog.SetDefault(logger)
}

package slogtest

import (
	"context"
	"io"
	"log/slog"

	"github.com/wlynch/slogctx"
)

var (
	_ io.Writer = &logAdapter{}
)

type logAdapter struct {
	l Logger
}

func (l *logAdapter) Write(b []byte) (int, error) {
	l.l.Log(string(b))
	return len(b), nil
}

type Logger interface {
	Log(args ...any)
	Logf(format string, args ...any)
}

// TestLogger gets a logger to use in unit and end to end tests
func TestLogger(t Logger) *slogctx.Logger {
	return slogctx.New(slog.NewTextHandler(&logAdapter{l: t}, nil))
}

// TestContextWithLogger returns a context with a logger to be used in tests
func TestContextWithLogger(t Logger) context.Context {
	return slogctx.WithLogger(context.Background(), TestLogger(t))
}

// RemoveTime removes the top-level time attribute.
// It is intended to be used as a ReplaceAttr function,
// to make example output deterministic.
//
// This is taken from slog/internal/slogtest.RemoveTime.
func RemoveTime(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	return a
}

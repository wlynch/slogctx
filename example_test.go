package slogctx_test

import (
	"context"
	"log/slog"
	"os"

	"github.com/chainguard-dev/slogctx"
	"github.com/chainguard-dev/slogctx/slogtest"
)

func ExampleHandler() {
	log := slog.New(slogctx.NewHandler(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// Remove time for repeatable results
		ReplaceAttr: slogtest.RemoveTime,
	})))

	ctx := context.Background()
	ctx = slogctx.WithValues(ctx, "foo", "bar")
	log.InfoContext(ctx, "hello world", slog.Bool("baz", true))

	// Output:
	// level=INFO msg="hello world" baz=true foo=bar
}

func ExampleLogger() {
	log := slogctx.NewLogger(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// Remove time for repeatable results
		ReplaceAttr: slogtest.RemoveTime,
	})))
	log = log.With("a", "b")
	ctx := slogctx.WithLogger(context.Background(), log)

	// Grab logger from context and use
	// Note: this is a formatter aware method, not an slog.Attr method.
	slogctx.FromContext(ctx).With("foo", "bar").Infof("hello %s", "world")

	// Package level context loggers are also aware
	slogctx.ErrorContext(ctx, "asdf", slog.Bool("baz", true))

	// Output:
	// level=INFO msg="hello world" a=b foo=bar
	// level=ERROR msg=asdf a=b baz=true
}

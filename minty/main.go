package main

import (
	"context"
	"github.com/runabol/tork"
	"github.com/runabol/tork/engine"
	"github.com/runabol/tork/middleware/task"
	"os"
)

func main() {
	engine.SetMode(engine.ModeStandalone)
	mw := func(next task.HandlerFunc) task.HandlerFunc {
		return func(ctx context.Context, e task.EventType, t *tork.Task) error {
			pass := os.Getenv("MINTY_DEVICE_PASSWORD")
			if pass != "" {
				if t.Env == nil {
					t.Env = make(map[string]string)
				}
				t.Env["WARP_DEVICE_PASSWORD"] = pass
			}
			return next(ctx, e, t)
		}

	}

	engine.RegisterTaskMiddleware(mw)
	engine.Run()
}

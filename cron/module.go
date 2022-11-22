package cron

import (
	"context"
	"time"
)

type Module struct{}

func New() *Module {
	return &Module{}
}

func (m Module) RunPeriodically(ctx context.Context, task func(), interval time.Duration) *Worker {
	return runPeriodically(ctx, task, interval)
}

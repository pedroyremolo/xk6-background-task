package cron

import "time"

type Module struct{}

func New() *Module {
	return &Module{}
}

// RunPeriodically receives a task, as function, which runs on a interval passed in milliseconds
func (m Module) RunPeriodically(task func(), interval int) *Worker {
	return runPeriodically(task, time.Duration(interval)*time.Millisecond)
}

package cron

import (
	"context"
	"time"
)

type Worker struct {
	ticker *time.Ticker

	stop chan bool
}

func runPeriodically(_ context.Context, task func(), interval time.Duration) *Worker {
	stop := make(chan bool)
	w := &Worker{ticker: time.NewTicker(interval), stop: stop}
	go func() {
		for {
			select {
			case <-stop:
				close(w.stop)
				w.ticker.Stop()
				return
			case <-w.ticker.C:
				task()
			}
		}
	}()

	return w
}

func (w *Worker) Stop() {
	w.stop <- true
}

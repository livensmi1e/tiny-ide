package cee

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/livensmi1e/tiny-ide/pkg/logger"
	"github.com/livensmi1e/tiny-ide/queue"
	"github.com/livensmi1e/tiny-ide/store"
)

type WorkerPool struct {
	Sandbox    Sandbox
	Queue      queue.Queue
	Store      *store.Store
	PollDelay  time.Duration
	Logger     logger.Logger
	NumWorkers int
}

func NewWorkerPool(s *store.Store, q queue.Queue, l logger.Logger, pd time.Duration, n int) *WorkerPool {
	return &WorkerPool{
		Sandbox:    NewDockerContainer("sandbox", 24*time.Hour),
		Queue:      q,
		Store:      s,
		Logger:     l,
		PollDelay:  pd,
		NumWorkers: n,
	}
}

func (w *WorkerPool) Start(ctx context.Context) {
	var wg sync.WaitGroup
	for i := 0; i < w.NumWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			w.Logger.Info(fmt.Sprintf("worker %d started", workerID))
			w.Run(ctx, workerID)
		}(i)
	}
}

func (w *WorkerPool) Run(ctx context.Context, workerID int) {
	for {
		select {
		case <-ctx.Done():
			w.Logger.Info(fmt.Sprintf("worker %d stopped", workerID))
			return
		default:
			submission, err := w.Queue.Pop()
			if err != nil {
				// w.Logger.Info(fmt.Sprintf("worker %d: queue empty, retrying", workerID))
				time.Sleep(w.PollDelay)
				continue
			}
			output, err := w.Sandbox.Run(submission)
			status := "success"
			if err != nil {
				status = "fail"
			}
			_, err = w.Store.UpdateSubmission(&store.UpdateSubmission{
				ID:         submission.ID,
				LanguageID: submission.LanguageID,
				Status:     &status,
				Stdout:     &output.Stdout,
				Stderr:     &output.Stderr,
				Time:       &output.Time,
				Memory:     &output.Memory,
			})
			if err != nil {
				w.Logger.Info(fmt.Sprintf("worker %d: failed to update submission", workerID))
			}
		}

	}
}

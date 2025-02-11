package cee

import (
	"context"
	"time"

	"github.com/livensmi1e/tiny-ide/pkg/logger"
	"github.com/livensmi1e/tiny-ide/queue"
	"github.com/livensmi1e/tiny-ide/store"
)

type Worker struct {
	Sandbox   Sandbox
	Queue     queue.Queue
	Store     *store.Store
	PollDelay time.Duration
	Logger    logger.Logger
}

func NewWorker(s *store.Store, q queue.Queue, l logger.Logger, pd time.Duration) *Worker {
	return &Worker{
		Sandbox:   NewDockerContainer("sandbox", 24*time.Hour),
		Queue:     q,
		Store:     s,
		Logger:    l,
		PollDelay: pd,
	}
}

func (w *Worker) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			w.Logger.Info("worker stopped")
			return
		default:
			submission, err := w.Queue.Pop()
			if err != nil {
				w.Logger.Info("queue empty, retrying")
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
				w.Logger.Debug("fail to update session")
			}
		}

	}
}

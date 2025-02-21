package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "github.com/livensmi1e/tiny-ide/executor/proto"
	"github.com/livensmi1e/tiny-ide/pkg/logger"
	"github.com/livensmi1e/tiny-ide/queue"
	"github.com/livensmi1e/tiny-ide/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WorkerPool struct {
	Queue      queue.Queue
	Store      *store.Store
	PollDelay  time.Duration
	Logger     logger.Logger
	NumWorkers int
	GrpcClient pb.ExecutorClient
}

func NewWorkerPool(s *store.Store, q queue.Queue, l logger.Logger, pd time.Duration, n int, addr string) *WorkerPool {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to executor service: %v", err))
	}
	client := pb.NewExecutorClient(conn)
	return &WorkerPool{
		Queue:      q,
		Store:      s,
		Logger:     l,
		PollDelay:  pd,
		NumWorkers: n,
		GrpcClient: client,
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
	wg.Wait()
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
				time.Sleep(w.PollDelay)
				continue
			}
			req := &pb.ExecuteRequest{
				SubmissionId: submission.ID,
				LanguageId:   submission.LanguageID,
				SourceCode:   submission.SourceCode,
			}
			resp, err := w.GrpcClient.Execute(ctx, req)
			if err != nil {
				w.Logger.Info(fmt.Sprintf("worker %d: failed to execute submission", workerID))
				w.Logger.Error(err.Error())
				continue
			}
			_, err = w.Store.UpdateSubmission(&store.UpdateSubmission{
				ID:         submission.ID,
				LanguageID: submission.LanguageID,
				Status:     &resp.Status,
				Stdout:     &resp.Stdout,
				Stderr:     &resp.Stderr,
				Time:       &resp.Time,
				Memory:     &resp.Memory,
			})
			if err != nil {
				w.Logger.Info(fmt.Sprintf("worker %d: failed to update submission", workerID))
			}
		}

	}
}

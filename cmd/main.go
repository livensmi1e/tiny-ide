package main

import (
	"context"
	"fmt"
	"time"

	"github.com/livensmi1e/tiny-ide/infra"
	"github.com/livensmi1e/tiny-ide/pkg/config"
	"github.com/livensmi1e/tiny-ide/pkg/logger"
	"github.com/livensmi1e/tiny-ide/pkg/validator"
	"github.com/livensmi1e/tiny-ide/queue"
	"github.com/livensmi1e/tiny-ide/server"
	"github.com/livensmi1e/tiny-ide/store"
	"github.com/livensmi1e/tiny-ide/store/db"
	"github.com/livensmi1e/tiny-ide/worker"
)

func main() {
	cfg := config.New()
	logger := logger.New()
	driver, err := db.New(cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
	store := store.New(driver)
	if err := store.Migrate(); err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("migrate completed")
	}
	validator := validator.New()
	queue := queue.New(cfg, "submissions")

	workerPool := worker.NewWorkerPool(store, queue, *logger, 1*time.Second, 5, cfg.Executor.Addr)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go workerPool.Start(ctx)

	infra := infra.NewInfrastructure(cfg, logger, store, validator, queue)
	server := server.NewServer(infra)
	server.Start()
}

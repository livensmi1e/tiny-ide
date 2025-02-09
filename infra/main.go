package infra

import (
	"github.com/livensmi1e/tiny-ide/pkg/config"
	"github.com/livensmi1e/tiny-ide/pkg/logger"
	"github.com/livensmi1e/tiny-ide/pkg/validator"
	"github.com/livensmi1e/tiny-ide/queue"
	"github.com/livensmi1e/tiny-ide/store"
)

type Infrastructure interface {
	Config() *config.Config
	Logger() *logger.Logger
	Validator() validator.Validator
	Queue() *queue.SubmissionQueue
	Store() *store.Store
}

type AppInfra struct {
	config    *config.Config
	logger    *logger.Logger
	validator validator.Validator
	queue     *queue.SubmissionQueue
	store     *store.Store
}

func NewInfrastructure(cfg *config.Config, lg *logger.Logger, st *store.Store, vldt validator.Validator, que *queue.SubmissionQueue) Infrastructure {
	return &AppInfra{
		config:    cfg,
		logger:    lg,
		store:     st,
		validator: vldt,
		queue:     que,
	}
}

func (i *AppInfra) Config() *config.Config {
	return i.config
}

func (i *AppInfra) Logger() *logger.Logger {
	return i.logger
}

func (i *AppInfra) Store() *store.Store {
	return i.store
}

func (i *AppInfra) Validator() validator.Validator {
	return i.validator
}

func (i *AppInfra) Queue() *queue.SubmissionQueue {
	return i.queue
}

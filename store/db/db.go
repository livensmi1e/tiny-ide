package db

import (
	"errors"
	"fmt"

	"github.com/livensmi1e/tiny-ide/pkg/config"
	"github.com/livensmi1e/tiny-ide/store"
	"github.com/livensmi1e/tiny-ide/store/db/postgres"
)

func New(cfg *config.Config) (store.Driver, error) {
	var driver store.Driver
	var err error

	switch cfg.Store.Driver {
	case "postgres":
		driver, err = postgres.New(cfg)
	default:
		return nil, errors.New("unknown db driver")
	}
	if err != nil {
		return nil, fmt.Errorf("fail to create db driver: %w", err)
	}
	return driver, nil
}

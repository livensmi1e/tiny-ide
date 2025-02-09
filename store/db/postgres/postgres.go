package postgres

import "github.com/livensmi1e/tiny-ide/pkg/config"

type postgres struct{}

func New(cfg *config.Config) (*postgres, error) {
	return &postgres{}, nil
}

package env

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/romanfomindev/microservices-chat-server/internal/config"
)

const (
	dsnEnvName = "PG_DSN"
	timeout    = "PG_TIMEOUT"
)

var _ config.PGConfig = (*pgConfig)(nil)

type pgConfig struct {
	dsn     string
	timeout time.Duration
}

func NewPGConfig() (config.PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	pgTimeout := os.Getenv(timeout)
	if len(dsn) == 0 {
		return nil, errors.New("pg  not found")
	}

	timeout, err := strconv.Atoi(pgTimeout)
	if err != nil {
		timeout = 30
	}

	return &pgConfig{
		dsn:     dsn,
		timeout: time.Duration(timeout),
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}

func (cfg *pgConfig) Timeout() time.Duration {
	return cfg.timeout
}

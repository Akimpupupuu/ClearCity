package db

import (
	"context"
	"fmt"

	"github.com/Akimpupupuu/ClearCity/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(cfg *config.Config) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DB)

	pgxcfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	pgxcfg.MaxConns = 20

	return pgxpool.NewWithConfig(context.Background(), pgxcfg)
}

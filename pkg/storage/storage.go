package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DB struct {
	*pgxpool.Pool
}

// New creates new connections pool and runs migrations.
func New() (*DB, error) {
	name := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	mode := os.Getenv("POSTGRES_MODE")
	port := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", name, password, host, port, dbName, mode)

	m, err := migrate.New("file://pkg/storage/migrations", connStr)
	if err != nil {
		return nil, err
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return nil, err
		}
	}

	pgxCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Err(err).Msg("failed to parse pgx config")
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxCfg)
	if err != nil {
		log.Err(err).Msg("failed to create pgxpool")
		return nil, err
	}

	if err = pool.Ping(context.Background()); err != nil {
		log.Err(err).Msg("failed to ping database")
		return nil, err
	}

	return &DB{pool}, nil
}

package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Pool       *pgxpool.Pool
	once       sync.Once
	dbpool_err error
)

func Init() error {
	once.Do(func() {
		// Connect to database.
		Pool, dbpool_err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	})

	if dbpool_err != nil {
		return fmt.Errorf("unable to create connection pool: %v", dbpool_err)
	}

	return nil
}

func Close() {
	Pool.Close()
}

package utilslibs

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pgPoolInstance *pgxpool.Pool
	pgOnce         sync.Once
)

// InitializePostgresPool creates (or returns) a singleton *pgxpool.Pool connection.
// This function uses the sync.Once pattern to ensure the pool is only set up once.
//
// Example usage:
//
//	cfg := utilslibs.PostgresConfig{
//	    DBUser:  "username",
//	    DBPass:  "password",
//	    DBHost:  "localhost",
//	    DBPort:  "5432",
//	    DBName:  "database",
//	    SSLMode: "disable",
//	}
//	pool, err := utilslibs.InitializePostgresPool(context.Background(), cfg)
//	if err != nil {
//	    log.Fatalf("Failed to initialize Postgres pool: %v", err)
//	}
func InitializePostgresPool(ctx context.Context, config PostgresConfig) *pgxpool.Pool {
	var err error

	pgOnce.Do(func() {
		// Build the connection string
		databaseURL := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			config.DBUser,
			config.DBPass,
			config.DBHost,
			config.DBPort,
			config.DBName,
			config.SSLMode,
		)

		// Create a new connection pool
		pgPoolInstance, err = pgxpool.New(ctx, databaseURL)
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to Postgres: %v", err))
		}
	})

	// If pgOnce set up the pool successfully, poolInstance won't be nil (unless there was an error).
	return pgPoolInstance
}

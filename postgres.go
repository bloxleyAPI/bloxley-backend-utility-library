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
//	    DBTimezone: "UTC",
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
			config.DB_USER,
			config.DB_PASSWORD,
			config.DB_HOST,
			config.DB_PORT,
			config.DB_NAME,
			config.DB_SSLMODE,
		)

		// Create a new connection pool
		pgPoolInstance, err = pgxpool.New(ctx, databaseURL)
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to Postgres: %v", err))
		} else {
			fmt.Println("Connected to Postgres")
		}

		// Ping the database to ensure the connection is alive
		if err := PingDB(ctx); err != nil {
			panic(fmt.Sprintf("Failed to ping Postgres: %v", err))
		}
	})

	// If pgOnce set up the pool successfully, poolInstance won't be nil (unless there was an error).
	return pgPoolInstance
}

// PingDB pings the database to check if the connection is alive.
func PingDB(ctx context.Context) error {
	if pgPoolInstance == nil {
		return fmt.Errorf("postgres pool is not initialized")
	}

	conn, err := pgPoolInstance.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %v", err)
	}
	defer conn.Release()

	if err := conn.Conn().Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Database connection is alive")
	return nil
}

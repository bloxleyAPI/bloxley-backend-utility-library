package utilslibs

type AppConfig struct {
	APP_PORT string
	ENV      string
}

// Config holds all required environment variables.
type PostgresConfig struct {
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_SSLMODE  string
}

// RedisConfig holds the connection details for Redis.
// The calling application is responsible for populating this struct.
type RedisConfig struct {
	REDIS_ADDRESS  string // "host:port" e.g. "localhost" or "10.0.0.5"
	REDIS_PORT     string // Redis port (e.g. 6379)
	REDIS_PASSWORD string // Redis password (can be empty if none)
	REDIS_DB       int    // Redis database number
}

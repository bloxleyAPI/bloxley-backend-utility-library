package utilslibs

type AppConfig struct {
	Port string
	Env  string
}

// Config holds all required environment variables.
type PostgresConfig struct {
	DBUser  string
	DBPass  string
	DBHost  string
	DBPort  string
	DBName  string
	SSLMode string
}

// RedisConfig holds the connection details for Redis.
// The calling application is responsible for populating this struct.
type RedisConfig struct {
	Address  string // "host:port" e.g. "localhost:6379" or "10.0.0.5:6379"
	User     string // Redis username (can be empty if none)
	Password string // Redis password (can be empty if none)
	DB       int    // Redis DB index (e.g. 0, 1, 2, ...)
}

package utilslibs

type AppConfig struct {
	Port string
	Env  string
}

// Config holds all required environment variables.
type PostgresConfig struct {
	DBUser     string
	DBPass     string
	DBHost     string
	DBPort     string
	DBName     string
	SSLMode    string
	DBTimezone string
}

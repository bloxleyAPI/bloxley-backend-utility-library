package utilslibs

type AppConfig struct {
	Port string
	Env  string
}

// Config holds all required environment variables.
type DBConfig struct {
	DBUser  string
	DBPass  string
	DBHost  string
	DBPort  string
	DBName  string
	SSLMode string
}

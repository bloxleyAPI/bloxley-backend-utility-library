package utilslibs

type AppConfig struct {
	Port string
	Env  string
}

// Config holds all required environment variables.
type Config struct {
	DBUser  string
	DBPass  string
	DBHost  string
	DBPort  string
	DBName  string
	SSLMode string
}

// LoadDBConfig loads the application configuration from environment variables.
func LoadDBConfig() *Config {
	return &Config{
		DBUser:  GetEnvOrFail(DBUserKey),
		DBPass:  GetEnvOrFail(DBPassKey),
		DBHost:  GetEnvOrFail(DBHostKey),
		DBPort:  GetEnvOrFail(DBPortKey),
		DBName:  GetEnvOrFail(DBNameKey),
		SSLMode: GetEnvOrFail(SSLModeKey),
	}
}

// LoadAppConfig loads the application configuration from environment variables.
func LoadAppConfig() *AppConfig {
	return &AppConfig{
		Port: GetEnvOrFail(AppPortKey),
		Env:  GetEnvOrFail(EnvKey),
	}
}

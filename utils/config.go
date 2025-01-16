package utilslib

import "github.com/bloxleyAPI/bloxley-backend-utility-library/constants"

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

// LoadAppConfig loads the application configuration from environment variables.
func LoadDBConfig() *Config {
	return &Config{
		DBUser:  GetEnvOrFail(constants.DBUserKey),
		DBPass:  GetEnvOrFail(constants.DBPassKey),
		DBHost:  GetEnvOrFail(constants.DBHostKey),
		DBPort:  GetEnvOrFail(constants.DBPortKey),
		DBName:  GetEnvOrFail(constants.DBNameKey),
		SSLMode: GetEnvOrFail(constants.SSLModeKey),
	}
}

// LoadAppConfig loads the application configuration from environment variables.
func LoadAppConfig() *AppConfig {
	return &AppConfig{
		Port: GetEnvOrFail(constants.AppPortKey),
		Env:  GetEnvOrFail(constants.EnvKey),
	}
}

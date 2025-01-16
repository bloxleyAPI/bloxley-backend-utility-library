package configs

import (
	"github.com/bloxleyAPI/bloxley-backend-utility-library/constants"
	"github.com/bloxleyAPI/bloxley-backend-utility-library/utils"
)

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
		DBUser:  utils.GetEnvOrFail(constants.DBUserKey),
		DBPass:  utils.GetEnvOrFail(constants.DBPassKey),
		DBHost:  utils.GetEnvOrFail(constants.DBHostKey),
		DBPort:  utils.GetEnvOrFail(constants.DBPortKey),
		DBName:  utils.GetEnvOrFail(constants.DBNameKey),
		SSLMode: utils.GetEnvOrFail(constants.SSLModeKey),
	}
}

// LoadAppConfig loads the application configuration from environment variables.
func LoadAppConfig() *AppConfig {
	return &AppConfig{
		Port: utils.GetEnvOrFail(constants.AppPortKey),
		Env:  utils.GetEnvOrFail(constants.EnvKey),
	}
}

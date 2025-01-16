package config

import (
	"github.com/bloxleyAPI/bloxley-backend-utility-library/constant"
	"github.com/bloxleyAPI/bloxley-backend-utility-library/util"
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
		DBUser:  util.GetEnvOrFail(constant.DBUserKey),
		DBPass:  util.GetEnvOrFail(constant.DBPassKey),
		DBHost:  util.GetEnvOrFail(constant.DBHostKey),
		DBPort:  util.GetEnvOrFail(constant.DBPortKey),
		DBName:  util.GetEnvOrFail(constant.DBNameKey),
		SSLMode: util.GetEnvOrFail(constant.SSLModeKey),
	}
}

// LoadAppConfig loads the application configuration from environment variables.
func LoadAppConfig() *AppConfig {
	return &AppConfig{
		Port: util.GetEnvOrFail(constant.AppPortKey),
		Env:  util.GetEnvOrFail(constant.EnvKey),
	}
}

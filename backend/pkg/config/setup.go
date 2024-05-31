package config

import (
	"os"
)

var defaultCfg = defaultConfig()
var Cfg *Config

type Config struct {
	Database *DBConfig
	Cors     *CorsConfig
}

func New() *Config {
	return &Config{
		Database: NewDBConfig(),
		Cors:     NewCorsConfig(),
	}
}

func EnvOrDefault(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return def
}

func defaultConfig() *Config {
	return &Config{
		Database: &DBConfig{
			Host:        "localhost",
			Port:        "3306",
			User:        "root",
			Password:    "",
			Database:    "stellar",
			LoggerLevel: "warn",
		},
		Cors: &CorsConfig{
			AllowedOrigin: "*",
		},
	}
}

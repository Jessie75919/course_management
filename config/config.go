package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strings"
)

type envConfig struct {
	DB_URL           string
	DEBUG bool
}

var EnvConfig envConfig

func Setup() {
	EnvConfig = envConfig{
		DB_URL:           os.Getenv("DB_URL"),
		DEBUG: parseDebug(),
	}
}

func parseDebug() bool {
	if strings.ToLower(os.Getenv("DEBUG")) == "true" {
		return true
	} else {
		return false
	}
}

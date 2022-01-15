package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strings"
)

type envConfig struct {
	DBUrl string
	DEBUG bool
}

var EnvConfig envConfig

func Setup() {
	EnvConfig = envConfig{
		DBUrl: os.Getenv("DB_URL"),
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

package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"strings"
)

type envConfig struct {
	DB_URL           string
	DEBUG            bool
	JWT_KEY          string
	JWT_LIFE_MINUTES int
}

var EnvConfig envConfig

func Setup() {
	EnvConfig = envConfig{
		DB_URL:           os.Getenv("DB_URL"),
		DEBUG:            parseDebug(),
		JWT_KEY:          os.Getenv("JWT_KEY"),
		JWT_LIFE_MINUTES: parseJwtLifeMinutes(),
	}
}

func parseJwtLifeMinutes() int {
	minutes, _ := strconv.Atoi(os.Getenv("JWT_LIFE_MINUTES"))
	return minutes
}

func parseDebug() bool {
	if strings.ToLower(os.Getenv("DEBUG")) == "true" {
		return true
	} else {
		return false
	}
}

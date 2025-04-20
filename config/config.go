package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURI        string
	LoggerURL    string
	LoggerAppKey string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println(".env not found, using environment variables as default")
		} else {
			log.Fatal("error loading .env", err)
		}
	}

	dbURI, ok := os.LookupEnv("DB_URI")
	if !ok {
		log.Fatal("DB_URI not configured")
	}

	loggerURL, ok := os.LookupEnv("LOGGER_URL")
	if !ok {
		log.Fatal("LOGGER_URL not configured")
	}

	loggerAppKey, ok := os.LookupEnv("LOGGER_APP_KEY")
	if !ok {
		log.Fatal("LOGGER_APP_KEY not configured")
	}

	return Config{
		DBURI:        dbURI,
		LoggerURL:    loggerURL,
		LoggerAppKey: loggerAppKey,
	}
}

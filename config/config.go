package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURI string
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

	return Config{
		DBURI: dbURI,
	}
}

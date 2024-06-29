package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DB  DBConfig
	JWT JWTConfig
	API APIConfig
}

type DBConfig struct {
	DSN string
}

type JWTConfig struct {
	SecretKey string
}

type APIConfig struct {
	ResumeParserAPIKey string
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	return Config{
		DB: DBConfig{
			DSN: os.Getenv("DSN"),
		},
		JWT: JWTConfig{
			SecretKey: os.Getenv("JWT_SECRET_KEY"),
		},
		API: APIConfig{
			ResumeParserAPIKey: os.Getenv("RESUME_PARSER_API_KEY"),
		},
	}, nil
}

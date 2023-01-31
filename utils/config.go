package utils

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)


type Config struct {
	DBSource			string
	DBDriver			string
	TokenSymetricKey	string
	TokenDuration		time.Duration
}

func LoadConfig()(*Config, error){
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	duration, err := time.ParseDuration(os.Getenv("TOKEN_DURATION"))
	config := Config{
		DBSource: os.Getenv("DB_SOURCE"),
		DBDriver: os.Getenv("DB_DRIVER"),
		TokenSymetricKey: os.Getenv("TOKEN_SYMETRIC_KEY"),
		TokenDuration: duration,
	}
	return &config, nil
}

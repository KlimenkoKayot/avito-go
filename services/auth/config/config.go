package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Logger                 string
	ServerPort             int
	DatabaseDSN            string
	TokenExpirationMinutes time.Duration
	ReadTimeoutSeconds     time.Duration
	WriteTimeoutSeconds    time.Duration
}

func Load(path string) (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	readTimeoutString := os.Getenv("READ_TIMEOUT")
	readTimeoutInt, err := strconv.Atoi(readTimeoutString)
	if err != nil {
		return nil, err
	}
	readTimeoutSeconds := time.Second * time.Duration(readTimeoutInt)

	writeTimeoutString := os.Getenv("WRITE_TIMEOUT")
	writeTimeoutInt, err := strconv.Atoi(writeTimeoutString)
	if err != nil {
		return nil, err
	}
	writeTimeoutSeconds := time.Second * time.Duration(writeTimeoutInt)

	tokenExpirationString := os.Getenv("TOKEN_EXPIRATION_TIMEOUT")
	tokenExpirationInt, err := strconv.Atoi(tokenExpirationString)
	if err != nil {
		return nil, err
	}
	tokenExpirationMinutes := time.Minute * time.Duration(tokenExpirationInt)

	databaseDSN := os.Getenv("DATABASE_DSN")

	logger := os.Getenv("LOGGER")

	serverPortString := os.Getenv("SERVER_PORT")
	serverPort, err := strconv.Atoi(serverPortString)
	if err != nil {
		return nil, err
	}

	return &Config{
		Logger:                 logger,
		ServerPort:             serverPort,
		DatabaseDSN:            databaseDSN,
		TokenExpirationMinutes: tokenExpirationMinutes,
		ReadTimeoutSeconds:     readTimeoutSeconds,
		WriteTimeoutSeconds:    writeTimeoutSeconds,
	}, nil
}

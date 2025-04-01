package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Logger              string
	Router              string
	ServerPort          int
	ReadTimeoutSeconds  time.Duration
	WriteTimeoutSeconds time.Duration
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

	logger := os.Getenv("LOGGER")
	router := os.Getenv("ROUTER")

	serverPortString := os.Getenv("SERVER_PORT")
	serverPort, err := strconv.Atoi(serverPortString)
	if err != nil {
		return nil, err
	}

	return &Config{
		Router:              router,
		Logger:              logger,
		ServerPort:          serverPort,
		ReadTimeoutSeconds:  readTimeoutSeconds,
		WriteTimeoutSeconds: writeTimeoutSeconds,
	}, nil
}

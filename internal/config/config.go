package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port          string
	MaxUploadSize int64
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	maxUploadSizeStr := os.Getenv("MAX_UPLOAD_SIZE")
	if maxUploadSizeStr == "" {
		maxUploadSizeStr = "10485760" // 10 mb
	}

	maxUploadSize, err := strconv.ParseInt(maxUploadSizeStr, 10, 64)
	if err != nil {
		log.Println("error parsing MAX_UPLOAD_SIZE")
		maxUploadSize = 10 << 20
	}

	return &Config{
		Port:          port,
		MaxUploadSize: maxUploadSize,
	}
}

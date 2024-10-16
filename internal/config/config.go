package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnectionString string
	ServerAddress      string
	AllowedOrigins     []string
	VideoPath          string
}

var AppConfig Config

func LoadConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	AppConfig = Config{
		DBConnectionString: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
		ServerAddress:  os.Getenv("SERVER_ADDRESS"),
		AllowedOrigins: strings.Split(os.Getenv("ALLOWED_ORIGIN"), ","),
		VideoPath:      fmt.Sprintf("%s/%s", homeDir, os.Getenv("VIDEO_PATH")),
	}
}

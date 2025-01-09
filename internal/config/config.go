package config

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type RateLimitValue struct {
	Duration int   `yaml:"duration"` // in seconds
	Limit    int64 `yaml:"limit"`
}

type RateLimitConfig struct {
	Search RateLimitValue `yaml:"search"`
	Frame  RateLimitValue `yaml:"frame"`
	GIF    RateLimitValue `yaml:"gif"`
}

type CacheConfig struct {
	CacheDuration   int `yaml:"duration"` // in seconds
	CleanupInterval int `yaml:"clean"`    // in seconds
}
type yamlConfig struct {
	Cache          CacheConfig     `yaml:"cache"`
	RateLimit      RateLimitConfig `yaml:"rate_limit"`
	ServerAddress  string          `yaml:"server"`
	AllowedOrigins []string        `yaml:"cors_origin"`
	VideoPath      string          `yaml:"video_path"`
}
type secretConfig struct {
	DBConnectionString string
}
type Config struct {
	secretConfig
	yamlConfig
}

var AppConfig Config

func LoadConfig() {
	file, err := os.Open("internal/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	yamlContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	yamlConfig := yamlConfig{}
	if err := yaml.Unmarshal(yamlContent, &yamlConfig); err != nil {
		log.Fatal(err)
	}

	AppConfig = Config{
		secretConfig: secretConfig{
			DBConnectionString: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"),
			),
		},
		yamlConfig: yamlConfig}
}

package config

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
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
	// Try multiple possible config file locations
	configPaths := []string{
		os.Getenv("CONFIG_PATH"),      // First try environment variable
		"internal/config/config.yaml", // Original path
		"/app/config/config.yaml",     // Common Docker path
		"config.yaml",                 // Root directory
	}

	var yamlContent []byte
	var err error
	var loadedPath string

	// Try each path until we successfully load the file
	for _, path := range configPaths {
		if path == "" {
			continue
		}
		if yamlContent, err = loadFile(path); err == nil {
			loadedPath = path
			break
		}
	}

	if yamlContent == nil {
		log.Fatalf("Could not find config file in any of the paths: %v", configPaths)
	}

	log.Printf("Loaded config from: %s", loadedPath)

	// Load .env file with multiple possible locations
	envFiles := []string{".env", "/app/.env"}
	envLoaded := false
	for _, envFile := range envFiles {
		if err := godotenv.Load(envFile); err == nil {
			envLoaded = true
			break
		}
	}
	if !envLoaded {
		log.Println("Warning: .env file not found")
	}

	yamlConfigObj := yamlConfig{}
	if err := yaml.Unmarshal(yamlContent, &yamlConfigObj); err != nil {
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
		yamlConfig: yamlConfigObj,
	}
}

func loadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

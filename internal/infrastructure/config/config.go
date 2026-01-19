package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	DBPort     string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: .env not found")
	}

	return &Config{
		Port:       getEnv("PORT", "3000"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBName:     getEnv("DB_NAME", "Shop"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBPort:     getEnv("DB_PORT", "5432"),
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

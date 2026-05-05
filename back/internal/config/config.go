package config

import (
	"log/slog"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTExpirationSeconds int
	JWTSecret            string
}

var (
	instance *Config
	once     sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load("../.env"); err != nil {
			slog.Warn("Env file not found, using environment variables")
		}

		instance = &Config{
			DBHost:     getEnv("DB_HOST", ""),
			DBPort:     getEnv("DB_PORT", ""),
			DBUser:     getEnv("DB_USER", ""),
			DBPassword: getEnv("DB_PASSWORD", ""),
			DBName:     getEnv("DB_NAME", ""),

			JWTExpirationSeconds: getEnvAsInt("JWT_EXPIRATION_SECONDS", 3600),
			JWTSecret:            getEnv("JWT_SECRET", ""),
		}
	})
	return instance
}

func GetConfig() *Config {
	if instance == nil {
		return LoadConfig()
	}
	return instance
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

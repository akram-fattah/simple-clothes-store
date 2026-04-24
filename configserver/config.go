package configserver

import (
  "fmt"
  "github.com/joho/godotenv"
  "os"
)

type Config struct {
  ServerPort string
  DatabaseUrl string
  Environment string
  LogLevel string
}

func LoadConfig() (*Config, error) {
  if err := godotenv.Load(); err != nil {
    return nil, fmt.Errorf("Faild to loading env file: %v", err)
  }
  return &Config{
    ServerPort: getEnv("SERVER_PORT", "8080"),
    DatabaseUrl: getEnv("DB_URL", "postgres"),
    Environment: getEnv("ENVIRONMENT", "development"),
    LogLevel: getEnv("LOG_LEVEL", "info"),
  }, nil
}

func getEnv(key, defaultValue string) string {
  if value, exists := os.LookupEnv(key); exists {
    return value
  }
  return defaultValue
}
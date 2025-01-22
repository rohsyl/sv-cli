package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ Warning: No .env file found, using default values.")
	}
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

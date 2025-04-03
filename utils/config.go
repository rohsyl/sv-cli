package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"runtime"
)

func LoadEnv() {

	var defaultEnv = "sv.config"
	if runtime.GOOS == "linux" {
		defaultEnv = "/data/sv/.config"
	}

	err := godotenv.Load(".env", defaultEnv)
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

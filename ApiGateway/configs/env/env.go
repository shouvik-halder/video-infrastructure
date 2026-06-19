package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading env variables")
	}
}

func GetString(key, fallback string) string {
	value, available := os.LookupEnv(key)
	if !available {
		return fallback
	}
	return value
}

func GetInt(key string, fallback int) int {
	value, available := os.LookupEnv(key)
	if !available {
		return fallback
	}

	if intValue, err := strconv.Atoi(value); err != nil {
		return intValue
	}

	return fallback

}

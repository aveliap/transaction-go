package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string  {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetEnvAsInt(key string, fallback int64) int64 {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file, using fallback values")
	}

	value:=os.Getenv(key)
	if value == "" {
		return fallback
	}

	i,err:=strconv.ParseInt(value, 10, 64)
	if err != nil{
		log.Printf("Error parsing %s: %v, using fallback value", key, err)
		return fallback
	}
	return i
}
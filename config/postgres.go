package config

import (
	"log"
	"strconv"
)


type Config struct{
	DBHost string
	DBPort int
	DBUser string
	DBPassword string
	DBName string
}

var Envs = initConfig()

func initConfig() Config  {
	dbPort, err := strconv.Atoi(GetEnv("POSTGRES_PORT"))
	if err != nil {
		log.Fatalf("Error converting PORT into int: %v",err)
	}
	return Config{
		DBHost: GetEnv("POSTGRES_HOST"),
		DBPort: dbPort,
		DBUser: GetEnv("POSTGRES_USER"),
		DBPassword: GetEnv("POSTGRES_PASSWORD"),
		DBName: GetEnv("POSTGRES_DB"),
	}
}
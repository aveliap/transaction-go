package config

import
	"github.com/aveliap/transaction-go/config/readerconf"

type Config struct{
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

func initConfig() Config  {
	return Config{
		DBHost: readerconf.GetEnv(),
	}
}
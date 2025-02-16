package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RoutConf ConfigRout
	DbCfg    DbCOnfig
}
type ConfigRout struct {
	Port string //формат 8080
}
type DbCOnfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewConfig() *Config {
	err := godotenv.Load("cfg.env")
	if err != nil {
		// потом хуйню
	}
	cfg := Config{}

	cfg.RoutConf.Port = os.Getenv("PORT")
	if cfg.RoutConf.Port == "" {
		cfg.RoutConf.Port = "8080"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "bookstore"
	}

	return &cfg

}

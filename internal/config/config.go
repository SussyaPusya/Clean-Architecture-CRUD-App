package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string //Формат типа ":8080"
}

func NewConfig() *Config {
	err := godotenv.Load("cfg.env")
	if err != nil {
		// потом хуйню
	}
	cfg := Config{}

	cfg.Port = os.Getenv("PORT")

	return &cfg

}

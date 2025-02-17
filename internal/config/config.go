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
	Dns string
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

	cfg.DbCfg.Dns = "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable"

	return &cfg

}

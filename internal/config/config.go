package config

import (
	"os"
)

type Config struct {
	RoutConf ConfigRout
}
type ConfigRout struct {
	Port string //формат 8080
}

func NewConfig() *Config {

	cfg := Config{}

	cfg.RoutConf.Port = os.Getenv("PORT")
	if cfg.RoutConf.Port == "" {
		cfg.RoutConf.Port = "8080"
	}

	return &cfg

}

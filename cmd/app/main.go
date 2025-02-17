package main

import (
	cfg "App/internal/config"
	"App/internal/repository"
	"App/internal/service"
	router "App/internal/transport/http"
	"App/pkg/postgres"
)

func main() {

	//тут сверху логика всей хуйни

	config := cfg.NewConfig()

	db, err := postgres.NewPosgrs(config.DbCfg)
	if err != nil {

	}

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	handler := router.NewHandler(service) //там внутри должен быть сервис

	r := router.NewRouter(handler, config)

	r.Run()

}

package main

import (
	"App/internal/repository"
	"App/internal/service"
	router "App/internal/transport/http"
	"App/pkg/postgres"
)

func main() {

	//тут сверху логика всей хуйни

	db, err := postgres.NewPosgrs()
	if err != nil {

	}

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	handler := router.NewHandler(service) //там внутри должен быть сервис

	r := router.NewRouter(handler)

	r.Run()

}

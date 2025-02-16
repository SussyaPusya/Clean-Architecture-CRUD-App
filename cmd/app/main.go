package main

import (
	router "App/internal/transport/http"
)

func main() {

	//тут сверху логика всей хуйни

	handler := router.NewHandler() //там внутри должен быть сервис

	r := router.NewRouter(handler)

	r.Run()

}

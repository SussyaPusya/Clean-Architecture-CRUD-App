package http

import (
	"App/internal/entity"
	"App/pkg"
	"io"
	"log"
	"net/http"
)

type Service interface {
	CreateData(entity.User) (*entity.User, error)
	GetData(entity.User) (*entity.User, error)
	UpdateData(entity.User) (*entity.User, error)
	DeleteData(entity.User) (*entity.User, error)
}

type Handler struct {
	service Service
}

func (h *Handler) GetData(w http.ResponseWriter, r *http.Request) {
	bytew, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("invalid json")
	}
	defer r.Body.Close()
	var user entity.User
	person := pkg.ParseJson(bytew, user)
	h.service.GetData(person)
	// тут дальше надо так чтобы челу отправился ответ с гео данными там дата рождения и тп
}

func (h *Handler) UpdateData(w http.ResponseWriter, r *http.Request) {
	bytew, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("invalid json")
	}
	defer r.Body.Close()
	var user entity.User
	person := pkg.ParseJson(bytew, user)
	h.service.UpdateData(person)
	// вывод обновлённых данных

}

func (h *Handler) DeleteData(w http.ResponseWriter, r *http.Request) {
	bytew, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("invalid json")
	}
	defer r.Body.Close()
	var user entity.User
	person := pkg.ParseJson(bytew, user)
	h.service.DeleteData(person)
	//вывод что всё адлилос
}

package http

import (
	"App/internal/entity"
	"App/pkg"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Service interface {
	CreateData(entity.User) error
	GetData(entity.User) (*entity.User, error)
	UpdateData(entity.User) error
	DeleteData(entity.User) error
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// Чтобы получить дату надо отправить json с полем name
func (h *Handler) GetData(w http.ResponseWriter, r *http.Request) {
	bytew, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("invalid json")
	}
	defer r.Body.Close()
	var user entity.User
	person := pkg.ParseJson(bytew, user)

	jsnResp, _ := h.service.GetData(person)

	jsonBytes, err := pkg.ToJson(*jsnResp)
	if err != nil {
		fmt.Fprintln(w, "Filed to get data, err: ", err)

	}

	fmt.Fprintln(w, string(jsonBytes))

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
	err = h.service.UpdateData(person)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintln(w, "Updated done")

	// вывод обновлённых данных

}

func (h *Handler) DeleteData(w http.ResponseWriter, r *http.Request) {
	bytew, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("invalid json", err)
	}
	defer r.Body.Close()
	var user entity.User
	person := pkg.ParseJson(bytew, user)
	h.service.DeleteData(person)

	fmt.Fprintln(w, "Deleted Done!")
	//вывод что всё адлилос
}

func (h *Handler) CreateData(w http.ResponseWriter, r *http.Request) {
	bytew, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("invalid json")
	}
	defer r.Body.Close()
	var user entity.User
	person := pkg.ParseJson(bytew, user)
	err = h.service.CreateData(person)
	if err != nil {
		fmt.Fprintln(w, "server error")
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Fprintln(w, "Data created!)")

	//вывод что всё адлилос
}

package http

import (
	"App/internal/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Router  *mux.Router
	Handler Handler
	Config  *config.Config
}

func NewRouter(h *Handler) *Router {
	rout := mux.NewRouter()

	rout.HandleFunc("/createdata", h.CreateData)

	rout.HandleFunc("/getdata", h.GetData)
	rout.HandleFunc("/updatedata", h.UpdateData)

	rout.HandleFunc("/deletedata", h.DeleteData)

	return &Router{Router: rout, Config: config.NewConfig()}
}

func (r *Router) Run() {
	log.Fatal(http.ListenAndServe(r.Config.Port, r.Router))
}

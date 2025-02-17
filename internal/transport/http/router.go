package http

import (
	cfg "App/internal/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Router  *mux.Router
	Handler Handler
	Config  *cfg.Config
}

func NewRouter(h *Handler, cfg *cfg.Config) *Router {
	rout := mux.NewRouter()

	rout.HandleFunc("/createdata", h.CreateData)

	rout.HandleFunc("/getdata", h.GetData)
	rout.HandleFunc("/updatedata", h.UpdateData)

	rout.HandleFunc("/deletedata", h.DeleteData)

	return &Router{Router: rout, Config: cfg}
}

func (r *Router) Run() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", r.Config.RoutConf.Port), r.Router))
}

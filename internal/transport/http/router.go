package http

import "github.com/gorilla/mux"

type Router struct {
	Router  *mux.Router
	Handler Handler
}

func NewRouter(h *Handler) *Router {
	rout := mux.NewRouter()

	// rout.HandleFunc("/", ).Methods("GET")

	return &Router{Router: rout}
}

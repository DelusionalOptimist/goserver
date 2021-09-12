package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DelusionalOptimist/goserver/handlers"
)

type Router struct {
	mux  *http.ServeMux
	port int
}

func NewRouter(port int, l *log.Logger) *Router {
	mux := http.NewServeMux()

	handler := handlers.NewHandlerInstance(l)
	mux.HandleFunc("/", handler.SimpleReply)

	return &Router{
		mux:  mux,
		port: port,
	}
}

func (r *Router) Run() error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", r.port), r.mux)
	if err != nil {
		return err
	}
	return nil
}

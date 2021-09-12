package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DelusionalOptimist/goserver/handlers"
)

// type Router implements a router
// it consists of a ServeMux for listening to and multiplexing requests
// it also has a port which is used for specifying the port to listen on
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

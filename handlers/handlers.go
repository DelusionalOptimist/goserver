package handlers

import (
	"log"
)

// type Handler is a common type for all Handlers
// It has a logger and a request number
type Handler struct {
	log          *log.Logger
	requestIndex int64
}

// func NewHandlerInstance() initializes and returns a new Handler
func NewHandlerInstance(l *log.Logger) *Handler {
	return &Handler{l, 0}
}

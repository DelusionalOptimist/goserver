package handlers

import (
	"log"
)

type Handler struct {
	log          *log.Logger
	requestIndex int64
}

func NewHandlerInstance(l *log.Logger) *Handler {
	return &Handler{l, 0}
}

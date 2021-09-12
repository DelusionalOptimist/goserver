package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// func SimpleReply() is a dead simple http handler.
// For GET: It returns the request number
// For POST: It acts as an echo server
func (h *Handler) SimpleReply(rw http.ResponseWriter, r *http.Request) {
	h.requestIndex = h.requestIndex + 1

	switch r.Method {
	case http.MethodGet:
		h.log.Println("Request ", h.requestIndex)
		fmt.Fprintf(rw, "Request %d!\n", h.requestIndex)

	case http.MethodPost:
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			h.log.Println("Error while reading request body")
			http.Error(rw, "Bad request body", http.StatusBadRequest)
			return
		}
		h.log.Println(string(body))
		fmt.Fprintln(rw, string(body))

	default:
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"app/model"
)

// A HelloHandler implements health check endpoint.
type HelloHandler struct {
}

// NewHelloHandler returns HelloHandler based http.Handler.
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World\n")
	hr := &model.HelloResponse{Message: "Hello"}
	enc := json.NewEncoder(w)
	if err := enc.Encode(hr); err != nil {
		log.Println(err)
	}
}

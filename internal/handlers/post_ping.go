package handlers

import (
	"gochat/internal/templates/components"
	"net/http"

	"github.com/a-h/templ"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	templ.Handler(components.PingResponse()).ServeHTTP(w, r)
}

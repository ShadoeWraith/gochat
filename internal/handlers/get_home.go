package handlers

import (
	"gochat/internal/templates"
	"net/http"

	"github.com/a-h/templ"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Push-Url", "/")
		templ.Handler(templates.Content()).ServeHTTP(w, r)
		return
	} else {
		pageComponent := templates.Layout("My Awesome Page Title", templates.Content())
		templ.Handler(pageComponent).ServeHTTP(w, r)
	}
}

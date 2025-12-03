package main

import (
	"fmt"
	"net/http"

	"gochat/components"

	"github.com/a-h/templ"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/test", testHandler)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pageComponent := layout("My Awesome Page Title", components.Content())
	templ.Handler(pageComponent).ServeHTTP(w, r)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	templ.Handler(components.PingResponse()).ServeHTTP(w, r)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Push-Url", "/test")
		templ.Handler(components.Test()).ServeHTTP(w, r)
	} else {
		pageComponent := layout("Test Page", components.Test())
		templ.Handler(pageComponent).ServeHTTP(w, r)
	}
}

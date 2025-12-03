package main

import (
	"fmt"
	"gochat/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.NewHomeHandler().ServeHTTP)
	http.HandleFunc("/ping", handlers.NewPingHandler().ServeHTTP)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

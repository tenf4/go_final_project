package server

import (
	"net/http"
	"os"
)

func LaunchServer() error {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}
	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	return http.ListenAndServe(":"+port, nil)
}

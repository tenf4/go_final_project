package api

import (
	"encoding/json"
	"net/http"
)

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
}
func writeJson(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return json.NewEncoder(w).Encode(data)
}

func writeJsonError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

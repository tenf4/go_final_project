package api

import (
	"fmt"
	"go_final_project/pkg/db"
	"net/http"
)

func taskHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	// обработка других методов будет добавлена на следующих шагах
	case http.MethodPost:
		addTaskHandler(w, req)
	case http.MethodPut:
		//updateTaskHandler(w, req)
		addTaskHandler(w, req)
	case http.MethodGet:
		getTaskHandler(w, req)
	}

}

func getTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		writeJsonError(w, http.StatusBadRequest, fmt.Errorf("incorrect id"))
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJsonError(w, http.StatusBadRequest, fmt.Errorf("can't find task"))
		return
	}
}

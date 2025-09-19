package api

import (
	"encoding/json"
	"fmt"
	"go_final_project/pkg/db"
	"io"
	"net/http"
	"time"
)

func taskHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	// обработка других методов будет добавлена на следующих шагах
	case http.MethodPost:
		addTaskHandler(w, req)
	case http.MethodPut:
		updateTaskHandler(w, req)

	case http.MethodGet:
		getTaskHandler(w, req)

	case http.MethodDelete:
		deleteTaskHandler(w, req)

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

	writeJson(w, task)
}

func updateTaskHandler(w http.ResponseWriter, req *http.Request) {
	var task db.Task
	now := time.Now()
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writeJsonError(w, http.StatusBadRequest, fmt.Errorf("invalid body request"))
		return
	}

	if err := json.Unmarshal(data, &task); err != nil {
		writeJsonError(w, http.StatusBadRequest, fmt.Errorf("json format invalid"))
		return
	}

	if task.ID == "" {
		writeJsonError(w, http.StatusBadRequest, fmt.Errorf("id is empty"))
		return
	}

	if task.Title == "" {
		writeJsonError(w, http.StatusBadRequest, fmt.Errorf("title is empty"))
		return
	}

	if task.Date == "" {
		task.Date = now.Format("20060102")
	}

	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		writeJsonError(w, http.StatusBadRequest, err)
		return
	}

	if now.After(t) {
		if task.Repeat == "" {
			task.Date = now.Format("20060102")
		} else {
			next_date, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				http.Error(w, `{"error" : "incorrect repeat format"}`, http.StatusBadRequest)
				return
			}
			task.Date = next_date
		}

	}
	err = db.UpdateTask(&task)
	if err != nil {
		writeJsonError(w, http.StatusBadRequest, err)
		return
	}
	writeJson(w, map[string]interface{}{})
}

func deleteTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		writeJsonError(w, http.StatusBadRequest, fmt.Errorf("can't find id parameter"))
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		writeJsonError(w, http.StatusBadRequest, fmt.Errorf("error while deleting task: %v", err))
		return
	}

	writeJson(w, map[string]interface{}{})
}

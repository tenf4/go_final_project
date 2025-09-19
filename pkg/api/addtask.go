package api

import (
	"encoding/json"
	"fmt"
	"go_final_project/pkg/db"
	"io"
	"net/http"
	"time"
)

func addTaskHandler(w http.ResponseWriter, req *http.Request) {
	var task db.Task

	data, _ := io.ReadAll(req.Body)
	json.Unmarshal(data, &task)
	now := time.Now()

	if task.Title == "" {
		http.Error(w, `{"error" : "task title is empty"}`, http.StatusBadRequest)
		return
	}

	if task.Date == "" {
		task.Date = now.Format("20060102")
	} else {
		t, err := time.Parse("20060102", task.Date)
		if err != nil {
			http.Error(w, `{"error" : "incorrect time format"}`, http.StatusBadRequest)
			return
		}

		var next_date string
		if task.Repeat != "" {
			next_date, err = NextDate(now, task.Date, task.Repeat)
			if err != nil {
				http.Error(w, `{"error" : "incorrect repeat format"}`, http.StatusBadRequest)
				return
			}
		}

		nowDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		taskDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)

		if nowDate.After(taskDate) {
			if task.Repeat == "" {
				task.Date = now.Format("20060102")
			} else {
				task.Date = next_date
			}
		}
	}

	taskId, err := db.AddTask(&task)
	if err != nil {
		http.Error(w, `{"error" : "error while adding the task to database"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	_, err = w.Write([]byte(fmt.Sprintf(`{"id" : "%d"}`, taskId)))
	if err != nil {
		http.Error(w, `{"error" : "server error"}`, http.StatusInternalServerError)
		return
	}
}

package api

import (
	"fmt"
	"go_final_project/pkg/db"
	"net/http"
	"time"
)

func taskDoneHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	now := time.Now()

	if id == "" {
		writeJsonError(w, http.StatusBadRequest, fmt.Errorf("id not found"))
		return
	}
	task, err := db.GetTask(id)
	if err != nil {
		writeJsonError(w, http.StatusBadRequest, fmt.Errorf("can't find task"))
		return
	}

	if task.Repeat == "" {
		err := db.DeleteTask(id)
		if err != nil {
			writeJsonError(w, http.StatusBadRequest, fmt.Errorf("error while deleting task: %v", err))
			return
		}
	} else {
		next_date, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJsonError(w, http.StatusBadRequest, fmt.Errorf("error while calculating date: %v", err))
			return
		}
		err = db.UpdateDate(id, next_date)
		if err != nil {
			writeJsonError(w, http.StatusBadRequest, fmt.Errorf("error while updating date: %v", err))
			return
		}

	}
	writeJson(w, map[string]interface{}{})

}

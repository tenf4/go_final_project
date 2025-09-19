package api

import (
	"fmt"
	"go_final_project/pkg/db"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, fmt.Errorf("maximum value of entries reached"))
		return
	}

	if r.Method != http.MethodGet {
		writeJsonError(w, http.StatusBadRequest, fmt.Errorf("this method is not allowed"))
		return
	}
	if tasks == nil {
		tasks = []*db.Task{}
	}

	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}

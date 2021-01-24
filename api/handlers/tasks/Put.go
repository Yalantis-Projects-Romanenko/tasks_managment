package tasks

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/fdistorted/task_managment/db/projects"
	dbTasks "github.com/fdistorted/task_managment/db/tasks"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/logger"
	"github.com/fdistorted/task_managment/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Put(w http.ResponseWriter, r *http.Request) {
	userId, ok := middlewares.GetUserID(r.Context())
	if !ok {
		common.SendResponse(w, http.StatusBadRequest, common.FailedToGetUserId)
		return
	}

	vars := mux.Vars(r)
	projectId := vars["projectId"]
	columnId := vars["columnId"]
	taskId := vars["taskId"]

	var task models.Task

	// decode the json request to task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, common.FailedToParseJson)
		return
	}

	// check if task exist and owned by a user
	_, err = projects.GetById(r.Context(), userId, projectId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			common.SendResponse(w, http.StatusNotFound, common.ResourceIsNotOwned)
			return
		}
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	task.Id = taskId
	task.ProjectId = projectId
	task.ColumnId = columnId

	err = dbTasks.Update(r.Context(), task)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	common.SendResponse(w, http.StatusOK, "task updated")
}

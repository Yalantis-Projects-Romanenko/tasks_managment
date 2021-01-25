package tasks

import (
	"encoding/json"
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

	// check if project exist and owned by a user
	if !common.CheckUsersProperty(w, r, userId, projectId) {
		return
	}

	task.Id = taskId
	task.ProjectId = projectId

	err = dbTasks.Update(r.Context(), columnId, task)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	common.SendResponse(w, http.StatusOK, "task updated")
}

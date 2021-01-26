package tasks

import (
	"database/sql"
	"encoding/json"
	"errors"
	dbTasks "github.com/fdistorted/task_managment/db/tasks"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/logger"
	"github.com/fdistorted/task_managment/models"
	vld "github.com/fdistorted/task_managment/validator"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	userId, ok := middlewares.GetUserID(r.Context())
	if !ok {
		common.SendResponse(w, http.StatusInternalServerError, common.ErrFailedToGetUserId)
		return
	}

	var task models.Task

	// decode the json request to task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, common.ErrFailedToParseJson)
		return
	}

	validate := vld.Get()
	err = validate.Struct(task)
	if err != nil {
		errors := vld.ParseValidationErrors(err)
		common.SendResponse(w, http.StatusUnprocessableEntity, errors)
		return
	}

	vars := mux.Vars(r)
	projectId := vars["projectId"]
	columnId := vars["columnId"]

	// check if project exist and owned by a user
	if !common.CheckUsersProperty(w, r, userId, projectId) {
		return
	}

	taskId, err := dbTasks.CreateTask(r.Context(), columnId, projectId, task)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			common.SendResponse(w, http.StatusNotFound, common.ErrNotFound)
		}
		common.SendResponse(w, http.StatusInternalServerError, common.ErrDatabaseError)
		return
	}

	task.Id = taskId

	logger.WithCtxValue(r.Context()).Info("created task", zap.Any("task", task))
	common.SendResponse(w, http.StatusOK, task)
}

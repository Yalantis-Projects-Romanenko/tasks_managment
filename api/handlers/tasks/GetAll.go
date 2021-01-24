package tasks

import (
	"database/sql"
	"errors"
	dbProjects "github.com/fdistorted/task_managment/db/projects"
	dbTasks "github.com/fdistorted/task_managment/db/tasks"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/logger"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	userId, ok := middlewares.GetUserID(r.Context())
	if !ok {
		common.SendResponse(w, http.StatusBadRequest, common.FailedToGetUserId)
		return
	}

	vars := mux.Vars(r)
	projectId := vars["projectId"]
	columnId := vars["columnId"]

	// check if project exist and owned by a user
	_, err := dbProjects.GetById(r.Context(), userId, projectId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			common.SendResponse(w, http.StatusNotFound, common.ResourceIsNotOwned)
			return
		}
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	gotTasks, err := dbTasks.GetAll(r.Context(), columnId, projectId)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	logger.WithCtxValue(r.Context()).Info("got gotTasks from the database ", zap.Int("columns_len", len(gotTasks)))
	common.SendResponse(w, http.StatusOK, gotTasks)
}

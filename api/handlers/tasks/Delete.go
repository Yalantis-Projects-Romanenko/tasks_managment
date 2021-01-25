package tasks

import (
	dbTasks "github.com/fdistorted/task_managment/db/tasks"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/logger"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	columnId := vars["columnId"]
	taskId := vars["taskId"]
	userId, ok := middlewares.GetUserID(r.Context())
	if !ok {
		common.SendResponse(w, http.StatusInternalServerError, common.FailedToGetUserId)
		return
	}

	// check if project exist and owned by a user
	if !common.CheckIfUserExists(w, r, userId, projectId) {
		return
	}

	affected, err := dbTasks.DeleteById(r.Context(), projectId, columnId, taskId)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	logger.WithCtxValue(r.Context()).Info("Total rows/record affected %v", zap.Int64("rowsAffected", affected))

	common.SendResponse(w, http.StatusOK, "deleted")
}

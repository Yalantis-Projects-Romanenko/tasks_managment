package columns

import (
	"github.com/fdistorted/task_managment/db/columns"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/logger"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

// TODO delete everything under column: tasks, comments
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	columnId := vars["columnId"]
	userId, ok := middlewares.GetUserID(r.Context())
	if !ok {
		common.SendResponse(w, http.StatusInternalServerError, common.FailedToGetUserId)
		return
	}

	// check if project exist and owned by a user
	if !common.CheckUsersProperty(w, r, userId, projectId) {
		return
	}

	affected, err := columns.DeleteById(r.Context(), projectId, columnId)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	logger.WithCtxValue(r.Context()).Info("Total rows/record affected %v", zap.Int64("rowsAffected", affected))

	common.SendResponse(w, http.StatusOK, "deleted")
}

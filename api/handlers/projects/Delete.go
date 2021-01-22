package projects

import (
	projects2 "github.com/fdistorted/task_managment/db/projects"
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
	userId, ok := middlewares.GetUserID(r.Context())
	if !ok {
		common.SendResponse(w, http.StatusInternalServerError, common.FailedToGetUserId)
		return
	}

	affected, err := projects2.DeleteById(r.Context(), userId, projectId)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	logger.WithCtxValue(r.Context()).Info("Total rows/record affected %v", zap.Int64("rowsAffected", affected))

	common.SendResponse(w, http.StatusOK, "deleted")
}

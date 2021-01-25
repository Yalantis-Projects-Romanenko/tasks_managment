package comments

import (
	"github.com/fdistorted/task_managment/db/comments"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/logger"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	userId, ok := middlewares.GetUserID(r.Context())
	if !ok {
		common.SendResponse(w, http.StatusBadRequest, common.FailedToGetUserId)
		return
	}

	vars := mux.Vars(r)
	projectId := vars["projectId"]
	taskId := vars["taskId"]
	commentId := vars["commentId"]

	// check if project exist and owned by a user
	if !common.CheckUsersProperty(w, r, userId, projectId) {
		return
	}

	gotTask, err := comments.GetById(r.Context(), projectId, taskId, commentId)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	common.SendResponse(w, http.StatusOK, gotTask)
}

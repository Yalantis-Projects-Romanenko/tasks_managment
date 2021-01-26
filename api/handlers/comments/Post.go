package comments

import (
	"database/sql"
	"encoding/json"
	"errors"
	dbComments "github.com/fdistorted/task_managment/db/comments"
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

	var comment models.Comment

	// decode the json request to comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, common.ErrFailedToParseJson)
		return
	}

	validate := vld.Get()
	err = validate.Struct(comment)
	if err != nil {
		errors := vld.ParseValidationErrors(err)
		common.SendResponse(w, http.StatusUnprocessableEntity, errors)
		return
	}

	vars := mux.Vars(r)
	projectId := vars["projectId"]
	//commentId := vars["commentId"]
	taskId := vars["taskId"]

	// check if project exist and owned by a user
	if !common.CheckUsersProperty(w, r, userId, projectId) {
		return
	}

	comment.ProjectId = projectId
	comment.TaskId = taskId
	comment.UserId = userId
	commentId, err := dbComments.CreateComment(r.Context(), comment)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			common.SendResponse(w, http.StatusNotFound, common.ErrNotFound)
		}
		common.SendResponse(w, http.StatusInternalServerError, common.ErrDatabaseError)
		return
	}

	comment.Id = commentId
	logger.WithCtxValue(r.Context()).Info("created comment", zap.Any("comment", comment))
	common.SendResponse(w, http.StatusOK, comment)
}

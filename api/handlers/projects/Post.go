package projects

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/fdistorted/task_managment/db/projects"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/logger"
	"github.com/fdistorted/task_managment/models"
	vld "github.com/fdistorted/task_managment/validator"
	"go.uber.org/zap"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	userId, ok := middlewares.GetUserID(r.Context())
	if !ok {
		common.SendResponse(w, http.StatusInternalServerError, common.ErrFailedToGetUserId)
		return
	}

	var project models.Project

	// decode the json request to project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, common.ErrFailedToParseJson)
		return
	}

	validate := vld.Get()
	err = validate.Struct(project)
	if err != nil {
		errors := vld.ParseValidationErrors(err)
		common.SendResponse(w, http.StatusUnprocessableEntity, errors)
		return
	}

	project.UserId = userId

	id, err := projects.CreateProject(r.Context(), project)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			common.SendResponse(w, http.StatusNotFound, common.ErrNotFound)
		}
		common.SendResponse(w, http.StatusInternalServerError, common.ErrDatabaseError)
		return
	}
	project.Id = id
	logger.WithCtxValue(r.Context()).Info("created project", zap.Any("project", project))
	common.SendResponse(w, http.StatusOK, project)
}

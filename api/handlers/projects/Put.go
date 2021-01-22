package projects

import (
	"encoding/json"
	"github.com/fdistorted/task_managment/db/projects"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/logger"
	"github.com/fdistorted/task_managment/models"
	vld "github.com/fdistorted/task_managment/validator"
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

	var project models.Project

	// decode the json request to project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, common.FailedToParseJson)
		return
	}

	validate := vld.Get()
	err = validate.Struct(project)
	if err != nil {
		errors := vld.ParseValidationErrors(err)
		common.SendResponse(w, http.StatusUnprocessableEntity, errors)
		return
	}

	project.Id = projectId
	project.UserId = userId
	affected, err := projects.Update(r.Context(), project)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	logger.WithCtxValue(r.Context()).Info("Total rows/record affected %v", zap.Int64("rowsAffected", affected))
	common.SendResponse(w, http.StatusOK, "project updated")
}

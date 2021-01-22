package columns

import (
	"database/sql"
	"encoding/json"
	"errors"
	dbColumns "github.com/fdistorted/task_managment/db/columns"
	"github.com/fdistorted/task_managment/db/projects"
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

	var column models.Column

	// decode the json request to column
	err := json.NewDecoder(r.Body).Decode(&column)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, common.FailedToParseJson)
		return
	}

	// check if column exist and owned by a user
	_, err = projects.GetById(r.Context(), userId, projectId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			common.SendResponse(w, http.StatusNotFound, common.ResourceIsNotOwned)
			return
		}
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	column.Id = columnId
	err = dbColumns.Update(r.Context(), userId, projectId, column)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	common.SendResponse(w, http.StatusOK, "column updated")
}

package columns

import (
	"encoding/json"
	"github.com/fdistorted/task_managment/db/columns"
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
		common.SendResponse(w, http.StatusInternalServerError, common.FailedToGetUserId)
		return
	}

	var column models.Column

	// decode the json request to column
	err := json.NewDecoder(r.Body).Decode(&column)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, common.FailedToParseJson)
		return
	}

	validate := vld.Get()
	err = validate.Struct(column)
	if err != nil {
		errors := vld.ParseValidationErrors(err)
		common.SendResponse(w, http.StatusUnprocessableEntity, errors)
		return
	}

	vars := mux.Vars(r)
	projectId := vars["projectId"]

	columnId, err := columns.CreateColumn(userId, projectId, column, r.Context())
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	logger.WithCtxValue(r.Context()).Info("New record ID is:", zap.String("projectId", columnId))

	common.SendResponse(w, http.StatusOK, "column created")
	return
}

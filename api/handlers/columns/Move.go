package columns

import (
	"encoding/json"
	database "github.com/fdistorted/task_managment/db"
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

func Move(w http.ResponseWriter, r *http.Request) {
	userId, ok := middlewares.GetUserID(r.Context())
	if !ok {
		common.SendResponse(w, http.StatusInternalServerError, common.FailedToGetUserId)
		return
	}

	var column models.ColumnMove

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

	err = columns.ChangeIndex(userId, projectId, column.Id, column.Index, r.Context())
	if err != nil {
		if database.ErrInvalidParameters.Is(err) {
			common.SendResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.SendResponse(w, http.StatusInternalServerError, common.DatabaseError)
		return
	}

	common.SendResponse(w, http.StatusOK, "column created") // todo change response accrding to rest
	return
}

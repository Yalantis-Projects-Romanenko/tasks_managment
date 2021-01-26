package projects

import (
	"database/sql"
	"errors"
	projects2 "github.com/fdistorted/task_managment/db/projects"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	userId, ok := middlewares.GetUserID(r.Context())
	if !ok {
		common.SendResponse(w, http.StatusInternalServerError, common.ErrFailedToGetUserId)
		return
	}
	project, err := projects2.GetById(r.Context(), userId, projectId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			common.SendResponse(w, http.StatusNotFound, common.ErrNotFound)
		}
		common.SendResponse(w, http.StatusInternalServerError, common.ErrDatabaseError)
		return
	}
	common.SendResponse(w, http.StatusOK, project)
}

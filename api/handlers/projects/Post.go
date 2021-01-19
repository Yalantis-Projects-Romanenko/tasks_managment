package projects

import (
	"encoding/json"
	"github.com/fdistorted/task_managment/db/projects"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/logger"
	"github.com/fdistorted/task_managment/models"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {

	// create an empty user of type models.User
	var project models.Project

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
		return
	}

	userId, ok := middlewares.GetUserID(r.Context())
	if !ok {
		common.SendResponse(w, http.StatusInternalServerError, "failed to get userId")
		return
	}

	id, err := projects.Insert(userId, project)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "failed to create project")
		return
	}

	logger.Get().Info("New record ID is:", zap.String("projectId", id))

	common.SendResponse(w, http.StatusOK, "project created")
	return
}

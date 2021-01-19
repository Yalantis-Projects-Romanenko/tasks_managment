package projects

import (
	"encoding/json"
	"github.com/fdistorted/task_managment/db/projects"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
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

	project.Id = id
	projects.Update(userId, project)
	common.SendResponse(w, http.StatusOK, "project updated")
}

package projects

import (
	"github.com/fdistorted/task_managment/handlers/common"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//vars["project_id"]

	common.SendResponse(w, http.StatusInternalServerError, "not implemented yet")
}

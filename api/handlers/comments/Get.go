package comments

import (
	"github.com/fdistorted/task_managment/handlers/common"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]

	common.SendResponse(w, http.StatusInternalServerError, "not implemented yet")
}

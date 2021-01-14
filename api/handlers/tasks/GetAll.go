package tasks

import (
	"github.com/fdistorted/task_managment/handlers/common"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	common.SendResponse(w, http.StatusInternalServerError, "not implemented yet")
}

package projects

import (
	projects2 "github.com/fdistorted/task_managment/db/projects"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/logger"
	"go.uber.org/zap"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	userId, _ := middlewares.GetUserID(r.Context())
	projects := projects2.GetAllByUserId(userId) // todo use user id from context
	logger.Get().Info("got projects from the database ", zap.Int("projects_len", len(projects)))
	common.SendResponse(w, http.StatusOK, projects)
}

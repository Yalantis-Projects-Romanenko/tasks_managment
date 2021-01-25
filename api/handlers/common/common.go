package common

import (
	"encoding/json"
	dbProjects "github.com/fdistorted/task_managment/db/projects"
	"github.com/fdistorted/task_managment/logger"
	"go.uber.org/zap"
	"net/http"
)

// SendResponse - encode response to json and send it.
func SendResponse(w http.ResponseWriter, statusCode int, respBody interface{}) {
	binRespBody, err := json.Marshal(respBody)
	if err != nil {
		logger.Get().Error("failed to marshal response body to json", zap.Error(err))
		statusCode = http.StatusInternalServerError
	}

	SendRawResponse(w, statusCode, binRespBody)
}

// SendRawResponse sends any raw ([]byte) response.
func SendRawResponse(w http.ResponseWriter, statusCode int, binBody []byte) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	w.WriteHeader(statusCode)
	_, err := w.Write(binBody)
	if err != nil {
		logger.Get().Error("failed to write response body", zap.Error(err))
	}
}

func CheckUsersProperty(w http.ResponseWriter, r *http.Request, userId, projectId string) bool {
	// check if project exist and owned by a user
	exists, err := dbProjects.Exists(r.Context(), userId, projectId)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		SendResponse(w, http.StatusInternalServerError, DatabaseError)
		return false
	}
	if !exists {
		SendResponse(w, http.StatusNotFound, ResourceIsNotOwned)
		return false
	}

	return true
}

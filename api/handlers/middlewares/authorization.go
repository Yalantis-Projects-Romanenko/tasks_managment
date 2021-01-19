package middlewares

import (
	"context"
	"encoding/base64"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/logger"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"unicode"
)

const authorizationHeader = "Authorization"
const UserIdFieldName = "user_id"

// AuthToken gets the auth token from the context.
func GetUserID(ctx context.Context) (string, bool) {
	tokenStr, ok := ctx.Value(UserIdFieldName).(string)
	return tokenStr, ok
}

func CheckUsername(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			header := r.Header.Get(authorizationHeader)
			if header == "" {
				common.SendResponse(w, http.StatusUnauthorized, "no auth token")
				return
			}

			tokenParts := strings.Split(header, " ")

			if len(tokenParts) != 2 {
				common.SendResponse(w, http.StatusUnauthorized, "wrong token format")
				return
			}

			data, err := base64.StdEncoding.DecodeString(tokenParts[1])
			if err != nil {
				logger.Get().Error("failed to decode token", zap.Error(err))
				common.SendResponse(w, http.StatusUnauthorized, "failed to decode token")
				return
			}

			username := string(data)
			if !CheckUsername(username) {
				common.SendResponse(w, http.StatusUnauthorized, "wrong username")
				return
			}

			logger.Get().Debug("got user id", zap.String(UserIdFieldName, username))
			ctx = context.WithValue(ctx, UserIdFieldName, username)

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

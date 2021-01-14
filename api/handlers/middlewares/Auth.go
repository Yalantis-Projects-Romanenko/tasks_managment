package middlewares

import (
	"context"
	"encoding/base64"
	"github.com/fdistorted/task_managment/handlers/common"
	"github.com/fdistorted/task_managment/logger"
	"go.uber.org/zap"
	"net/http"
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
		if !unicode.IsLetter(r) || !unicode.IsNumber(r) {
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
				common.SendRawResponse(w, 401, []byte("no auth token"))
				return
			}

			data, err := base64.StdEncoding.DecodeString(header)
			if err != nil {
				logger.Get().Error("failed to decode token", zap.Error(err))
				common.SendRawResponse(w, 401, []byte("failed to decode token"))
				return
			}

			username := string(data)
			if !CheckUsername(username) {
				common.SendRawResponse(w, 401, []byte("wrong username"))
			}

			logger.Get().Info("got user id", zap.String(UserIdFieldName, username))
			ctx = context.WithValue(ctx, UserIdFieldName, username)

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

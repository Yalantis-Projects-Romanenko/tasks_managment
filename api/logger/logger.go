package logger

import (
	"context"
	"github.com/fdistorted/task_managment/config"
	"go.uber.org/zap"
)

type RequestIdType string

const requestIDKey RequestIdType = "request_id"

var logger *zap.Logger

func Get() *zap.Logger {
	return logger
}

func Load(cfg *config.Config) (err error) {
	logger, err = zap.NewProduction() // todo add log pattern here
	return err
}

func WithCtxValue(ctx context.Context) *zap.Logger {
	return logger.With(zapFieldsFromContext(ctx)...)
}

func zapFieldsFromContext(ctx context.Context) []zap.Field {
	return []zap.Field{
		zap.String(string(requestIDKey), GetRequestID(ctx)),
	}
}

func GetRequestID(ctx context.Context) (value string) {
	value, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return ""
	}

	return value
}

func WithRequestID(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, requestIDKey, value)
}

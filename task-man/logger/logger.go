package logger

import (
	"context"
	"github.com/DimaKuptsov/task-man/config"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"sync"
)

const RequestIDContextField = "request_id"

var loggerInstance *zap.Logger
var once sync.Once

func GetLogger() *zap.Logger {
	return loggerInstance
}

func GetWithContext(context context.Context) *zap.Logger {
	requestId := context.Value(middleware.RequestIDKey).(string)
	return GetLogger().With(
		zap.String(RequestIDContextField, requestId),
	)
}

func Init() error {
	var err error
	once.Do(func() {
		cfg := config.GetConfig()
		switch cfg.Environment {
		case config.ProductionEnvironment:
			loggerInstance, err = zap.NewProduction()
		default:
			loggerInstance, err = zap.NewDevelopment()
		}
	})
	return err
}

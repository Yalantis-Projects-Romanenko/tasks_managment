package main

import (
	"fmt"
	"github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"

	"github.com/fdistorted/task_managment/config"
	"github.com/fdistorted/task_managment/handlers"
)

const defaultConfigPath = "config.json"

// nolint
func main() {
	err := config.Load(defaultConfigPath)
	if err != nil {
		//log.Fatalf("Failed to load config: %s", err.Error())
		log.Printf("Failed to load config: %s", err.Error())
	}

	err = logger.Load(config.Get())
	if err != nil {
		log.Fatalf("Failed to laod logger: %s", err.Error())
	}

	// err = validator.Load()
	// if err != nil {
	// 	logger.Get().Error("Failed to load validator", zap.Error(err))
	// }

	db.NewDb(config.Get().Postgres)

	server := &http.Server{
		Addr:    config.Get().ListenURL,
		Handler: handlers.NewRouter(),
	}

	logger.Get().Info("Listening...", zap.String("listen_url", config.Get().ListenURL))
	err = server.ListenAndServe()
	if err != nil {
		// logger.Get().Error("Failed to initialize HTTP server", zap.Error(err))
		fmt.Println("failed to start server")
		os.Exit(1)
	}
}

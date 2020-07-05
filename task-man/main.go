package main

import (
	"github.com/DimaKuptsov/task-man/config"
	"github.com/DimaKuptsov/task-man/handlers"
	"github.com/DimaKuptsov/task-man/logger"
	"log"
	"net/http"
)

const DefaultConfigFilePath = "config.json"

func main() {
	err := config.InitFromFile(DefaultConfigFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize config: %s", err.Error())
	}
	appConfig := config.GetConfig()
	err = logger.Init()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %s", err.Error())
	}
	router := handlers.NewRouter()
	server := &http.Server{
		Addr:    appConfig.ListenURL,
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}

package main

import (
	"github.com/DimaKuptsov/task-man/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := handlers.NewRouter()
	server := &http.Server{
		Addr:    ":80",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("%s", err.Error())
		os.Exit(1)
	}
}

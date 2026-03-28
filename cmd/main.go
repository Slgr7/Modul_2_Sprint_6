package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("Запуск сервера :8080")
	if err := srv.HTTPServer.ListenAndServe(); err != nil {

		logger.Fatal(err)
	}
}

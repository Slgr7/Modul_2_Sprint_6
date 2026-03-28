package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Ser struct {
	Logger     *log.Logger
	HTTPServer http.Server
}

func NewServer(logger *log.Logger) *Ser {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.FirstHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Ser{
		Logger:     logger,
		HTTPServer: srv,
	}
}

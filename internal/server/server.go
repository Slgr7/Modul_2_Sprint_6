package server

import (
	"log"
	"net/http"
	"time"
)

type Ser struct {
	Logger     *log.Logger
	HTTPServer http.Server
}

func NewServer(logger *log.Logger) *Ser {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

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

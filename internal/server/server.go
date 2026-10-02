package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func StartServer(logger *log.Logger) *Server {
	NewRtr := http.NewServeMux()

	NewRtr.HandleFunc("/", handlers.GetBack)
	NewRtr.HandleFunc("/upload", handlers.Worker)

	myServer := &http.Server{
		Addr:         ":8080",
		Handler:      NewRtr,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: myServer,
	}
}

package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "myServer:", log.LstdFlags)

	srv := server.StartServer(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Printf("Ошибка запуска сервера: %v", err)
	}
}

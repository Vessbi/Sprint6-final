package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

// regIncomRequest регистрация входящих запросов HTTP
func regIncomRequest(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Request: %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// ServerNew создание нового сервера
func ServerNew(logger *log.Logger) *Server {
	router := http.NewServeMux()

	// Обработчик для обслуживания index.html.
	router.HandleFunc("/", handlers.FileHandler)
	// Обработчик для загруженных файлов
	router.HandleFunc("POST /upload", handlers.ParcerHandler)

	// Создание экземпляра структуры http.Server для настройки сервера
	loggedRouter := regIncomRequest(logger, router)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      loggedRouter,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HTTPServer: server,
	}
}

package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger     *log.Logger
	HttpServer *http.Server
}

// New создает новый экземпляр сервера
func New(log *log.Logger) (*Server, error) {
	// Создаем роутер
	mux := http.NewServeMux()

	// Настраиваем HTTP-сервер
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		handlers.IndexHandler(w, req, log)
	})
	mux.HandleFunc("/upload", func(w http.ResponseWriter, req *http.Request) {
		handlers.UploadHandler(w, req, log)
	})

	return &Server{
		logger:     log,
		HttpServer: httpServer,
	}, nil
}

//добьавление хендлеров
//старт и стоп

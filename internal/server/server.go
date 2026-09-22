// Package server отвечает за реализацию функции для создания http-сервера.
package server

import (
	"log"
	"net/http"
	"time"
)

// Настройки HTTP-сервера, числа взяты из задания.
const (
	portAddress  = ":8080" // ":8080" - это адрес прослушивания, пустой хост означает «на всех интерфейсах».
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
	idleTimeout  = 15 * time.Second
)

// APIServer объединяет зависимости приложения, которые живут рядом с сервером.
type APIServer struct {
	Logger *log.Logger
	Server *http.Server
}

// NewAPIServer возвращает незапущенный сервер.
// Вызывающий запускает сервер.
// Логгер попадает в поле ErrorLog, библиотека пишет в логгер сообщения об ошибках.
func NewAPIServer(logger *log.Logger) *APIServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handle)
	httpServer := &http.Server{
		Addr:         portAddress,
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
	return &APIServer{
		Logger: logger,
		Server: httpServer,
	}
}

// Handle отдаёт один и тот же текст на любой запрос.
func Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("сервер жив"))
}

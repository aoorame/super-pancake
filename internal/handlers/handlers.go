package handlers

import (
	"net/http"
)

// Handle отдаёт один и тот же текст на любой запрос.
func Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("сервер жив"))
}

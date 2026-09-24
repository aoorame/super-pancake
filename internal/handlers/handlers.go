package handlers

import (
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

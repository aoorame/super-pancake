// Package handlers содержит обработчики которые принимают запросы и производят ответы.

package handlers

import (
	"io"
	"log"
	"net/http"
)

// ServeIndex отправляет файл index.html клиенту.
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("myFile") //hdr
	if err != nil {
		http.Error(w, "неверный формат файла", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "поток оборван", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Print(err)
		return
	}
}

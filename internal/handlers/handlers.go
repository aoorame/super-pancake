// Package handlers содержит обработчики которые принимают запросы и производят ответы.
package handlers

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// ServeIndex отправляет файл index.html клиенту.
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, hdr, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "неверный формат файла", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	defer file.Close()

	ext := filepath.Ext(hdr.Filename)

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "невозможно прочитать содержимое файла", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	info, err := service.AutoDecodeOrEncode(string(data))
	if err != nil {
		http.Error(w, "пустой файл", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	err = os.WriteFile(time.Now().UTC().String()+ext, []byte(info), 0644)
	if err != nil {
		http.Error(w, "ошибка в записи файла", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	_, err = w.Write([]byte(info))
	if err != nil {
		log.Print(err)
		return
	}
}

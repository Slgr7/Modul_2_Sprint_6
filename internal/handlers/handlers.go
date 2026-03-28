package handlers

import (
	_ "embed"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

//go:embed index.html
var indexHTML string

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index").Parse(indexHTML))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Неверный формат формы", http.StatusBadRequest)
		return
	}

	var file multipart.File
	var header *multipart.FileHeader
	for _, headers := range r.MultipartForm.File {
		if len(headers) > 0 {
			header = headers[0]
			var err error
			file, err = header.Open()
			if err != nil {
				http.Error(w, "Ошибка открытия файла", http.StatusInternalServerError)
				return
			}
			defer file.Close()
			break
		}
	}
	if file == nil {
		http.Error(w, "Файл не передан", http.StatusBadRequest)
		return
	}

	contents, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}
	if len(contents) == 0 {
		http.Error(w, "Файл пуст", http.StatusBadRequest)
		return
	}

	converted, err := service.Service(string(contents))
	if err != nil {
		http.Error(w, "Ошибка пакета service", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	timeStr := time.Now().UTC().String()

	timeStr = strings.Map(func(r rune) rune {
		if r == ' ' || r == ':' || r == '+' {
			return '_'
		}
		return r
	}, timeStr)
	fileName := timeStr + ext

	if err := os.WriteFile(fileName, []byte(converted), 0644); err != nil {
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(converted))
}

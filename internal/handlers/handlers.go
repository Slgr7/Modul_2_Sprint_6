package handlers

import (
	_ "embed"
	"html/template"
	"io"
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

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Файл не передан", http.StatusBadRequest)
		return
	}
	defer file.Close()

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
		switch r {
		case ' ', ':', '+':
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

package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
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

	convertedString, err := service.Service(string(contents))
	if err != nil {
		http.Error(w, "Ошибка пакета service", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	fileName := time.Now().UTC().String() + ext

	// Запись результата в локальный файл
	err = os.WriteFile(fileName, []byte(convertedString), 0644)
	if err != nil {
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Результат: %s", convertedString)
}

package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func GetBack(w http.ResponseWriter, r *http.Request) {

	file, errOpen := os.Open("../index.html")
	if errOpen != nil {
		http.Error(w, "can't read file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-type", "text/html; charset=utf-8")

	_, errCopy := io.Copy(w, file)
	if errCopy != nil {
		http.Error(w, "read mistace", http.StatusInternalServerError)
		return
	}
}

func Worker(w http.ResponseWriter, r *http.Request) {
	// 1. Парсим html-форму из файла index.html
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "can't parse form", http.StatusInternalServerError)
		return
	}

	// 2. Получить файл из формы (не забудьте его закрыть).
	file, header, errGetFile := r.FormFile("myFile")
	if errGetFile != nil {
		http.Error(w, "can't take file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 3. Прочитать данные из файла
	data, errRead := io.ReadAll(file)
	if errRead != nil {
		http.Error(w, "Error of read", http.StatusInternalServerError)
		return
	}

	// 4. Передать эти данные в функцию автоопределения из пакета service,
	// которую вы создали, чтобы получить переконвертируемую строку
	dataString, errTranslate := service.MorseTranslater(string(data))
	if errTranslate != nil {
		http.Error(w, "Error of read", http.StatusInternalServerError)
		return
	}

	// 5. Создать локальный файл с уникальным названием
	fileName := fmt.Sprintf("%d%s", time.Now().Unix(), filepath.Ext(header.Filename))
	out, errCerate := os.Create(fileName)
	if errCerate != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// 6. Записать в локальный файл результат конвертации строки.
	_, errWrite := out.WriteString(dataString)
	if errWrite != nil {
		http.Error(w, "Error writing at file", http.StatusInternalServerError)
		return
	}

	// 7. Возвращаем результат конвертации строки в HTTP-ответе.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, errWriteHTTP := w.Write([]byte(dataString))
	if errWriteHTTP != nil {
		http.Error(w, "Error writing at HTTP", http.StatusInternalServerError)
		return
	}
}

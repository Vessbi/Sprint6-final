package handlers

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// FileHandler возвращает HTML из файла корневой директории к которой обращаются
func FileHandler(res http.ResponseWriter, req *http.Request) {
	req.Header.Add("Content-Type", "text/html")

	http.ServeFile(res, req, "./index.html")
}

// ParcerHandler парсит и записывает результат конвертации строки
func ParcerHandler(res http.ResponseWriter, req *http.Request) {

	if err := req.ParseMultipartForm(10 << 20); err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}

	// получаем файл из формы
	file, handler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "error when receiving file", http.StatusInternalServerError)
		return
	}
	// закрываем файл
	defer file.Close()
	ext := filepath.Ext(handler.Filename)

	// Прочитать весь полученный файл
	scanner := bufio.NewScanner(file)
	var dataFile string
	for scanner.Scan() {
		dataFile += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		http.Error(res, "error reading file", http.StatusInternalServerError)
		return
	}
	// Конвертируем данные функцией из пакета service
	converted, err := service.ConvMorseString(dataFile)
	if err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}

	//Создание нового файла и пути к нему
	fileName := "index_" + time.Now().UTC().Format("20060102150405") + ext
	filePath := filepath.Join(".", fileName)

	outputFile, err := os.Create(filePath)
	if err != nil {
		http.Error(res, "error creating file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Запись данных во вновь созданный файл.
	if _, err = outputFile.Write([]byte(converted)); err != nil {
		http.Error(res, "error while writing file", http.StatusInternalServerError)
		return
	}

	// Возврат результата пользователю в HTTP-ответе.
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if _, err := res.Write([]byte(converted)); err != nil {
		log.Printf("error sending response: %v", err)
	}
}

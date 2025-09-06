package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, req *http.Request, log *log.Logger) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("The server does not support %s requests", req.Method), http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Println("Reading file error:", err)
		http.Error(w, "Reading file error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, req *http.Request, log *log.Logger) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("The server does not support %s requests", req.Method), http.StatusBadRequest)
		return
	}
	file, handler, err := req.FormFile("myFile")
	if err != nil {
		log.Println("Receiving file error:", err)
		http.Error(w, "Receiving file error", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fileData, err := io.ReadAll(file)
	if err != nil {
		log.Println("Opening file error:", err)
		http.Error(w, "Opening file error", http.StatusInternalServerError)
		return
	}
	convStr := service.ReverseMorse(string(fileData))
	finalFile, err := os.Create(time.Now().UTC().String() + filepath.Ext(handler.Filename))
	if err != nil {
		log.Println("Creation file error:", err)
		http.Error(w, "Creation file error", http.StatusInternalServerError)
		return
	}
	defer finalFile.Close()
	_, err = fmt.Fprint(finalFile, convStr)
	if err != nil {
		log.Println("Writing to file error:", err)
		http.Error(w, "Writing to file error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convStr))
}
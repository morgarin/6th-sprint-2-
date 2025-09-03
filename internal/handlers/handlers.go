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

//корневой эндпоинт "/" возвращает индекс html

//второй эндпоинт "/upload" принимает JSON с текстом или морзе и сохраняет JSON c переведенным текстом или морзе

//принимаем файл парсим строку передаем в service status401
// http.StatusBadRequest
// http.StatusInternalServerError

// Construct absolute path to index.html
//htmlPath := filepath.Join(h.baseDir, "index.html")

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
	file, handler, err := req.FormFile("newFile")
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

/*
//go:embed templates/index.html
var indexHTML embed.FS

// Request represents the JSON payload for text/morse conversion
type Request struct {
    Text string json:"text"
}

// Response represents the JSON response
type Response struct {
    Result string json:"result"
    Error  string json:"error,omitempty"
}

// IndexHandler serves the HTML page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    content, err := indexHTML.ReadFile("templates/index.html")
    if err != nil {
        http.Error(w, "Could not load page", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html")
    w.Write(content)
}

// UploadHandler handles text/morse conversion
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.JSON(w, Response{Error: "Invalid JSON"}, http.StatusBadRequest)
        return
    }

    result, err := service.TextOrMorse(req.Text)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(Response{Error: err.Error()})
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Response{Result: result})
}
*/

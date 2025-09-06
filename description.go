//handlers

//корневой эндпоинт "/" возвращает индекс html

//второй эндпоинт "/upload" принимает JSON с текстом или морзе и сохраняет JSON c переведенным текстом или морзе

//принимаем файл парсим строку передаем в service status401
// http.StatusBadRequest
// http.StatusInternalServerError

// Construct absolute path to index.html
//htmlPath := filepath.Join(h.baseDir, "index.html")

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

// функция которя берет строчку из файла и проверяет ее на наличие кирилицы
// если есть кирилица то текст в морзе, если нет то наоборот

/*
func isTextInput(input string) bool {
	hasLettersOrNumbers := false
	hasNonMorseChars := false

	for _, r := range input {
		switch {
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			hasLettersOrNumbers = true
		case r != '.' && r != '-' && r != ' ' && !unicode.IsSpace(r):
			hasNonMorseChars = true
		}
	}

	return hasLettersOrNumbers || hasNonMorseChars
}
*/
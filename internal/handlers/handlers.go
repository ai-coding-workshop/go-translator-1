package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	homeTemplate   *template.Template
	resultTemplate *template.Template
)

func init() {
	// Parse templates
	homeTemplate = template.Must(template.ParseFiles(filepath.Join("web", "templates", "home.html")))
	resultTemplate = template.Must(template.ParseFiles(filepath.Join("web", "templates", "result.html")))
}

// HomeHandler serves the main web page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Render home template
	if err := homeTemplate.Execute(w, nil); err != nil {
		log.Printf("Error rendering home template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// TranslateHandler processes translation requests from the web form
func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get form values
	text := r.FormValue("text")
	model := r.FormValue("model")

	// For now, return a simple response
	// This will be replaced with actual translation logic
	data := struct {
		Original    string
		Model       string
		Translation string
	}{
		Original:    text,
		Model:       model,
		Translation: "[Translation will appear here]",
	}

	// Render result template
	if err := resultTemplate.Execute(w, data); err != nil {
		log.Printf("Error rendering result template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// APIHandler handles REST API requests for translation
func APIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// For now, return a simple JSON response
	// This will be replaced with actual translation logic
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(`{
    "status": "success",
    "message": "API endpoint ready for translation",
    "translation": "[Translation will appear here]",
    "data": {
        "original": "[text to translate]",
        "model": "[selected model]"
    }
}`))
	if err != nil {
		log.Printf("Error writing API response: %v", err)
	}
}

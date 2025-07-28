package handlers

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"translator-service/internal/models"
	"translator-service/internal/services"
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
func NewHomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

// TranslateHandler processes translation requests from the web form
type TranslateHandler struct {
	translatorService *services.TranslatorService
}

func NewTranslateHandler(translatorService *services.TranslatorService) http.HandlerFunc {
	handler := &TranslateHandler{
		translatorService: translatorService,
	}

	return handler.ServeHTTP
}

func (h *TranslateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get form values
	text := r.FormValue("text")
	model := r.FormValue("model")

	// Create translation request
	req := &models.TranslationRequest{
		Text:  text,
		Model: model,
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	// Perform translation
	response, err := h.translatorService.Translate(ctx, req)
	if err != nil {
		log.Printf("Translation error: %v", err)
		http.Error(w, "Translation failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Render result template
	if err := resultTemplate.Execute(w, response); err != nil {
		log.Printf("Error rendering result template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// APIHandler handles REST API requests for translation
type APIHandler struct {
	translatorService *services.TranslatorService
}

func NewAPIHandler(translatorService *services.TranslatorService) http.HandlerFunc {
	handler := &APIHandler{
		translatorService: translatorService,
	}

	return handler.ServeHTTP
}

func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON request
	var req models.TranslationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Text == "" {
		http.Error(w, "Text field is required", http.StatusBadRequest)
		return
	}

	if req.Model == "" {
		http.Error(w, "Model field is required", http.StatusBadRequest)
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	// Perform translation
	response, err := h.translatorService.Translate(ctx, &req)
	if err != nil {
		log.Printf("Translation error: %v", err)
		http.Error(w, "Translation failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

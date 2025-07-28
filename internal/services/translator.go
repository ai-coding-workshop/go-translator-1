package services

import (
	"context"
	"fmt"
	"log"

	"translator-service/internal/models"
)

// TranslatorService manages multiple translation providers
type TranslatorService struct {
	translators map[string]models.Translator
}

// NewTranslatorService creates a new translator service
func NewTranslatorService() *TranslatorService {
	service := &TranslatorService{
		translators: make(map[string]models.Translator),
	}

	// Register supported translators
	service.registerTranslators()

	return service
}

// registerTranslators registers all supported translation providers
func (ts *TranslatorService) registerTranslators() {
	// For now, we'll register mock translators
	// In a real implementation, these would be actual LLM API clients
	ts.translators["gpt-3.5"] = NewMockTranslator("GPT-3.5")
	ts.translators["gpt-4"] = NewMockTranslator("GPT-4")
	ts.translators["claude"] = NewMockTranslator("Claude")
	ts.translators["llama"] = NewMockTranslator("Llama")
}

// Translate translates text using the specified model
func (ts *TranslatorService) Translate(ctx context.Context, req *models.TranslationRequest) (*models.TranslationResponse, error) {
	// Find the appropriate translator
	translator, exists := ts.translators[req.Model]
	if !exists {
		return nil, fmt.Errorf("unsupported model: %s", req.Model)
	}

	// Perform translation
	response, err := translator.Translate(ctx, req)
	if err != nil {
		log.Printf("Translation error with %s: %v", req.Model, err)
		return nil, fmt.Errorf("failed to translate with %s: %w", req.Model, err)
	}

	return response, nil
}

// GetSupportedModels returns a list of supported models
func (ts *TranslatorService) GetSupportedModels() []string {
	models := make([]string, 0, len(ts.translators))
	for model := range ts.translators {
		models = append(models, model)
	}
	return models
}

// IsModelSupported checks if a model is supported
func (ts *TranslatorService) IsModelSupported(model string) bool {
	_, exists := ts.translators[model]
	return exists
}

package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler serves the main web page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Simple response for now, will be replaced with proper HTML template
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Translation Service</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" type="text/css" href="/static/css/style.css">
</head>
<body>
    <h1>Translation Service</h1>
    <form action="/translate" method="POST">
        <label for="text">Enter English text to translate:</label><br><br>
        <textarea id="text" name="text" rows="4" cols="50" placeholder="Enter English word or sentence"></textarea><br><br>

        <label for="model">Select LLM Model:</label>
        <select id="model" name="model">
            <option value="gpt-3.5">GPT-3.5</option>
            <option value="gpt-4">GPT-4</option>
            <option value="claude">Claude</option>
            <option value="llama">Llama</option>
        </select><br><br>

        <input type="submit" value="Translate">
    </form>
</body>
</html>
`)
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
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Translation Result</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" type="text/css" href="/static/css/style.css">
</head>
<body>
    <h1>Translation Result</h1>
    <p><strong>Original:</strong> %s</p>
    <p><strong>Model:</strong> %s</p>
    <p><strong>Translation:</strong> [Translation will appear here]</p>
    <a href="/">Translate Another</a>
</body>
</html>
`, text, model)
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
	fmt.Fprintf(w, `{
    "status": "success",
    "message": "API endpoint ready for translation",
    "data": {
        "original": "[text to translate]",
        "translation": "[translation will appear here]",
        "model": "[selected model]"
    }
}`)
}

package controllers

import (
	"fmt"
	"github.com/NerdsvilleCEO/go-chuck-norrisify/models"
	"net/http"
)

// GET /
// Success: 200
// Error: 400
// Content-Type: application/json
// Data: {"result": "JOKE", "error": []}
// TODO: Add further error handling middleware
func (app *App) GetJokeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var errors []error
	name, err := models.GetName()
	if err != nil {
		errors = append(errors, err)
	}

	joke, err := models.GetJoke(name)
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"result": "", "error": %q}`, errors)
		return
	}

	fmt.Fprintf(w, `{"result": %s, "error": %s}`, joke.Value["joke"], errors)
}

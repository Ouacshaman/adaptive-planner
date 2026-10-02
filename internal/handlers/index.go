package handlers

import (
	"adaptive-planner/internal/app"
	"log"
	"net/http"
)

func HandlersIndex(cfg *app.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := cfg.PageTemplate.Execute(w, nil); err != nil {
			log.Printf("Error rendering page template: %s", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}

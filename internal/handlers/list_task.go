package handlers

import (
	"adaptive-planner/internal/app"
	"net/http"
)

func HandlersListTask(cfg *app.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Task string `json:"task"`
		}
	}
}

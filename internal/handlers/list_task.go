package handlers

import (
	"adaptive-planner/internal/app"
	"net/http"
)

func (cfg *app.APIConfig) handlersListTask(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Task string `json:"task"`
	}
}

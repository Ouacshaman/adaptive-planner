package handlers

import (
	"adaptive-planner/internal/app"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type task struct {
	Id        uuid.UUID `json:"id"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Task      string    `json:"task"`
}

func HandlersListTask(cfg *app.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Task string `json:"task"`
		}
	}
}

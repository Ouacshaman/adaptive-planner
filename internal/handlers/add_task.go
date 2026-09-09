package handlers

import (
	"adaptive-planner/internal/app"
	"adaptive-planner/internal/response"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func HandlersAddTask(cfg *app.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Task string `json:"task"`
		}

		type Task struct {
			ID        uuid.UUID `json:"id"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			Task      string    `json:"task"`
		}

		params := parameters{}

		err := response.DecodeJson(w, r, &params)
		if err != nil {
			return
		}

		task, err := cfg.DB.CreateTask(r.Context(), params.Task)
		if err != nil {
			return
		}
		response.RespondWithJSON(w, 200, Task{
			ID:        task.ID,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
			Task:      task.Task,
		})
	}
}

package handlers

import (
	"adaptive-planner/internal/app"
	"adaptive-planner/internal/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Task      string    `json:"task"`
}

func HandlersListTask(cfg *app.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbTasks, err := cfg.DB.GetTasks(r.Context())
		if err != nil {
			return
		}

		tasks := []Task{}
		for _, dbTask := range dbTasks {
			tasks = append(tasks, Task{
				ID:        dbTask.ID,
				CreatedAt: dbTask.CreatedAt,
				UpdatedAt: dbTask.UpdatedAt,
				Task:      dbTask.Task,
			})
		}

		json.RespondWithJSON(w, 200, tasks)
	}
}

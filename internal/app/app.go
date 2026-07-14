package app

import "adaptive-planner/internal/database"

type APIConfig struct {
	DB       *database.Queries
	Platform string
}

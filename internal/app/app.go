package app

import (
	"adaptive-planner/internal/database"
	"html/template"
)

type APIConfig struct {
	DB           *database.Queries
	Platform     string
	PageTemplate *template.Template
}

package main

import (
	"adaptive-planner/internal/app"
	"adaptive-planner/internal/database"
	"adaptive-planner/internal/handlers"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	const filepathRoot = "."
	const port = "8080"
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	dbQueries := database.New(dbConn)

	apiCfg := app.APIConfig{
		DB:       dbQueries,
		Platform: platform,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", http.FileServer(http.Dir(".")))

	mux.HandleFunc("GET /api/tasks", handlers.HandlersListTask(&apiCfg))
	mux.HandleFunc("POST /api/tasks", handlers.HandlersAddTask(&apiCfg))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())

}

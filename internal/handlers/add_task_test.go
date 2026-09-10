package handlers

import (
	"adaptive-planner/internal/app"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlersAddTask(t *testing.T){

	type params struct{
		Task string `json:"task"`
	}

	entry := params{Task: "Brush teeth"}

	data, err := json.Marshal(entry)
	if err != nil{
		t.Fatalf("Unable to marshal json: %v",err)
	}

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(data))

	recorder := httptest.NewRecorder()

	cfg := app.APIConfig{}

	handler := HandlersAddTask(&cfg)

	handler(recorder, req)
}

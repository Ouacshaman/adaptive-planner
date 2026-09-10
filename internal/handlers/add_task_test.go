package handlers

import (
	"adaptive-planner/internal/app"
	"bytes"
	"encoding/json"

	//	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlersAddTask(t *testing.T){

//	type params struct{
//		Task string `json:"task"`
//	}
//
//	entry := params{Task: "Brush teeth"}
//
//	data, err := json.Marshal(entry)
//	if err != nil{
//		t.Fatalf("Unable to marshal json: %v",err)
//	}
	
	malformedJSON := `{"app": "joe,`

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(malformedJSON)))

	recorder := httptest.NewRecorder()

	cfg := app.APIConfig{}

	handler := HandlersAddTask(&cfg)

	handler(recorder, req)

	if recorder.Code != 400{
		t.Fatalf("Code returned %d, instead of expected %d", recorder.Code, 400)
	}

	type params struct{
		Error string `json:"error"`
	}

	out := params{}

	data := json.Unmarshal(recorder.Body.Bytes(), &out)

	if data.Error() != "Json is unreadable"{
		t.Fatalf("Expected %q got %q", "Json is unreadable", data.Error())
	}

}

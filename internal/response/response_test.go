package response

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	type TestJson struct {
		Name string `json:"name"`
	}

	val := TestJson{
		Name: "Joe",
	}

	data, err := json.Marshal(val)
	if err != nil {
		t.Fatalf("failed to marshal json: %v", err)
	}

	expected := "Joe"

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(data))
	recorder := httptest.NewRecorder()

	params := TestJson{}

	err = DecodeJson(recorder, req, &params)

	if err != nil {
		t.Fatalf("failed to decode json: %v", err)
	}

	if params.Name != expected {
		t.Errorf("DecodeJson() failed, got %q, want %q", params.Name, expected)
	}

}

func TestRespondWithError(t *testing.T) {

	code := 404
	msg := "error"
	err := errors.New("test error")

	recorder := httptest.NewRecorder()

	RespondWithError(recorder, code, msg, err)

	fmt.Println(recorder.Body.String())

	type param struct {
		Error string `json:"error"`
	}

	res := param{}

	json.Unmarshal(recorder.Body.Bytes(), &res)

	if recorder.Code != code {
		t.Errorf("ResponseWithError Code is not valid received %d, expected %d", recorder.Code, code)
	}
	if res.Error != msg {
		t.Errorf("ResponseWithError msg failed got %q and %q", recorder.Body.String(), msg)
	}
}

func TestRespondWithJson(t *testing.T) {
	type param struct {
		Test string `json:"test"`
	}

	code := 200
	pl := param{Test: "test00"}

	rec := httptest.NewRecorder()

	RespondWithJSON(rec, code, pl)

	if rec.Code != code {
		t.Errorf("RespondWithJson code is not valid got %d, want %d", rec.Code, code)
	}

	res := param{}

	json.Unmarshal(rec.Body.Bytes(), &res)

	if res.Test != pl.Test {
		t.Errorf("RespondWithJson Payload is not right got %q, want %q", res.Test, pl.Test)
	}

}

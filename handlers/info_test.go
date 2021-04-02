package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestHandler_Info(t *testing.T) {

	handler := NewHandler()

	req := httptest.NewRequest("GET", "/health", nil)
	res := httptest.NewRecorder()

	h := handler.Info()
	h(res, req, nil)

	result := res.Body.String()
	expected := `{"version": "0.0.1"}`

	if res.Code != 200 {
		t.Errorf("Expected HTTP response 200 but got %d", res.Code)
	}

	if result != expected {
		t.Errorf("Expected body to contain value %q but got %q", expected, result)
	}
}

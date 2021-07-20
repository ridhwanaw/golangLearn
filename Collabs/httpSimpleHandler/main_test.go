package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	recorder := httptest.NewRecorder()

	requester, err := http.NewRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	hello(recorder, requester)

	result := recorder.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d\n", http.StatusOK, result.StatusCode)
	}

	defer result.Body.Close()

	body, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Errorf("want body, got error %q\n", err)
	}

	if string(body) != "Hello" {
		t.Errorf("want body 'Hello', got %s\n", string(body))
	}
}

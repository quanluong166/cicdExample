package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleTodos(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(`{"task":"Learn CI/CD"}`))
	w := httptest.NewRecorder()
	handleTodos(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/todos", nil)
	w = httptest.NewRecorder()
	handleTodos(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

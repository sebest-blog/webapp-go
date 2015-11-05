package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPath(t *testing.T) {
	h := http.HandlerFunc(handler)
	req, _ := http.NewRequest("GET", "/path", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/path didn't return %v\n", http.StatusOK)
	}
	if w.Body.String() != "Request /path (go)\n" {
		t.Errorf("/path didn't return the expected body: %s\n", w.Body.String())
	}
}

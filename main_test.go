package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandleWhenOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecoder := httptest.NewRecorder()
	handle := http.HandlerFunc(mainHandle)
	handle.ServeHTTP(responseRecoder, req)

	if status := responseRecoder.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}
}

func TestMainHandleWhenMissingCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecoder := httptest.NewRecorder()
	handle := http.HandlerFunc(mainHandle)
	handle.ServeHTTP(responseRecoder, req)

	if status := responseRecoder.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, status)
	}

	expected := `count missing`
	if responseRecoder.Body.String() != expected {
		t.Errorf("expected body: %s, got %s", expected, responseRecoder.Body.String())
	}
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, status)
	}

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	if len(list) != totalCount {
		t.Errorf("expected cafe count: %d, got %d", totalCount, len(list))
	}
}

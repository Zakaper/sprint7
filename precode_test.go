package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	/*
	   if status := responseRecorder.Code; status != http.StatusOK {
	       t.Errorf("expected status code: %d, got %d", http.StatusOK, status)
	   }
	*/
}

func TestMainHandlerWhenMissingCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=omsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), `count missing`)
	/*
		if status := responseRecorder.Code; status != http.StatusBadRequest {
			t.Errorf("expected status code: %d, got %d", http.StatusBadRequest, status)
		}

		expected := `count missing`
		if responseRecorder.Body.String() != expected {
			t.Errorf("expected body: %s, got %s", expected, responseRecorder.Body.String())
		}
	*/
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//if status := responseRecorder.Code; status != http.StatusOK {
	//	t.Fatalf("expected status code: %d, got %d", http.StatusOK, status)
	//}

	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, len(list), totalCount)
}

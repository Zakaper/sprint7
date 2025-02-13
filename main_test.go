package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 2
	URL := fmt.Sprintf("/cafe?count=%d&city=moscow", totalCount)
	req := httptest.NewRequest("GET", URL, nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := "Мир кофе,Сладкоежка"
	assert.NotEmpty(t, responseRecorder.Body.String())
	assert.Equal(t, responseRecorder.Body.String(), expected)
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
}
func TestMainHandlerWhenCountMoreThanTotal2(t *testing.T) {
	totalCount := 4
	URL := fmt.Sprintf("/cafe?count=%d&city=moscow2", totalCount)
	req := httptest.NewRequest("GET", URL, nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := "wrong city value"
	assert.NotEmpty(t, responseRecorder.Body.String())
	assert.Equal(t, responseRecorder.Body.String(), expected)
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
}
func TestMainHandlerWhenCountMoreThanTotal3(t *testing.T) {
	totalCount := 6
	URL := fmt.Sprintf("/cafe?count=%d&city=moscow", totalCount)
	req := httptest.NewRequest("GET", URL, nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.NotEmpty(t, responseRecorder.Body.String())
	assert.Equal(t, responseRecorder.Body.String(), expected)
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
}

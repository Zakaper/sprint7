package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест №1: Корректный запрос (код 200, тело ответа не пустое)
func TestMainHandlerCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=5", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected code: 200. Actual code: %d", responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String(), "Value should not be empty")
}

// Тест №2: Город не поддерживается (код 400, сообщение об ошибке)
func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=tula&count=7", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected code: 400. Actual code: %d", responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Expected: wrong city value. Actual: %s", responseRecorder.Body.String())
}

// Тест №3: Количество кафе больше, чем доступно (возвращаются все кафе)
func TestMainHandlerCountExceeds(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=10", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	actualCount := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, totalCount, len(actualCount), "Expected count: %d. Actual code: %d", totalCount, len(actualCount))
}

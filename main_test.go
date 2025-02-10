package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Проверка статуса ответа и тела ответа
func TestMainHandler_StatusOkAndBodyNotEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

// Проверка неподдерживаемого города
func TestMainHandler_WrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=wocsom", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

// Проверка количества кафе
func TestMainHandler_WhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=100&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	checkCafeString := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"

	list := strings.Split(responseRecorder.Body.String(), ",")

	assert.GreaterOrEqual(t, totalCount, (len(list)))
	assert.Equal(t, responseRecorder.Body.String(), checkCafeString)
}

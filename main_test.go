package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест 1: если count больше чем доступно, вывести все кафе
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=8&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	// проверяем код ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	// проверяем , что вывелись все кафе
	body := strings.Join(cafeList["moscow"][:totalCount], ",")
	assert.Equal(t, body, responseRecorder.Body.String())
}

// Тест 2: запрос сформирован корректно, код ответа - 200, и тело ответа не пустое
func TestStatusOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=8&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверяем код ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	// проверяем что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body.String())
}

// Тест 3: указанный город не поддерживается
func TestInvalidCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=8&city=omsk", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверяем код ответа
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// проверяем сообщение об ошибке
	body := "wrong city value"
	assert.Equal(t, body, responseRecorder.Body.String())
}

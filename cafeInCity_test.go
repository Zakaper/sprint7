package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandleRequestOk(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil) // запрос с корректными параметрами

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// сервис вернет код 200 и тело ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())

	expectedResponce := "Мир кофе, Сладкоежка, Кофе и завтраки, Сытый студент"
	assert.Equal(t, expectedResponce, responseRecorder.Body.String())

}

func TestMainHandleCityWrong(t *testing.T) {

	expectedBody := "wrong city value"

	req := httptest.NewRequest("GET", "/cafe?count=4&city=Gotem", nil) // некорректный параметр city

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// сервис вернет код 400 и ошибку в теле ответа
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, expectedBody, responseRecorder.Body.String())

}

func TestMainHandleCountMore(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // некорректный параметр count >4

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// сервис вернет все доступные кафе, если параметр count больше, чем всего кафе
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	actualResponce := strings.Split(responseRecorder.Body.String(), ", ")
	assert.Equal(t, totalCount, len(actualResponce))
}

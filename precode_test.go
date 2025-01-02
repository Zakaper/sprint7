package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// проверяем, что сервис возвращает код ответа 200 и тело ответа не пустое
	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// проверяем, что в случае, если в параметре count указано количество кафе,
	// превыщающее total count, то возвращаются все доступные кафе
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Equal(t, len(list), totalCount)
}

func TestMainHandlerWhenCityNotIncluded(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=tver", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// проверяем, что в случае, если город, который передаётся в параметре city, не поддерживается,
	// то сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа
	status := responseRecorder.Code
	body := responseRecorder.Body.String()
	assert.Equal(t, status, http.StatusBadRequest)
	assert.Equal(t, body, "wrong city value")
}

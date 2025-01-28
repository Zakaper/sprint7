package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидается код ответа 200")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Ответ не должен быть пустым")

	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, cafes, 4, "Количество кафе должно соответствовать максимальному количеству доступных")
}

func TestMainHandlerWithUnsupportedCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=novosibirsk", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидается код ответа 400")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Тело ответа должно содержать ошибку 'wrong city value'")
}

func TestMainHandlerCorrectRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидается код ответа 200")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Ответ не должен быть пустым")

	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, cafes, 2, "Количество кафе должно быть равно 2")
}

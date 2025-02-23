package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=8&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// здесь нужно добавить необходимые проверки

	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusOK)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, "пустой запрос")

	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)
}
func TestMainHandlerGood(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// здесь нужно добавить необходимые проверки

	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusOK)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, "пустой запрос")

	assert.Equal(t, body, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент")
}

func TestMainHandlerBad(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=rostov", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// здесь нужно добавить необходимые проверки

	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusBadRequest)

	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body)
}

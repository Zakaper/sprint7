package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	require.NotEmpty(t, body)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

}

func TestMainHandlerWhenWrongCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscownew", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	bodyEcpected := "wrong city value"
	bodyActual := responseRecorder.Body.String()

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, bodyEcpected, bodyActual)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	// здесь нужно создать запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//   здесь нужно добавить необходимые проверки

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()

	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)

}

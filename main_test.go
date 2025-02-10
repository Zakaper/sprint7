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

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)

}

func TestMainHandlerWhenNoCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=tula", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	ecpected := 400
	bodyEcpected := "wrong city value"
	require.Equal(t, http.StatusBadRequest, ecpected)
	assert.Equal(t, responseRecorder.Body.String(), bodyEcpected)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=12&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Equal(t, len(list), totalCount)

}

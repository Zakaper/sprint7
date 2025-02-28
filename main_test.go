package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenGoodRequest(t *testing.T) {
	URL := fmt.Sprintf("/cafe?count=%d&city=moscow", 4)
	req := httptest.NewRequest("GET", URL, nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
	require.Equal(t, responseRecorder.Code, http.StatusOK)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	URL := fmt.Sprintf("/cafe?count=%d&city=spb", 4)
	req := httptest.NewRequest("GET", URL, nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	expected := "wrong city value"
	assert.Equal(t, body, expected)
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	URL := fmt.Sprintf("/cafe?count=%d&city=moscow", 5)
	req := httptest.NewRequest("GET", URL, nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)
}

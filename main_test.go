package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerStatusOkBodyNotEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerCityNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=Moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerCountMoreThanTotal(t *testing.T) {

	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=7&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	cnt := responseRecorder.Body.String()
	answerCount := len(strings.Split(cnt, ","))

	assert.Equal(t, totalCount, answerCount)
	require.Equal(t, totalCount, answerCount)
}

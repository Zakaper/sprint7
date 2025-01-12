package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("Get", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, len(list), totalCount)
	assert.LessOrEqual(t, t, totalCount)

	status := responseRecorder.Code

	require.Equal(t, http.StatusOK, status)
}

func TestMainAnswer200Ok(t *testing.T) {
	req := httptest.NewRequest("Get", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	body := responseRecorder.Body

	require.Equal(t, http.StatusOK, status)
	assert.NotEmpty(t, body)
}

func TestСity(t *testing.T) {
	req := httptest.NewRequest("Get", "/cafe?count=4&city=murmansk", nil)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	require.Equal(t, http.StatusBadRequest, status)

	body := responseRecorder.Body.String()

	assert.Equal(t, body, "wrong city value")
}

package main

import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.NotEmpty(t, responseRecorder)

    body := responseRecorder.Body.String()
    list := strings.Split(body, ",")
    assert.Len(t, list, totalCount)
}

func TestMainHandlerWhenOk(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.NotEmpty(t, responseRecorder.Body.String())

    body := responseRecorder.Body.String()
    list := strings.Split(body, ",")
    assert.Len(t, list, 2)
}

func TestMainHandlerWhenBad(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?city=alibaba", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.NotEmpty(t, responseRecorder.Body.String())

    expected := `count missing`
    require.Equal(t, expected, responseRecorder.Body.String())
}

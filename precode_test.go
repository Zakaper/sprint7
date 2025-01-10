package main

import (
    "net/http"
    "strings"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
)
func TestMainHandlerStatusOk(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)
    assert.Equal(t, http.StatusOK, responseRecorder.Code, "Полученный статус не соответствует 200")

    body := responseRecorder.Body.String()
    assert.NotEmpty(t,body, "Тело ответа не должно быть пустым")
    
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)
    assert.Equal(t, http.StatusOK, responseRecorder.Code, "Полученный статус не соответствует 200")

    body := responseRecorder.Body.String()
    assert.NotEmpty(t,body, "Тело ответа не должно быть пустым")
    list := strings.Split(body, ",")
    assert.Len(t, list, totalCount)
}
func TestCorrectCity(t *testing.T){
req := httptest.NewRequest("GET", "/cafe?count=2&city=omsk", nil)
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)
    assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

    body := responseRecorder.Body.String()
    assert.Equal(t,"wrong city value", body, "Города нет в базе")

} 
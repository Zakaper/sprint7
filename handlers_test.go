package main

import (
	"net/http/httptest"
	"testing"
	"net/http"
	"strings"
	"github.com/stretchr/testify/assert"
)

// тест для случая если переданное число count больше чем длина слайса
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := len(cafeList["moscow"])
    // создаем запрос к серверу через httptest, укажем count = 10
    req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

    // получаем объект ResponseRecorder, который содержит ответ от сервера
    responseRecorder := httptest.NewRecorder()
    // создаем хендлер и указываем какой обработчик
    handler := http.HandlerFunc(mainHandle)
    // делаем запрос и указываем куда записываем ответ (responseRecorder) и что за запрос (req)
    handler.ServeHTTP(responseRecorder, req)

    // проверяем полученную информацию
    expected := strings.Join(cafeList["moscow"], ",")
    receivedStr := responseRecorder.Body.String()

    // проверяем длину полученного слайса
    receivedLen := len(strings.Split(receivedStr, ","))
    
	// проверяем на равенство
    assert.Equal(t, expected, receivedStr)
    assert.Equal(t, totalCount, receivedLen)
}

func TestMainHandlerWhenCityIsInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=milan", nil)
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

	// проверяем полученное тело
	expectedBody := "wrong city value"
	receivedBody := responseRecorder.Body.String()

	// получаем статус код ответа
	receivedCode := responseRecorder.Code

	assert.Equal(t, expectedBody, receivedBody)
	assert.Equal(t, http.StatusBadRequest, receivedCode)
}

func TestMainHandlerWhenRequestIsValid(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

	// проверяем полученное тело
	receivedBody := responseRecorder.Body.String()

	// получаем статус код ответа
	receivedCode := responseRecorder.Code

	assert.NotEmpty(t, receivedBody)
	assert.Equal(t, http.StatusOK, receivedCode)
}
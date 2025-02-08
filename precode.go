package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=5", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался код 200")

	assert.NotEmpty(t, responseRecorder.Body.String(), "Ответ сервера не должен быть пустым")

	expectedLength := len("Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент")
	assert.Len(t, responseRecorder.Body.String(), expectedLength, "Длина ответа отличается от ожидаемой")

	expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Ответ сервера отличается от ожидаемого")
}

func TestMainHandlerWhenRequestIsValid(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался код 200")

	assert.NotEmpty(t, responseRecorder.Body.String(), "Ответ сервера не должен быть пустым")

	expectedBody := "Мир кофе,Сладкоежка"
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Ответ сервера отличается от ожидаемого")
}

func TestMainHandlerWhenCityIsInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=ufa&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидался код 400")

	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Ошибка в ответе сервера")
}

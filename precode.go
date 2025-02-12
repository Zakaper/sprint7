package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var cafeList = map[string][]string{
	"moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

// Ваш основной обработчик запросов
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

// Тест 1: Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
}

// Тест 2: Город, который передается в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=spb", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	body := responseRecorder.Body.String()
	require.Contains(t, body, "wrong city value")
}

// Тест 3: Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	body := responseRecorder.Body.String()
	expectedCafes := strings.Join(cafeList["moscow"], ",")
	require.Equal(t, expectedCafes, body)
}

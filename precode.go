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

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK) // проверил что вернул 200
	body := responseRecorder.Body.String()                 // перевел тело ответа в строку
	assert.NotEmpty(t, body)                               // проверяю что тело в ответе не пустое

}

func TestMainHandlerWhenNoCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=bryansk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	mistake := 400 // указываю номер ошибки
	bodyEcpected := "wrong city value"
	require.Equal(t, http.StatusBadRequest, mistake)              // сравниваю код ошибки
	assert.Equal(t, responseRecorder.Body.String(), bodyEcpected) // сравниваю город и вывожу тело ответа если город некорректный

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=6&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()      // Получаю тело ответа (кол-во городов)
	splitBody := strings.Split(body, ",")       // Привожу тело ответа в слайс по запятой
	assert.Equal(t, len(splitBody), totalCount) // Проверяю равенство индексов в получившимся слайсе и вывожу весь список

}

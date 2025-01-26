package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
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
	// Создание запроса с параметрами "count" и "city".
	req := httptest.NewRequest("GET", "/?city=moscow&count=10", nil)

	// Создание ResponseRecorder для записи ответа.
	responseRecorder := httptest.NewRecorder()

	// Создание обработчика.
	handler := http.HandlerFunc(mainHandle)

	// Выполнение запроса.
	handler.ServeHTTP(responseRecorder, req)

	// Проверка статуса ответа.
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Проверка содержимого ответа.
	expected := "Мир кофе, Сладкоежка, Кофе и завтраки, Сытый студент"
	if responseRecorder.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), expected)
	}
}

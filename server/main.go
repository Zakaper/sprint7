package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// cafeList - мапа, где указатель - город, содержимое - перечень кафе в этом городе
var cafeList = map[string][]string{
	"moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

// mainHandle обрабатывает запросы на сервер по адресу "/cafe"
func mainHandle(w http.ResponseWriter, req *http.Request) {

	// countStr содержит значение заголовка "count" типа string
	countStr := req.URL.Query().Get("count")

	// проверяем, не пустая ли строка countStr
	// если пустая, передаем код ошибки 400 и соответствующее сообщение
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	// count содержит конвертированное значение countStr в int
	// если конвертация не удалась, передаем код ошибки 400 и соответствующее сообщение
	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	// city содержит значение заголовка "city"
	city := req.URL.Query().Get("city")

	// проверяем наличие города в нашей мапе
	// если отсутствует, передаем код ошибки 400 и соответствующее сообщение
	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	// если запрашиваемое кол-во кафе больше, чем в нашей мапе
	// возвращаем кол-во из нашей мапы
	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

func main() {
	http.HandleFunc(`/cafe`, mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("failed to listen and serve: %s\n", err)
	}
}

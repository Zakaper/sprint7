package main

import (
    "net/http"
    "strconv"
    "strings"
)

// Список кафе для разных городов
var cafeList = map[string][]string{
    "moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

// Обработчик HTTP-запросов
func mainHandle(w http.ResponseWriter, req *http.Request) {
    countStr := req.URL.Query().Get("count")
    if countStr == "" {
        http.Error(w, "count missing", http.StatusBadRequest)
        return
    }

    count, err := strconv.Atoi(countStr)
    if err != nil {
        http.Error(w, "wrong count value", http.StatusBadRequest)
        return
    }

    city := req.URL.Query().Get("city")
    cafe, ok := cafeList[city]
    if !ok {
        http.Error(w, "wrong city value", http.StatusBadRequest) // Используем http.Error (добавляет \n)
        return
    }

    if count > len(cafe) {
        count = len(cafe)
    }

    answer := strings.Join(cafe[:count], ",")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(answer))
}

func main() {
    http.HandleFunc("/cafe", mainHandle)
    http.ListenAndServe(":8080", nil)
}

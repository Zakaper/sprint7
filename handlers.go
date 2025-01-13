package main

import (
    "net/http"
    "strconv"
    "strings"
)

var cafeList = map[string][]string{
    "moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
    // получаем параметр count из запроса
    countStr := req.URL.Query().Get("count")
    if countStr == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("count missing"))
        return
    }

    // переводим count в int
    count, err := strconv.Atoi(countStr)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("wrong count value"))
        return
    }

    // получаем city из параметров запроса
    city := req.URL.Query().Get("city")

    // проверяем наличие города в мапе cafeList
    cafe, ok := cafeList[city]
    if !ok {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("wrong city value"))
        return
    }

    // если число в запросе превысило длину слайса, то вернуть приравнять их
    if count > len(cafe) {
        count = len(cafe)
    }

    // ответ будет строкой со значениями из слайса через запятую
    answer := strings.Join(cafe[:count], ",")

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(answer))
}



package main

import (
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "testing"
)

var cafeList = map[string][]string{
    "moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
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

// Тест 1: Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerCorrectRequest(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
    if err != nil {
        t.Fatal(err)
    }

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    if status := responseRecorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    if responseRecorder.Body.String() == "" {
        t.Errorf("handler returned empty body")
    }
}

// Тест 2: Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWrongCityValue(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?city=london&count=2", nil)
    if err != nil {
        t.Fatal(err)
    }

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    if status := responseRecorder.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    expected := "wrong city value"
    if responseRecorder.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            responseRecorder.Body.String(), expected)
    }
}

// Тест 3: Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerCountMoreThanTotal(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
    if err != nil {
        t.Fatal(err)
    }

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    if status := responseRecorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
    if responseRecorder.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            responseRecorder.Body.String(), expected)
    }
}

func main() {}

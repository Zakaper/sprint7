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

	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=52&city=Petropavlovsk-Kamchatsky", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	status := responseRecorder.Code
	body := responseRecorder.Body.String()

	// "Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа."
	expectedNoCity := `wrong city value`
	require.Equal(t, http.StatusBadRequest, status, "status is not 400")
	require.Equal(t, expectedNoCity, body, "body is not 'wrong city value'")

	// "Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе."
	list := strings.Split(body, ",")
	assert.Equal(t, totalCount, len(list), "incorrect count in the response")

	// "Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое"
	assert.Equal(t, http.StatusOK, status, "status is not 200")
	assert.NotEmpty(t, body, "response body is empty")
}

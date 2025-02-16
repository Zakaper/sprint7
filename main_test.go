package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	require.NotEmpty(t, body, "Body Empty") // Проверяю не пустое ли тело

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Return code status is not 200") // проверка статус кода на 200

}

func TestMainHandlerWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=isengard", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	assert.Equal(t, "wrong city value", body, "Retutn not as expexted")                            //Провека на несуществующий город
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Return code status is not 400") // Проверка статус кода на 400

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	testCount := totalCount + 1

	requestURL := fmt.Sprintf("/cafe?count=%d&city=moscow", testCount)

	req := httptest.NewRequest("GET", requestURL, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	body := responseRecorder.Body.String()
	itemsInBody := strings.Split(body, ",")

	assert.Len(t, itemsInBody, totalCount, "More values returned than expected") //Проверка на количество возвращенных элементов
}

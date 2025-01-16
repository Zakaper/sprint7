package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorrectRequest(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())

	expectedResponce := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedResponce, responseRecorder.Body.String())

}

func TestMainHandlerCity(t *testing.T) {

	expectedBody := "wrong city value"

	req := httptest.NewRequest("GET", "/cafe?count=4&city=failedcity", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, expectedBody, responseRecorder.Body.String())

}

func TestMainHandlerCountMore(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=999&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	actualResponce := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, totalCount, len(actualResponce))
}

/*Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
	list := strings.Split(body, ",")

	if !assert.Equal(t, totalCount, len(list)) {
		t.Errorf("expected cafe count: %d, got %d", totalCount, len(list))
	}

}*/

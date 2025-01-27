package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)

	require.NoError(t, err)

	test1 := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(test1, req)

	assert.Equal(t, http.StatusOK, test1.Code, "Статус код 200")

	assert.NotEmpty(t, test1.Body.String(), "Тело ответа не пустое")
}

func TestMainHandlerCity(t *testing.T) {

	req, err := http.NewRequest("GET", "/cafe?city=spb&count=2", nil)

	require.NoError(t, err)

	test2 := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(test2, req)

	assert.Equal(t, http.StatusBadRequest, test2.Code, "Статус-код 400")

	assert.Equal(t, "wrong city value", test2.Body.String(), "Сообщение об ошибке в теле ответа")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=10", nil)

	require.NoError(t, err)

	test3 := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(test3, req)

	assert.Equal(t, http.StatusOK, test3.Code, "Cтатус-код 200")

	body := test3.Body.String()

	assert.Contains(t, body, "Мир кофе")

	assert.Contains(t, body, "Сладкоежка")

	assert.Contains(t, body, "Кофе и завтраки")

	assert.Contains(t, body, "Сытый студент")
}

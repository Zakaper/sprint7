package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func PreRespRecorder(count int, city string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=%s", count, city), nil)
	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(respRec, req)
	return respRec
}

func TestCodeResponse(t *testing.T) {
	expexcedCode := 200

	respRec := PreRespRecorder(10, "moscow")

	if assert.Equal(t, respRec.Code, expexcedCode) && assert.NotEmpty(t, respRec.Body.String()) {
		fmt.Printf("\nКод ответа: %d\nТело ответа: %s\n\n", respRec.Code, respRec.Body.String())
	}
}

func TestCityValue(t *testing.T) {
	expexcedCode := 400
	expectedMessage := "wrong city value"

	respRec := PreRespRecorder(10, "krasnokorsk")

	if assert.Equal(t, respRec.Code, expexcedCode) && assert.Equal(t, respRec.Body.String(), expectedMessage) {
		fmt.Printf("\nКод ответа: %d\nТело ответа: %s\n\n", respRec.Code, respRec.Body.String())
	}
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	respRec := PreRespRecorder(10, "moscow")
	respArray := strings.Split(respRec.Body.String(), ",")

	if assert.Len(t, respArray, totalCount) {
		fmt.Printf("\n%s\n\n", respRec.Body.String())
	}
}

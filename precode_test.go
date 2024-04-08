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

func getCafes(req *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	responseRecorder := getCafes(httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?count=5&city=moscow"), nil))

	assert.Equal(t, (strings.Join(cafeList["moscow"][:totalCount], ",")), responseRecorder.Body.Bytes())
	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body)

}

func TestMainHandlerWhenOk(t *testing.T) {

	responseRecorder := getCafes(httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?count=2&city=moscow"), nil))
	require.NotEmpty(t, responseRecorder.Code)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, (strings.Join(cafeList["moscow"][:2], ",")), responseRecorder.Body.Bytes())
}

func TestWhenWrongCity(t *testing.T) {

	responseRecorder := getCafes(httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?count=2&city=london"), nil))
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, ("wrong city value"), responseRecorder.Body.String())
}

package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Тест, если параметр count больше, чем количество доступных кафе
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String())
}

// Тест, если запрос сформирован корректно
func TestMainHandlerWithValidRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

// Тест, если передан неподдерживаемый город
func TestMainHandlerWithUnsupportedCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=saint-petersburg", nil)
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

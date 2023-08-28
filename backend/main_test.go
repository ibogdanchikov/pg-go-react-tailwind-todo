package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// Old test, just as an example.
	router := setupRouter(nil)

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(writer, req)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "{\"message\":\"New phone, who dis?\"}", writer.Body.String())
}

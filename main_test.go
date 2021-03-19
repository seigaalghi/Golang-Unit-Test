package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestHelloWorld(test *testing.T) {
	// Create response and expected struct
	type structure struct {
		Hello string
		Array []string
	}
	// Build our expected body
	body := structure{
		Hello: "world",
		Array: []string{
			"satu",
			"duar",
		},
	}
	// Grab our router
	router := SetupRouter()
	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(test, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	var response map[string]structure
	err := json.Unmarshal([]byte(w.Body.Bytes()), &response)
	// Grab the value & whether or not it exists
	value, exists := response["result"]
	// Make some assertions on the correctness of the response.
	assert.Nil(test, err)
	assert.True(test, exists)
	assert.Equal(test, body, value)
}

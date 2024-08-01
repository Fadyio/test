package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestHelloWorld(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloworld))
	defer testServer.Close()

	testClint := testServer.Client()
	response, err := testClint.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "hello world", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

func TestHealth(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClint := testServer.Client()
	response, err := testClint.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, response.StatusCode)
}

package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldShouldSucced(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()
	testClient := testServer.Client()

	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Hello World!", string(body))

	assert.Equal(t, 200, response.StatusCode)
}
func TestHelloWorldShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()
	testClient := testServer.Client()

	body := strings.NewReader("Some Body")

	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}

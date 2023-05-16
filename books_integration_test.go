package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func runTestServer() *httptest.Server {
	return httptest.NewServer(setupServer())
}

func TestIntegrationGetBooksHandler(t *testing.T) {
	testServer := runTestServer()
	defer testServer.Close()

	resp, err := http.Get(fmt.Sprintf("%s/books", testServer.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expecte 200 got: %v", resp.StatusCode)
	}
}

package client

import (
	"net/http"
	"testing"
	"time"
)

//Client basic tests

func TestClientNew(t *testing.T) {

	fonixClient := New("This-Is-A-Test-Key")

	if fonixClient.apiKey != "This-Is-A-Test-Key" {
		t.Errorf("Expected api key to be set with initializer but got '%v'", fonixClient.apiKey)
	}

	if fonixClient.baseURL != "https://sonar.fonix.io" {
		t.Errorf("Expected default sonar URL to be correct but got '%v'", fonixClient.baseURL)
	}

	fonixClient.SetBaseURL("http://localhost:6677")

	if fonixClient.baseURL != "http://localhost:6677" {
		t.Errorf("Expected custom sonar URL to be correct but got '%v'", fonixClient.baseURL)
	}

	if fonixClient.httpClient.Timeout != CLIENT_TIMEOUT {
		t.Errorf("Expected client timeout to be set to default client timeout but got '%v'", fonixClient.httpClient.Timeout)
	}

	fonixClient.SetHttpClient(&http.Client{Timeout: 4 * time.Hour})

	if fonixClient.httpClient.Timeout != 4*time.Hour {
		t.Errorf("Expected custom client timeout to be set but got '%v'", fonixClient.httpClient.Timeout)
	}

}

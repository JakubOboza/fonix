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

func TestClientApiPathGenerationt(t *testing.T) {
	fonixClient := New("This-Is-A-Test-Key")

	apiUrl1, err := fonixClient.apiUrlPath(V2_SENDSMS)

	if err != nil {
		t.Errorf("didn't expect error but got '%v'", err)
	}

	if apiUrl1 != "https://sonar.fonix.io/v2/sendsms" {
		t.Errorf("expected correct path but got '%v'", apiUrl1)
	}

	fonixClient.SetBaseURL("https://lol.troll.test/")

	apiUrl2, err := fonixClient.apiUrlPath(V2_SENDSMS)

	if err != nil {
		t.Errorf("didn't expect error but got '%v'", err)
	}

	if apiUrl2 != "https://lol.troll.test/v2/sendsms" {
		t.Errorf("expected correct path but got '%v'", apiUrl2)
	}

}

func TestClientApiBaseAndUrlPathGenerationt(t *testing.T) {
	fonixClient := New("This-Is-A-Test-Key")

	apiUrl, err := fonixClient.apiBaseUrlAndUrlPath("http://test.sub.domain.lambdacu.be/", V2_SENDSMS)

	if err != nil {
		t.Errorf("didn't expect error but got '%v'", err)
	}

	if apiUrl != "http://test.sub.domain.lambdacu.be/v2/sendsms" {
		t.Errorf("expected correct path but got '%v'", apiUrl)
	}

}

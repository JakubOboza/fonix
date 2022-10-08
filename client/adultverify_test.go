package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//API MOCK RESPONSE TESTS

//AdultVerify

func TestAdultVerifySuccess(t *testing.T) {

	expectedRequestBody := `NETWORKRETRY=no&NUMBERS=447123456767890`

	response := `{
		"success": {
			"txguid": "7CDEB38F-4370-18FD-D7CE-329F21B99209",
			"numbers": "1"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/adultverify"

		if req.URL.String() != expectedURL {
			t.Errorf("Expected URL '%s' but got '%s'", expectedURL, req.URL.String())
		}

		if req.Method != "POST" {
			t.Errorf("Was expecting POST method but got %v", req.Method)
		}

		if req.Header.Get("X-API-KEY") != "ApiKeyThatWeUseForTest" {
			t.Errorf("Was expecting API KEY HEADER to be set correct but wasn't")
		}

		requestBody, _ := ioutil.ReadAll(req.Body)

		if string(requestBody) != expectedRequestBody {
			t.Errorf("Was expecting correct body but different \nEXP:\t'%v'\n GOT:\t'%v'", expectedRequestBody, string(requestBody))
		}

		rw.Write([]byte(response))
	}))
	defer server.Close()

	c := New("ApiKeyThatWeUseForTest")
	c.SetHttpClient(server.Client()).SetBaseURL(server.URL)

	ctx := context.Background()

	avParams := &AvParams{
		Numbers:      "447123456767890",
		NetworkRetry: "no",
	}

	result, err := c.AdultVerify(ctx, avParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.TxGuid != "7CDEB38F-4370-18FD-D7CE-329F21B99209" {
		t.Errorf("Expected status to be confirmed but got '%v'", result.TxGuid)
	}

}

func TestAdultVerifyFailure(t *testing.T) {

	expectedRequestBody := `NETWORKRETRY=no&NUMBERS=447123456767890`

	response := `{
		"failure": {
			"parameter": "numbers",
			"failcode": "IS_INVALID"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/adultverify"

		if req.URL.String() != expectedURL {
			t.Errorf("Expected URL '%s' but got '%s'", expectedURL, req.URL.String())
		}

		if req.Method != "POST" {
			t.Errorf("Was expecting POST method but got %v", req.Method)
		}

		if req.Header.Get("X-API-KEY") != "ApiKeyThatWeUseForTest" {
			t.Errorf("Was expecting API KEY HEADER to be set correct but wasn't")
		}

		requestBody, _ := ioutil.ReadAll(req.Body)

		if string(requestBody) != expectedRequestBody {
			t.Errorf("Was expecting correct body but different \nEXP:\t'%v'\n GOT:\t'%v'", expectedRequestBody, string(requestBody))
		}

		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(response))
	}))
	defer server.Close()

	c := New("ApiKeyThatWeUseForTest")
	c.SetHttpClient(server.Client()).SetBaseURL(server.URL)

	ctx := context.Background()

	avParams := &AvParams{
		Numbers:      "447123456767890",
		NetworkRetry: "no",
	}

	result, err := c.AdultVerify(ctx, avParams)

	if err.Error() != "failure! paramater: numbers error: IS_INVALID" {
		t.Errorf("Expected error but got '%v'", err)
	}

	if result != nil {
		t.Errorf("Should not have success response but got '%v'", result)
	}

}

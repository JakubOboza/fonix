package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//API MOCK RESPONSE TESTS

//SendSms

func TestChargeSmsSuccess(t *testing.T) {

	expectedRequestBody := `BODY=Hey+man+this+is+a+test&NUMBERS=447123456767890&ORIGINATOR=1234567`

	response := `{
		"success": {
			"txguid": "7CDEB38F-4370-18FD-D7CE-329F21B99209",
			"numbers": "1",
			"price": "50",
			"encoding": "gsm"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/chargesms"

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

	smsParams := &SmsParams{
		Originator: "1234567",
		Numbers:    "447123456767890",
		Body:       "Hey man this is a test",
	}

	result, err := c.ChargeSms(ctx, smsParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.SuccessData.TxGuid != "7CDEB38F-4370-18FD-D7CE-329F21B99209" {
		t.Errorf("Expected status to be confirmed but got '%v'", result.SuccessData.TxGuid)
	}

	if result.SuccessData.Price != "50" {
		t.Errorf("Expected price to be 50 but got '%v'", result.SuccessData.Price)
	}

}

func TestChargeSmsFailure(t *testing.T) {

	expectedRequestBody := `BODY=Hey+man+this+is+a+test&NUMBERS=447123456767890&ORIGINATOR=1234567`

	response := `{
		"failure": {
			"parameter": "body",
			"failcode": "TOO_MANY_CHARACTERS"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/chargesms"

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

	smsParams := &SmsParams{
		Originator: "1234567",
		Numbers:    "447123456767890",
		Body:       "Hey man this is a test",
	}

	result, err := c.ChargeSms(ctx, smsParams)

	if err.Error() != "failure! paramater: body error: TOO_MANY_CHARACTERS" {
		t.Errorf("Expected error but got '%v'", err)
	}

	if result != nil {
		t.Errorf("Should not have success response but got '%v'", result)
	}

}

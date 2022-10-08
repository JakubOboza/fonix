package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//API MOCK RESPONSE TESTS

//SendWapPush

func TestSendWapPushSuccess(t *testing.T) {

	expectedRequestBody := "NUMBERS=447123456767890&ORIGINATOR=1234567&PUSHLINK=http%3A%2F%2Fgoogle.com%2Flink%2Fto%2Fpush.wap&PUSHTITLE=amazing+Test+Promo"

	response := `{
		"success": {
			"txguid": "7CDEB38F-4370-18FD-D7CE-329F21B99209",
			"numbers": "1",
			"smsparts": "1"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/sendwappush"

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

	smsWapParams := &SmsWapParams{
		Originator: "1234567",
		Numbers:    "447123456767890",
		PushTitle:  "amazing Test Promo",
		PushLink:   "http://google.com/link/to/push.wap",
	}

	result, err := c.SendWapPush(ctx, smsWapParams)

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

func TestSendWapPushFailure(t *testing.T) {

	expectedRequestBody := `NUMBERS=447123456767890&ORIGINATOR=1234567&PUSHLINK=http%3A%2F%2Fgoogle.com%2Flink%2Fto%2Fpush.wap&PUSHTITLE=amazing+Test+Promo`

	response := `{
		"failure": {
			"parameter": "PUSHTITLE",
			"failcode": "IS_EMPTY"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/sendwappush"

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

	smsWapParams := &SmsWapParams{
		Originator: "1234567",
		Numbers:    "447123456767890",
		PushTitle:  "amazing Test Promo",
		PushLink:   "http://google.com/link/to/push.wap",
	}

	result, err := c.SendWapPush(ctx, smsWapParams)

	if err.Error() != "failure! paramater: PUSHTITLE error: IS_EMPTY" {
		t.Errorf("Expected error but got '%v'", err)
	}

	if result != nil {
		t.Errorf("Should not have success response but got '%v'", result)
	}

}

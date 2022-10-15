package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomApiError(t *testing.T) {

	expectedRequestBody := `BODY=Hey+man+this+is+a+test&NUMBERS=447123456767890&ORIGINATOR=1234567`

	response := `{
		"failure": {
			"parameter": "body",
			"failcode": "IS_EMPTY"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/sendsms"

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

	result, err := c.SendSms(ctx, smsParams)

	if err.Error() != "failure! paramater: body error: IS_EMPTY" {
		t.Errorf("Expected error but got '%v'", err)
	}

	if result != nil {
		t.Errorf("Should not have success response but got '%v'", result)
	}

	apiError, ok := err.(*ApiError)

	if ok != true {
		t.Errorf("returned error should be castable to *ApiError but wasnt")
	}

	if ok {
		if apiError.StatusCode != 400 {
			t.Errorf("expected StatusCode to be 404 but got %v", apiError.StatusCode)
		}

		if apiError.ErrorData.Failcode != "IS_EMPTY" {
			t.Errorf("expected FailCode to be 'IS_EMPTY' but got %v", apiError.ErrorData.Failcode)
		}

		if apiError.ErrorData.Parameter != "body" {
			t.Errorf("expected Parameter to be 'body' but got %v", apiError.ErrorData.Parameter)
		}

	}

}

func TestCustomApiErrorWrapError(t *testing.T) {

	errRes := errorResponse{ErrorData: errorResponseContent{
		Parameter: "body",
		Failcode:  "IS_EMPTY",
	}}

	apiError1 := wrapApiError(errRes, nil)

	if apiError1.StatusCode != 0 {
		t.Errorf("expected StatusCode to be 0 without response but got %v", apiError1.StatusCode)
	}

	res := &http.Response{
		StatusCode: http.StatusBadRequest,
	}

	apiError2 := wrapApiError(errRes, res)

	if apiError2.StatusCode != 400 {
		t.Errorf("expected StatusCode to be 400 without response but got %v", apiError2.StatusCode)
	}

}

func TestClientWithAllErrorFields(t *testing.T) {
	response := `{
		"failure":{
			"ifversion":"201001",
			"statuscode":"WINDOW_EXCEEDED",
			"statustext":"Too many requests made to the server in parallel"
		}
	}`

	expectedRequestBody := `NUMBERS=447123456767890&ORIGINATOR=1234567&PUSHLINK=http%3A%2F%2Fgoogle.com%2Flink%2Fto%2Fpush.wap&PUSHTITLE=amazing+Test+Promo`

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

	if result != nil {
		t.Errorf("Should not have success response but got '%v'", result)
	}

	apiError, ok := err.(*ApiError)

	if ok != true {
		t.Errorf("returned error should be castable to *ApiError but wasnt")
	}

	if ok {
		if apiError.ErrorData.StatusCode != "WINDOW_EXCEEDED" {
			t.Errorf("expected StatusCode to be WINDOW_EXCEEDED but got %v", apiError.ErrorData.StatusCode)
		}

		if apiError.ErrorData.StatusText != "Too many requests made to the server in parallel" {
			t.Errorf("expected FailCode to be 'Too many requests made to the server in parallel' but got '%v'", apiError.ErrorData.StatusText)
		}

	}

}

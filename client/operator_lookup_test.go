package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//API MOCK RESPONSE TESTS

//OperatorLookup

func TestOperatorLookupSuccess(t *testing.T) {

	expectedRequestBody := ``

	response := `{
		"success": {
			"mcc":"234",
			"mnc":"15",
			"operator":"voda-uk"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/operator_lookup?NUMBER=447123456767890"

		if req.URL.String() != expectedURL {
			t.Errorf("Expected URL '%s' but got '%s'", expectedURL, req.URL.String())
		}

		if req.Method != "GET" {
			t.Errorf("Was expecting GET method but got %v", req.Method)
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

	opParams := &OperatorLookupParams{
		Number: "447123456767890",
	}

	result, err := c.OperatorLookup(ctx, opParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.Mcc != "234" {
		t.Errorf("Expected mcc to be 234 but got '%v'", result.Mcc)
	}

	if result.Mnc != "15" {
		t.Errorf("Expected mnc to be 15 but got '%v'", result.Mnc)
	}

	if result.Operator != "voda-uk" {
		t.Errorf("Expected operator to be 'voda-uk' but got '%v'", result.Operator)
	}

}

func TestOperatorLookupFailure(t *testing.T) {

	response := `{
		"failure": {
			"parameter": "number",
			"failcode": "IS_INVALID"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/operator_lookup?NUMBER=447123456767890"

		if req.URL.String() != expectedURL {
			t.Errorf("Expected URL '%s' but got '%s'", expectedURL, req.URL.String())
		}

		if req.Method != "GET" {
			t.Errorf("Was expecting POST method but got %v", req.Method)
		}

		if req.Header.Get("X-API-KEY") != "ApiKeyThatWeUseForTest" {
			t.Errorf("Was expecting API KEY HEADER to be set correct but wasn't")
		}

		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(response))
	}))
	defer server.Close()

	c := New("ApiKeyThatWeUseForTest")
	c.SetHttpClient(server.Client()).SetBaseURL(server.URL)

	ctx := context.Background()

	opParams := &OperatorLookupParams{
		Number: "447123456767890",
	}

	result, err := c.OperatorLookup(ctx, opParams)

	if err.Error() != "failure! paramater: number error: IS_INVALID" {
		t.Errorf("Expected error but got '%v'", err)
	}

	if result != nil {
		t.Errorf("Should not have success response but got '%v'", result)
	}

}

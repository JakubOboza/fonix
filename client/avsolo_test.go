package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//API MOCK RESPONSE TESTS

//AvSolo

func TestAvSoloVerifiedSuccess(t *testing.T) {

	expectedRequestBody := `NETWORKRETRY=no&NUMBERS=447123456767890`

	response := `{
		"verified":{
			"ifversion":"201001",
			"operator":"o2-uk",
			"guid":"f2e84610-534a-4e02-bfad-9ccbaa14e09a"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/avsolo"

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
	c.SetHttpClient(server.Client()).SetBaseAvSoloURL(server.URL)

	ctx := context.Background()

	avParams := &AvParams{
		Numbers:      "447123456767890",
		NetworkRetry: "no",
	}

	result, err := c.AvSolo(ctx, avParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.Status != AV_VERIFIED {
		t.Errorf("Expected status to be 'verified' but got '%v'", result.Status)
	}

	if result.IfVersion != "201001" {
		t.Errorf("Expected ifversion to be '201001' but got '%v'", result.IfVersion)
	}

	if result.Guid != "f2e84610-534a-4e02-bfad-9ccbaa14e09a" {
		t.Errorf("Expected guid to be 'f2e84610-534a-4e02-bfad-9ccbaa14e09a' but got '%v'", result.Guid)
	}

	if result.Operator != "o2-uk" {
		t.Errorf("Expected operator to be 'o2-uk' but got '%v'", result.Operator)
	}

}

func TestAvSoloFailure(t *testing.T) {

	expectedRequestBody := `NETWORKRETRY=no&NUMBERS=447123456767890`

	response := `{
		"failure": {
			"parameter": "numbers",
			"failcode": "IS_INVALID"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/avsolo"

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
	c.SetHttpClient(server.Client()).SetBaseAvSoloURL(server.URL)

	ctx := context.Background()

	avParams := &AvParams{
		Numbers:      "447123456767890",
		NetworkRetry: "no",
	}

	result, err := c.AvSolo(ctx, avParams)

	if err.Error() != "failure! paramater: numbers error: IS_INVALID" {
		t.Errorf("Expected error but got '%v'", err)
	}

	if result != nil {
		t.Errorf("Should not have success response but got '%v'", result)
	}

}

// no_verified
func TestAvSoloNotVerifiedSuccess(t *testing.T) {

	expectedRequestBody := `NETWORKRETRY=no&NUMBERS=447123456767890`

	response := `{
		"not_verified":{
			"ifversion":"201001",
			"operator":"o2-uk",
			"guid":"f2e84610-534a-4e02-bfad-9ccbaa14e09a"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/avsolo"

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
	c.SetHttpClient(server.Client()).SetBaseAvSoloURL(server.URL)

	ctx := context.Background()

	avParams := &AvParams{
		Numbers:      "447123456767890",
		NetworkRetry: "no",
	}

	result, err := c.AvSolo(ctx, avParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.Status != AV_NOT_VERIFIED {
		t.Errorf("Expected status to be 'not_verified' but got '%v'", result.Status)
	}

	if result.IfVersion != "201001" {
		t.Errorf("Expected ifversion to be '201001' but got '%v'", result.IfVersion)
	}

	if result.Guid != "f2e84610-534a-4e02-bfad-9ccbaa14e09a" {
		t.Errorf("Expected guid to be 'f2e84610-534a-4e02-bfad-9ccbaa14e09a' but got '%v'", result.Guid)
	}

	if result.Operator != "o2-uk" {
		t.Errorf("Expected operator to be 'o2-uk' but got '%v'", result.Operator)
	}

}

// pending

func TestAvSoloPendingSuccess(t *testing.T) {

	expectedRequestBody := `NETWORKRETRY=no&NUMBERS=447123456767890`

	response := `{
		"pending":{
			"ifversion":"201001",
			"operator":"o2-uk",
			"guid":"f2e84610-534a-4e02-bfad-9ccbaa14e09a"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/avsolo"

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
	c.SetHttpClient(server.Client()).SetBaseAvSoloURL(server.URL)

	ctx := context.Background()

	avParams := &AvParams{
		Numbers:      "447123456767890",
		NetworkRetry: "no",
	}

	result, err := c.AvSolo(ctx, avParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.Status != AV_PENDING {
		t.Errorf("Expected status to be 'pending' but got '%v'", result.Status)
	}

	if result.IfVersion != "201001" {
		t.Errorf("Expected ifversion to be '201001' but got '%v'", result.IfVersion)
	}

	if result.Guid != "f2e84610-534a-4e02-bfad-9ccbaa14e09a" {
		t.Errorf("Expected guid to be 'f2e84610-534a-4e02-bfad-9ccbaa14e09a' but got '%v'", result.Guid)
	}

	if result.Operator != "o2-uk" {
		t.Errorf("Expected operator to be 'o2-uk' but got '%v'", result.Operator)
	}

}

// unknown

func TestAvSoloUnknownSuccess(t *testing.T) {

	expectedRequestBody := `NETWORKRETRY=no&NUMBERS=447123456767890`

	response := `{
		"unknown":{
			"ifversion":"201001",
			"operator":"o2-uk",
			"guid":"f2e84610-534a-4e02-bfad-9ccbaa14e09a"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/avsolo"

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
	c.SetHttpClient(server.Client()).SetBaseAvSoloURL(server.URL)

	ctx := context.Background()

	avParams := &AvParams{
		Numbers:      "447123456767890",
		NetworkRetry: "no",
	}

	result, err := c.AvSolo(ctx, avParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.Status != AV_UNKNOWN {
		t.Errorf("Expected status to be 'unknown' but got '%v'", result.Status)
	}

	if result.IfVersion != "201001" {
		t.Errorf("Expected ifversion to be '201001' but got '%v'", result.IfVersion)
	}

	if result.Guid != "f2e84610-534a-4e02-bfad-9ccbaa14e09a" {
		t.Errorf("Expected guid to be 'f2e84610-534a-4e02-bfad-9ccbaa14e09a' but got '%v'", result.Guid)
	}

	if result.Operator != "o2-uk" {
		t.Errorf("Expected operator to be 'o2-uk' but got '%v'", result.Operator)
	}

}

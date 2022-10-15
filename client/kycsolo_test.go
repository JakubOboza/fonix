package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//API MOCK RESPONSE TESTS

//KycSolo

func TestKycSoloSuccess(t *testing.T) {

	expectedRequestBody := "DOB=13-01-1900&HOUSE_NUMBER=143&NAME=test&NUMBER=447123456767890&POSTCODE=N13+ECR&REQUESTID=request01&SURNAME=surtest"

	response := `{
			"completed":{
				"ifversion":"201001",
				"statuscode":"OK",
				"statustext":"Successfully checked",
				"guid":"r-1-requestid",
				"requestid":"requestid",
				"status_time":"202106081801213",
				"first_name_match":true,
				"last_name_match":false,
				"full_name_match":false,
				"postcode_match":true,
				"house_match":"unknown",
				"full_address_match":"unknown",
				"birthday_match":true,
				"is_stolen":false,
				"contract_type":"PAYM"
			}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/kyc"

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
	c.SetHttpClient(server.Client()).SetBaseKycSoloURL(server.URL)

	ctx := context.Background()

	kycParams := &KycSoloParams{
		Name:        "test",
		Surname:     "surtest",
		Number:      "447123456767890",
		HouseNumber: "143",
		PostCode:    "N13 ECR",
		DOB:         "13-01-1900",
		RequestID:   "request01",
	}

	result, err := c.KycSolo(ctx, kycParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.CompletedData == nil {
		t.Errorf("Should have success response  data but got nil, dind't deserialize?")
	}

	if result.CompletedData.Guid != "r-1-requestid" {
		t.Errorf("Expected Guid to be match but got '%v'", result.CompletedData.Guid)
	}

	if result.CompletedData.IfVersion != "201001" {
		t.Errorf("Expected IfVersion to be match but got '%v'", result.CompletedData.IfVersion)
	}

	if result.CompletedData.StatusCode != "OK" {
		t.Errorf("Expected StatusCode to be match but got '%v'", result.CompletedData.StatusCode)
	}

	if result.CompletedData.StatusText != "Successfully checked" {
		t.Errorf("Expected StatusText to be match but got '%v'", result.CompletedData.StatusText)
	}

	if result.CompletedData.StatusTime != "202106081801213" {
		t.Errorf("Expected StatusTime to be match but got '%v'", result.CompletedData.StatusTime)
	}

	if result.CompletedData.ContractType != "PAYM" {
		t.Errorf("Expected ContractType to be match but got '%v'", result.CompletedData.ContractType)
	}

	val, ok := result.CompletedData.LastNameMatch.(bool)

	if ok != true {
		t.Errorf("Expected last name match to be bool")
	}

	if val != false {
		t.Errorf("Expected LastNameMatch in this case to be false but got '%v'", val)
	}

	unknownVal, ok := result.CompletedData.FullAddressMatch.(string)

	if ok != true {
		t.Errorf("Expected full address match to be string")
	}

	if unknownVal != "unknown" {
		t.Errorf("expected full address match to be unknown but got '%v'", unknownVal)
	}

	if !IsUnknownValueMatch(result.CompletedData.FullAddressMatch) {
		t.Errorf("Expected full address match to be unknown value")
	}

	_, err = GetBoolFromMaybeValue(result.CompletedData.FullAddressMatch)

	if err != ErrUnknownValue {
		t.Errorf("Expected error unknown value but got '%v'", err)
	}

	stolen, err := GetBoolFromMaybeValue(result.CompletedData.IsStolen)

	if err != nil {
		t.Errorf("Didnt expect error but got '%v'", err)
	}

	if stolen != false {
		t.Errorf("Expected not to be stolen but got '%v'", stolen)
	}

}

func TestKycSoloPending(t *testing.T) {

	expectedRequestBody := `DOB=13-01-1900&HOUSE_NUMBER=143&NAME=test&NUMBER=447123456767890&POSTCODE=N13+ECR&REQUESTID=request01&SURNAME=surtest`

	response := `{
		"pending":{
			"ifversion":"201001",
			"statuscode":"PENDING",
			"statustext":"The request is still processing",
			"guid":"r-1-test",
			"requestid":"test"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/kyc"

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
	c.SetHttpClient(server.Client()).SetBaseKycSoloURL(server.URL)

	ctx := context.Background()

	kycParams := &KycSoloParams{
		Name:        "test",
		Surname:     "surtest",
		Number:      "447123456767890",
		HouseNumber: "143",
		PostCode:    "N13 ECR",
		DOB:         "13-01-1900",
		RequestID:   "request01",
	}

	result, err := c.KycSolo(ctx, kycParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.CompletedData != nil {
		t.Errorf("Should not have success response  data but got nil, dind't deserialize?")
	}

	if result.PendingData == nil {
		t.Errorf("Should have pending response  data but got nil, dind't deserialize?")
	}

	if result.PendingData.StatusCode != "PENDING" {
		t.Errorf("Expected status code to be PENDING but got '%v'", result.PendingData.StatusCode)
	}

}

func TestKycSoloFailure(t *testing.T) {

	expectedRequestBody := `DOB=13-01-1900&HOUSE_NUMBER=143&NAME=test&NUMBER=447123456767890&POSTCODE=N13+ECR&REQUESTID=request01&SURNAME=surtest`

	response := `{
		"failure": {
			"parameter": "PUSHTITLE",
			"failcode": "IS_EMPTY"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/kyc"

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
	c.SetHttpClient(server.Client()).SetBaseKycSoloURL(server.URL)

	ctx := context.Background()

	kycParams := &KycSoloParams{
		Name:        "test",
		Surname:     "surtest",
		Number:      "447123456767890",
		HouseNumber: "143",
		PostCode:    "N13 ECR",
		DOB:         "13-01-1900",
		RequestID:   "request01",
	}

	result, err := c.KycSolo(ctx, kycParams)

	if err.Error() != "failure! paramater: PUSHTITLE error: IS_EMPTY" {
		t.Errorf("Expected error but got '%v'", err)
	}

	if result != nil {
		t.Errorf("Should not have success response but got '%v'", result)
	}

}

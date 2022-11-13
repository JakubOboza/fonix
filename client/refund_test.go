package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//API MOCK RESPONSE TESTS

//Refund

func TestRefundSuccess(t *testing.T) {

	expectedRequestBody := `CHARGE_GUID=bb4adwe-dasds-asdasd-asdasd&NUMBERS=447123456767890&REQUESTID=rid-1233433232`

	response := `{
		"success":{
			"ifversion":"201001",
			"statuscode":"OK",
			"statustext":"Successfully Refunded",
			"guid":"r-1-requestid1",
			"requestid":"requestid1",
			"charge_guid":"chargeguid1",
			"refund_time":"20210803153118",
			"refunded_amount_in_pence":150
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/refund"

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
	c.SetHttpClient(server.Client()).SetBaseRefundURL(server.URL)

	ctx := context.Background()

	refundParams := &RefundParams{
		Numbers:    "447123456767890",
		RequestID:  "rid-1233433232",
		ChargeGuid: "bb4adwe-dasds-asdasd-asdasd",
	}

	result, err := c.Refund(ctx, refundParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.Status != REFUND_SUCCESS {
		t.Errorf("Expected status to be 'success' but got '%v'", result.Status)
	}

	if result.IfVersion != "201001" {
		t.Errorf("Expected ifversion to be '201001' but got '%v'", result.IfVersion)
	}

	if result.Guid != "r-1-requestid1" {
		t.Errorf("Expected guid to be 'r-1-requestid1' but got '%v'", result.Guid)
	}

	if result.ChargeGuid != "chargeguid1" {
		t.Errorf("Expected guid to be 'chargeguid1' but got '%v'", result.ChargeGuid)
	}

	if result.RefundAmountInPence != 150 {
		t.Errorf("Expected refunded amount to be '150' but got '%v'", result.RefundAmountInPence)
	}

}

func TestRefundFailure(t *testing.T) {

	expectedRequestBody := `CHARGE_GUID=bb4adwe-dasds-asdasd-asdasd&NUMBERS=447123456767890&REQUESTID=rid-1233433232`

	response := `{
		"failure": {
			"parameter": "numbers",
			"failcode": "IS_INVALID"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/refund"

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
	c.SetHttpClient(server.Client()).SetBaseRefundURL(server.URL)

	ctx := context.Background()

	refundParams := &RefundParams{
		Numbers:    "447123456767890",
		RequestID:  "rid-1233433232",
		ChargeGuid: "bb4adwe-dasds-asdasd-asdasd",
	}

	result, err := c.Refund(ctx, refundParams)

	if err.Error() != "failure! paramater: numbers error: IS_INVALID" {
		t.Errorf("Expected error but got '%v'", err)
	}

	if result != nil {
		t.Errorf("Should not have success response but got '%v'", result)
	}

}

// failure with response

func TestRefundFailureWithResponse(t *testing.T) {

	expectedRequestBody := `CHARGE_GUID=bb4adwe-dasds-asdasd-asdasd&NUMBERS=447123456767890&REQUESTID=rid-1233433232`

	response := `{
		"failure":{
			"ifversion":"201001",
			"statuscode":"MNO_TX_NOT_FOUND",
			"statustext":"Charge Transaction Not Found",
			"guid":"r-1-requestid1",
			"requestid":"requestid1",
			"charge_guid":"chargeguid1",
			"refund_time":"20210714155121"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/refund"

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
	c.SetHttpClient(server.Client()).SetBaseRefundURL(server.URL)

	ctx := context.Background()

	refundParams := &RefundParams{
		Numbers:    "447123456767890",
		RequestID:  "rid-1233433232",
		ChargeGuid: "bb4adwe-dasds-asdasd-asdasd",
	}

	result, err := c.Refund(ctx, refundParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.Status != REFUND_FAILURE {
		t.Errorf("Expected status to be 'success' but got '%v'", result.Status)
	}

	if result.IfVersion != "201001" {
		t.Errorf("Expected ifversion to be '201001' but got '%v'", result.IfVersion)
	}

	if result.Guid != "r-1-requestid1" {
		t.Errorf("Expected guid to be 'r-1-requestid1' but got '%v'", result.Guid)
	}

	if result.ChargeGuid != "chargeguid1" {
		t.Errorf("Expected guid to be 'chargeguid1' but got '%v'", result.ChargeGuid)
	}

	if result.StatusCode != "MNO_TX_NOT_FOUND" {
		t.Errorf("Expected status code to be 'MNO_TX_NOT_FOUND' but got '%v'", result.StatusCode)
	}

}

// pending

func TestRefundPending(t *testing.T) {

	expectedRequestBody := `CHARGE_GUID=bb4adwe-dasds-asdasd-asdasd&NUMBERS=447123456767890&REQUESTID=rid-1233433232`

	response := `{
		"pending":{
			"ifversion":"201001",
			"statuscode":"PENDING",
			"statustext":"The request is still processing",
			"guid":"r-1-requestid1",
			"requestid":"requestid1"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/refund"

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
	c.SetHttpClient(server.Client()).SetBaseRefundURL(server.URL)

	ctx := context.Background()

	refundParams := &RefundParams{
		Numbers:    "447123456767890",
		RequestID:  "rid-1233433232",
		ChargeGuid: "bb4adwe-dasds-asdasd-asdasd",
	}

	result, err := c.Refund(ctx, refundParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.Status != REFUND_PENDING {
		t.Errorf("Expected status to be 'pending' but got '%v'", result.Status)
	}

	if result.IfVersion != "201001" {
		t.Errorf("Expected ifversion to be '201001' but got '%v'", result.IfVersion)
	}

	if result.Guid != "r-1-requestid1" {
		t.Errorf("Expected guid to be 'r-1-requestid1' but got '%v'", result.Guid)
	}

	if result.RequestID != "requestid1" {
		t.Errorf("Expected request id to be 'requestid1' but got '%v'", result.RequestID)
	}

}

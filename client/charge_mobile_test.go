package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//API MOCK RESPONSE TESTS

//ChargeMobile

func TestChargeMobileSuccess(t *testing.T) {

	expectedRequestBody := `AMOUNT=113&BODY=Thanks+for+playing+russian+roulette%2C+you+are+dead&CHARGEDESCRIPTION=Mobile+Games+Service&CHARGESILENT=no&CURRENCY=GBP&NUMBERS=eetmo-uk.440200000017&REQUESTID=F21B992E9257936E3D2F7CDEB38F217C&TIMETOLIVE=10`

	response := `{
		"success": {
			"txguid": "r-4-F21B992E9257936E3D2F7CDEB38F217C",
			"numbers": "1",
			"encoding": "gsm"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/chargemobile"

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

	chargeMobileParams := &ChargeMobileParams{
		Numbers:           "eetmo-uk.440200000017",
		Amount:            113,
		Currency:          "GBP",
		RequestID:         "F21B992E9257936E3D2F7CDEB38F217C",
		ChargeDescription: "Mobile Games Service",
		TimeToLive:        10,
		ChargeSilent:      "no",
		Body:              "Thanks for playing russian roulette, you are dead",
	}

	result, err := c.ChargeMobile(ctx, chargeMobileParams)

	if err != nil {
		t.Errorf("Didn't expect error but got '%v'", err)
	}

	if result == nil {
		t.Errorf("Should have success response but got nil")
	}

	if result.TxGuid != "r-4-F21B992E9257936E3D2F7CDEB38F217C" {
		t.Errorf("Expected TxGuid to be 'r-4-F21B992E9257936E3D2F7CDEB38F217C' but got '%v'", result.TxGuid)
	}

	if result.Numbers != "1" {
		t.Errorf("Expected numbers to be '1' but got '%v'", result.Numbers)
	}

	if result.Encoding != "gsm" {
		t.Errorf("Expected encoding to be 'gsm' but got '%v'", result.Encoding)
	}

}

func TesChargeMobileFailure(t *testing.T) {

	expectedRequestBody := `AMOUNT=113&BODY=Thanks+for+playing+russian+roulette%2C+you+are+dead&CHARGEDESCRIPTION=Mobile+Games+Service&CHARGESILENT=yes&CURRENCY=GBP&NUMBERS=eetmo-uk.440200000017&REQUESTID=F21B992E9257936E3D2F7CDEB38F217C&TIMETOLIVE=10`

	response := `{
		"failure": {
			"parameter": "numbers",
			"failcode": "IS_INVALID"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedURL := "/v2/chargemobile"

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

	chargeMobileParams := &ChargeMobileParams{
		Numbers:           "eetmo-uk.440200000017",
		Amount:            113,
		Currency:          "GBP",
		RequestID:         "F21B992E9257936E3D2F7CDEB38F217C",
		ChargeDescription: "Mobile Games Service",
		TimeToLive:        10,
		ChargeSilent:      "yes",
		Body:              "Thanks for playing russian roulette, you are dead",
	}

	result, err := c.ChargeMobile(ctx, chargeMobileParams)

	if err.Error() != "failure! paramater: numbers error: IS_INVALID" {
		t.Errorf("Expected error but got '%v'", err)
	}

	if result != nil {
		t.Errorf("Should not have success response but got '%v'", result)
	}

}

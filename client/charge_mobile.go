package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Charge Mobile API call

// -d NUMBERS=eetmo-uk.440200000017
// -d AMOUNT=100
// -d CURRENCY=GBP
// -d REQUESTID=F21B992E9257936E3D2F7CDEB38F217C
// -d CHARGEDESCRIPTION=Mobile%20Games%20Service
// -d TIMETOLIVE=10
// -d CHARGESILENT=no
// -d BODY=Thanks%20for%20playing%20lucky%20roulette.%20This%20message%20is%20charged%20%C2%A31.
// -d DUMMY=yes

type ChargeMobileParams struct {
	Numbers           string
	Amount            int
	Currency          string
	RequestID         string
	ChargeDescription string
	TimeToLive        int
	ChargeSilent      string //default to "no"
	Body              string
	Dummy             string
}

func (chargeParams *ChargeMobileParams) ToParams() string {
	data := url.Values{}
	data.Set("NUMBERS", chargeParams.Numbers)
	data.Set("AMOUNT", fmt.Sprintf("%d", chargeParams.Amount))
	data.Set("CURRENCY", chargeParams.Currency)
	data.Set("REQUESTID", chargeParams.RequestID)
	data.Set("CHARGEDESCRIPTION", chargeParams.ChargeDescription)
	data.Set("TIMETOLIVE", fmt.Sprintf("%d", chargeParams.TimeToLive))
	if strings.ToUpper(chargeParams.ChargeSilent) != "YES" {
		data.Set("CHARGESILENT", "no") //default to no for empty strings etc
	} else {
		data.Set("CHARGESILENT", "yes")
	}
	data.Set("BODY", chargeParams.Body)
	if strings.ToUpper(chargeParams.Dummy) == "YES" {
		data.Set("DUMMY", "yes")
	}
	return data.Encode()
}

// "txguid": "r-4-F21B992E9257936E3D2F7CDEB38F217C",
// "numbers": "1",
// "encoding": "gsm"

type ChargeMobileResponse struct {
	TxGuid   string `json:"txguid"`
	Numbers  string `json:"numbers"`
	Encoding string `json:"encoding"`
}

type ChargeMobileResponseWrapper struct {
	Success *ChargeMobileResponse `json:"success"`
}

func (sr *ChargeMobileResponse) ToConsole() string {
	return fmt.Sprintf("======Charge Mobile Request Result======\nTxGuid: %s\nNumbers: %s\nEncoding: %s\n", sr.TxGuid, sr.Numbers, sr.Encoding)
}

func (client *Client) ChargeMobile(ctx context.Context, chargeParams *ChargeMobileParams) (*ChargeMobileResponse, error) {

	apiUrl, err := client.apiBaseUrlAndUrlPath(client.baseURL, V2_CHARGEMOBILE)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", apiUrl, strings.NewReader(chargeParams.ToParams()))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := ChargeMobileResponseWrapper{}

	if err = client.sendRequest(req, &response); err != nil {
		return nil, err
	}

	if response.Success != nil {
		return response.Success, nil
	}

	return nil, errors.New("empty response or unknown response")
}

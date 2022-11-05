package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Async AV check, You need to implement handler to get response!
// https://www.fonix.com/documentation/messaging-sms-billing-api/
// Request will look like this
// POST /path/chargereport HTTP/1.1
// IFVERSION=201001&OPERATOR=eetmo-uk&MONUMBER=447932111222&GUID=7CDEB38F-4370-18FD-D7CE-329F21B99209&AV=VERIFIED

type AvParams struct {
	NetworkRetry string
	Numbers      string
	Dummy        string
}

// -d NUMBERS=447111222333
// -d NETWORKRETRY=no
// -d DUMMY=yes

func (avParams *AvParams) ToParams() string {
	data := url.Values{}
	data.Set("NUMBERS", avParams.Numbers)
	data.Set("NETWORKRETRY", avParams.NetworkRetry)
	if strings.ToUpper(avParams.NetworkRetry) == "" {
		//default to no network retry
		data.Set("NETWORKRETRY", "NO")
	}
	if strings.ToUpper(avParams.Dummy) == "YES" {
		data.Set("DUMMY", "yes")
	}
	return data.Encode()
}

type SuccessAvResponse struct {
	TxGuid  string `json:"txguid"`
	Numbers string `json:"numbers"`
}

func (sr *SuccessAvResponse) ToConsole() string {
	return fmt.Sprintf("======Response======\nGuid: %s\nNumbers: %s\n", sr.TxGuid, sr.Numbers)
}

type SuccessAvResponseWrapper struct {
	SuccessData SuccessAvResponse `json:"success"`
}

func (client *Client) AdultVerify(ctx context.Context, avParams *AvParams) (*SuccessAvResponse, error) {

	apiUrl, err := client.apiUrlPath(V2_ADULTVERIFY)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", apiUrl, strings.NewReader(avParams.ToParams()))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := SuccessAvResponseWrapper{}

	if err = client.sendRequest(req, &response); err != nil {
		return nil, err
	}

	return &response.SuccessData, nil
}

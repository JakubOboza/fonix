package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

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

type SuccessAvResponseWrapper struct {
	SuccessData SuccessAvResponse `json:"success"`
}

func (client *Client) AdultVerify(ctx context.Context, avParams *AvParams) (*SuccessAvResponse, error) {

	apiUrl, err := client.apiUrlPath(V2_ADULTVERIFY)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(avParams.ToParams()))

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

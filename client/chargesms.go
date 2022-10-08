package client

import (
	"context"
	"net/http"
	"strings"
)

// -d ORIGINATOR=84988
// -d NUMBERS=447111222333
// -d BODY=Welcome%20Home
// -d DUMMY=yes

type SuccessChargeResponse struct {
	TxGuid   string `json:"txguid"`
	Numbers  string `json:"numbers"`
	Price    string `json:"price"`
	Encoding string `json:"encoding"`
}

type SuccessChargeResponseWrapper struct {
	SuccessData SuccessChargeResponse `json:"success"`
}

func (client *Client) ChargeSms(ctx context.Context, smsParams *SmsParams) (*SuccessChargeResponse, error) {

	apiUrl, err := client.apiUrlPath(V2_CHARGESMS)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(smsParams.ToParams()))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := SuccessChargeResponseWrapper{}

	if err = client.sendRequest(req, &response); err != nil {
		return nil, err
	}

	return &response.SuccessData, nil
}

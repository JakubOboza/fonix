package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type SmsBinParams struct {
	Originator string
	Numbers    string
	BinBody    string
	Dummy      string
}

// -d ORIGINATOR=84988
// -d NUMBERS=447111222333
// -d BINBODY=....
// -d DUMMY=yes

func (smsParams *SmsBinParams) ToParams() string {
	data := url.Values{}
	data.Set("ORIGINATOR", smsParams.Originator)
	data.Set("NUMBERS", smsParams.Numbers)
	data.Set("BINBODY", smsParams.BinBody)
	if strings.ToUpper(smsParams.Dummy) == "YES" {
		data.Set("DUMMY", "yes")
	}
	return data.Encode()
}

func (client *Client) SendBinSms(ctx context.Context, smsParams *SmsBinParams) (*SuccessResponse, error) {

	apiUrl, err := client.apiUrlPath(V2_SENDSMSBIN)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(smsParams.ToParams()))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := SuccessResponseWrapper{}

	if err = client.sendRequest(req, &response); err != nil {
		return nil, err
	}

	return &response.SuccessData, nil
}

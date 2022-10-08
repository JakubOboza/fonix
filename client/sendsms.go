package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type SmsParams struct {
	Originator string
	Numbers    string
	Body       string
	Dummy      string
}

// -d ORIGINATOR=84988
// -d NUMBERS=447111222333
// -d BODY=Welcome%20Home
// -d DUMMY=yes

func (smsParams *SmsParams) ToParams() string {
	data := url.Values{}
	data.Set("ORIGINATOR", smsParams.Originator)
	data.Set("NUMBERS", smsParams.Numbers)
	data.Set("BODY", smsParams.Body)
	if strings.ToUpper(smsParams.Dummy) == "YES" {
		data.Set("DUMMY", "yes")
	}
	return data.Encode()
}

func (client *Client) SendSms(ctx context.Context, smsParams *SmsParams) (*SuccessResponse, error) {

	apiUrl, err := client.apiUrlPath(V2_SENDSMS)

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

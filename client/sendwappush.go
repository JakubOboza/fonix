package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// -d ORIGINATOR=84988
// -d NUMBERS=447111222333,447111222444
// -d PUSHTITLE=Welcome%20Home
// -d PUSHLINK=http://www.google.com
// -d DUMMY=no

type SmsWapParams struct {
	Originator string
	Numbers    string
	PushTitle  string
	PushLink   string
	Dummy      string
}

func (smsParams *SmsWapParams) ToParams() string {
	data := url.Values{}
	data.Set("ORIGINATOR", smsParams.Originator)
	data.Set("NUMBERS", smsParams.Numbers)
	data.Set("PUSHTITLE", smsParams.PushTitle)
	data.Set("PUSHLINK", smsParams.PushLink)
	if strings.ToUpper(smsParams.Dummy) == "YES" {
		data.Set("DUMMY", "yes")
	}
	return data.Encode()
}

type SuccessWapResponse struct {
	TxGuid   string `json:"txguid"`
	Numbers  string `json:"numbers"`
	SmsParts string `json:"smsparts"`
}

func (sr *SuccessWapResponse) ToConsole() string {
	return fmt.Sprintf("======Success======\nGuid: %s\nNumbers: %s\nParts: %s\n", sr.TxGuid, sr.Numbers, sr.SmsParts)
}

type SuccessWapResponseWrapper struct {
	SuccessData SuccessWapResponse `json:"success"`
}

func (client *Client) SendWapPush(ctx context.Context, smsParams *SmsWapParams) (*SuccessWapResponse, error) {

	apiUrl, err := client.apiUrlPath(V2_SENDWAPPUSH)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(smsParams.ToParams()))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := SuccessWapResponseWrapper{}

	if err = client.sendRequest(req, &response); err != nil {
		return nil, err
	}

	return &response.SuccessData, nil
}

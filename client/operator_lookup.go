package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OperatorLookupParams struct {
	Number string
	Dummy  string
}

// -d NUMBER=447111222333
// -d DUMMY=yes

func (oplkParams *OperatorLookupParams) ToParams() string {
	data := url.Values{}
	data.Set("NUMBER", oplkParams.Number)
	if strings.ToUpper(oplkParams.Dummy) == "YES" {
		data.Set("DUMMY", "yes")
	}
	return data.Encode()
}

type SuccessOperatorLookupResponse struct {
	Mnc      string `json:"mnc"`
	Mcc      string `json:"mcc"`
	Operator string `json:"operator"`
}

func (sr *SuccessOperatorLookupResponse) ToConsole() string {
	return fmt.Sprintf("======Operator Lookup Result======\nOperator: %s\nMnc: %s\nMcc: %s\n", sr.Operator, sr.Mnc, sr.Mcc)
}

type SuccessOperatorLookupResponseWrapper struct {
	SuccessData SuccessOperatorLookupResponse `json:"success"`
}

func (client *Client) OperatorLookup(ctx context.Context, oplkParams *OperatorLookupParams) (*SuccessOperatorLookupResponse, error) {

	apiUrl, err := client.apiUrlPath(V2_OPERATORLOOKUP)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", apiUrl, nil)

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = oplkParams.ToParams()

	req = req.WithContext(ctx)

	response := SuccessOperatorLookupResponseWrapper{}

	if err = client.sendRequest(req, &response); err != nil {
		return nil, err
	}

	return &response.SuccessData, nil
}

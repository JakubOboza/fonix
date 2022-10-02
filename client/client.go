package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	DEFAULT_URL = "https://sonar.fonix.io"
	V2_SENDSMS  = "v2/sendsms"
)

var (
	//default timeout for requests is 15 seconds, can be changed by using custom http client
	// or you can change it via this VAr before creation of client.
	CLIENT_TIMEOUT = 15 * time.Second
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

func New(apiKey string) *Client {
	httpClient := &http.Client{
		Timeout: CLIENT_TIMEOUT,
	}
	return &Client{apiKey: apiKey, baseURL: DEFAULT_URL, httpClient: httpClient}
}

// Configuration
func (c *Client) SetBaseURL(baseURL string) *Client {
	c.baseURL = baseURL
	return c
}

func (c *Client) SetHttpClient(httpClient *http.Client) *Client {
	c.httpClient = httpClient
	return c
}

// API Implementation

type errorResponseContent struct {
	Parameter string `json:"parameter"`
	Failcode  string `json:"failcode"`
}

type errorResponse struct {
	ErrorData errorResponseContent `json:"failure"`
}

func (errRes *errorResponse) ToError() error {
	return fmt.Errorf("failure! paramater: %s error: %s", errRes.ErrorData.Parameter, errRes.ErrorData.Failcode)
}

type SuccessResponseContent struct {
	TxGuid   string `json:"txguid"`
	Numbers  string `json:"numbers"`
	SmsParts string `json:"smsparts"`
	Encoding string `json:"encoding"`
}

type SuccessResponse struct {
	SuccessData SuccessResponseContent `json:"success"`
}

func (client *Client) sendRequest(req *http.Request, response interface{}) error {

	// Set API key for request
	req.Header.Add("X-API-KEY", client.apiKey)

	res, err := client.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errRes.ToError() //TODO needs better formatting or custom error
		}
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	//No content, no need to unmarshal
	if res.StatusCode == http.StatusNoContent {
		return nil
	}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}

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

	url := fmt.Sprintf("%s/%s", client.baseURL, V2_SENDSMS)

	req, err := http.NewRequest("POST", url, strings.NewReader(smsParams.ToParams()))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := SuccessResponse{}

	if err = client.sendRequest(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

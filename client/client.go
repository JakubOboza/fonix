package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	DEFAULT_URL        = "https://sonar.fonix.io"
	DEFAULT_URL_AVSOLO = "https://avsolo.fonix.io"
	V2_SENDSMS         = "v2/sendsms"
	V2_SENDSMSBIN      = "v2/sendbinsms"
	V2_CHARGESMS       = "v2/chargesms"
	V2_SENDWAPPUSH     = "v2/sendwappush"
	V2_ADULTVERIFY     = "v2/adultverify"
	V2_OPERATORLOOKUP  = "v2/operator_lookup"
	V2_AVSOLO          = "v2/avsolo"
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

// Generic Types

type ApiParams interface {
	ToParams() string
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

type SuccessResponse struct {
	TxGuid   string `json:"txguid"`
	Numbers  string `json:"numbers"`
	SmsParts string `json:"smsparts"`
	Encoding string `json:"encoding"`
}

type SuccessResponseWrapper struct {
	SuccessData SuccessResponse `json:"success"`
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
			return wrapApiError(errRes, res)
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

func (client *Client) apiUrlPath(endPointPath string) (string, error) {
	u, err := url.Parse(client.baseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, endPointPath)
	apiUrl := u.String()
	return apiUrl, nil
}

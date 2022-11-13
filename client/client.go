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
	DEFAULT_URL         = "https://sonar.fonix.io"
	DEFAULT_URL_AVSOLO  = "https://avsolo.fonix.io"
	DEFAULT_URL_KYCSOLO = "https://kycsolo.fonix.io"
	DEFAULT_URL_REFUND  = "https://refund.fonix.io"
	V2_SENDSMS          = "v2/sendsms"
	V2_SENDSMSBIN       = "v2/sendbinsms"
	V2_CHARGESMS        = "v2/chargesms"
	V2_SENDWAPPUSH      = "v2/sendwappush"
	V2_ADULTVERIFY      = "v2/adultverify"
	V2_OPERATORLOOKUP   = "v2/operator_lookup"
	V2_AVSOLO           = "v2/avsolo"
	V2_KYCSOLO          = "v2/kyc"
	V2_REFUND           = "v2/refund"
	V2_CHARGEMOBILE     = "v2/chargemobile"
)

var (
	//default timeout for requests is 15 seconds, can be changed by using custom http client
	// or you can change it via this VAr before creation of client.
	CLIENT_TIMEOUT = 15 * time.Second
)

type Client struct {
	apiKey         string
	baseURL        string
	baseURLAvSolo  string
	baseURLKycSolo string
	baseURLRefund  string
	httpClient     *http.Client
}

func New(apiKey string) *Client {
	httpClient := &http.Client{
		Timeout: CLIENT_TIMEOUT,
	}
	return &Client{apiKey: apiKey, baseURL: DEFAULT_URL, httpClient: httpClient, baseURLAvSolo: DEFAULT_URL_AVSOLO, baseURLKycSolo: DEFAULT_URL_KYCSOLO, baseURLRefund: DEFAULT_URL_REFUND}
}

// Configuration
func (c *Client) SetBaseURL(baseURL string) *Client {
	c.baseURL = baseURL
	return c
}

func (c *Client) SetBaseAvSoloURL(baseAvSoloURL string) *Client {
	c.baseURLAvSolo = baseAvSoloURL
	return c
}

func (c *Client) SetBaseKycSoloURL(baseKycSoloURL string) *Client {
	c.baseURLKycSolo = baseKycSoloURL
	return c
}

func (c *Client) SetBaseRefundURL(baseRefundURL string) *Client {
	c.baseURLRefund = baseRefundURL
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
	Parameter  string `json:"parameter"`
	Failcode   string `json:"failcode"`
	StatusCode string `json:"statuscode"`
	StatusText string `json:"statustext"`
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

func (sr *SuccessResponse) ToConsole() string {
	return fmt.Sprintf("======Success======\nGuid: %s\nNumbers: %s\nParts: %s\nEncoding: %s\n", sr.TxGuid, sr.Numbers, sr.SmsParts, sr.Encoding)
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
	return client.apiBaseUrlAndUrlPath(client.baseURL, endPointPath)
}

func (client *Client) apiBaseUrlAndUrlPath(baseURL, endPointPath string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, endPointPath)
	apiUrl := u.String()
	return apiUrl, nil
}

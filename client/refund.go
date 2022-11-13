package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Refund API call

const (
	REFUND_SUCCESS = "success"
	REFUND_FAILURE = "failure"
	REFUND_PENDING = "pending"
)

// -d REQUESTID=requestid1
// -d NUMBERS=o2-uk.447400000000
// -d CHARGE_GUID=chargeguid1
// -d DUMMY=YES

type RefundParams struct {
	RequestID  string
	Numbers    string
	ChargeGuid string
	Dummy      string
}

func (refundParams *RefundParams) ToParams() string {
	data := url.Values{}
	data.Set("REQUESTID", refundParams.RequestID)
	data.Set("NUMBERS", refundParams.Numbers)
	data.Set("CHARGE_GUID", refundParams.ChargeGuid)
	if strings.ToUpper(refundParams.Dummy) == "YES" {
		data.Set("DUMMY", "yes")
	}
	return data.Encode()
}

type RefundResponse struct {
	Status              string
	IfVersion           string `json:"ifversion"`
	StatusCode          string `json:"statuscode"`
	StatusText          string `json:"statustext"`
	Guid                string `json:"guid"`
	RequestID           string `json:"requestid"`
	ChargeGuid          string `json:"charge_guid"`
	RefundTime          string `json:"refund_time"`
	RefundAmountInPence int    `json:"refunded_amount_in_pence"`
}

type RefundResponseWrapper struct {
	Success *RefundResponse `json:"success"`
	Failure *RefundResponse `json:"failure"`
	Pending *RefundResponse `json:"pending"`
}

func (sr *RefundResponse) ToConsole() string {
	return fmt.Sprintf("======Refund Request Result======\nGuid: %s\nIfVersion: %s\nStatusCode: %s\nStatusText: %s\nRequestID: %s\nChargeGuid: %s\nRefundTime: %s\nRefundAmountInPence: %d\n", sr.Guid, sr.IfVersion, sr.StatusCode, sr.StatusText, sr.RequestID, sr.ChargeGuid, sr.RefundTime, sr.RefundAmountInPence)
}

func (client *Client) Refund(ctx context.Context, refundParams *RefundParams) (*RefundResponse, error) {

	apiUrl, err := client.apiBaseUrlAndUrlPath(client.baseURLRefund, V2_REFUND)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", apiUrl, strings.NewReader(refundParams.ToParams()))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := RefundResponseWrapper{}

	if err = client.sendRequest(req, &response); err != nil {
		return nil, err
	}

	if response.Success != nil {
		response.Success.Status = REFUND_SUCCESS
		return response.Success, nil
	}

	if response.Failure != nil {
		response.Failure.Status = REFUND_FAILURE
		return response.Failure, nil
	}

	if response.Pending != nil {
		response.Pending.Status = REFUND_PENDING
		return response.Pending, nil
	}

	return nil, errors.New("empty response or unknown response")
}

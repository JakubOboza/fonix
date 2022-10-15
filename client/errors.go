package client

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrUnknownValue = errors.New("unknown value")
)

type ApiError struct {
	StatusCode   int
	ErrorData    errorResponseContent
	HttpResponse *http.Response
}

func wrapApiError(errRes errorResponse, res *http.Response) *ApiError {
	httpResponseStatusCode := 0
	if res != nil {
		httpResponseStatusCode = res.StatusCode
	}
	return &ApiError{ErrorData: errRes.ErrorData, StatusCode: httpResponseStatusCode, HttpResponse: res}
}

func (ar *ApiError) Error() string {
	return fmt.Sprintf("failure! paramater: %s error: %s", ar.ErrorData.Parameter, ar.ErrorData.Failcode)
}

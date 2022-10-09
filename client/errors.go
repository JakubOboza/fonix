package client

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	StatusCode   int
	ErrorData    errorResponseContent
	HttpResponse *http.Response
}

func wrapApiError(errRes errorResponse, res *http.Response) *ApiError {
	return &ApiError{ErrorData: errRes.ErrorData, StatusCode: res.StatusCode, HttpResponse: res}
}

func (ar *ApiError) Error() string {
	return fmt.Sprintf("failure! paramater: %s error: %s", ar.ErrorData.Parameter, ar.ErrorData.Failcode)
}

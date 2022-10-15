package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// -d NUMBER=447400000000
// -d NAME=TestName
// -d SURNAME=TestSurname
// -d HOUSE_NUMBER=1
// -d POSTCODE=TESTPOSTCODE
// -d DOB=13-01-1900
// -d REQUESTID=requestid1
// -d DUMMY=YES

type KycSoloParams struct {
	Name        string
	Surname     string
	Number      string
	HouseNumber string
	PostCode    string
	DOB         string
	RequestID   string
	Dummy       string
}

func (kycParams *KycSoloParams) ToParams() string {
	data := url.Values{}
	data.Set("NAME", kycParams.Name)
	data.Set("SURNAME", kycParams.Surname)
	data.Set("NUMBER", kycParams.Number)
	data.Set("HOUSE_NUMBER", kycParams.HouseNumber)
	data.Set("POSTCODE", kycParams.PostCode)
	data.Set("DOB", kycParams.DOB)
	data.Set("REQUESTID", kycParams.RequestID)
	if strings.ToUpper(kycParams.Dummy) == "YES" {
		data.Set("DUMMY", "yes")
	}
	return data.Encode()
}

type SuccessKycSoloResponse struct {
	Guid             string      `json:"guid"`
	IfVersion        string      `json:"ifversion"`
	StatusCode       string      `json:"statuscode"`
	StatusText       string      `json:"statustext"`
	RequestID        string      `json:"requestid"`
	StatusTime       string      `json:"status_time"`
	FirstNameMatch   interface{} `json:"first_name_match"` // match fields and is_Stole are weird string or bool types
	LastNameMatch    interface{} `json:"last_name_match"`  // this means end client needs to do a cast and see
	FullNameMatch    interface{} `json:"full_name_match"`
	PostCodeMatch    interface{} `json:"postcode_match"`
	HouseMatch       interface{} `json:"house_match"`
	FullAddressMatch interface{} `json:"full_address_match"`
	BirthdayMatch    interface{} `json:"birthday_match"`
	IsStolen         interface{} `json:"is_stolen"`
	ContractType     string      `json:"contract_type"`
}

type KycSoloResponseWrapper struct {
	CompletedData *SuccessKycSoloResponse `json:"completed"`
	PendingData   *errorResponseContent   `json:"pending"`
}

func (client *Client) KycSolo(ctx context.Context, kycParams *KycSoloParams) (*KycSoloResponseWrapper, error) {

	apiUrl, err := client.apiBaseUrlAndUrlPath(client.baseURLKycSolo, V2_KYCSOLO)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(kycParams.ToParams()))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := KycSoloResponseWrapper{}

	if err = client.sendRequest(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func IsUnknownValueMatch(rawVal interface{}) bool {
	val, ok := rawVal.(string)
	if !ok {
		return false
	}
	if strings.ToLower(val) == "unknown" {
		return true
	}
	return false
}

func GetBoolFromMaybeValue(rawVal interface{}) (bool, error) {
	if IsUnknownValueMatch(rawVal) {
		return false, ErrUnknownValue
	}

	val, ok := rawVal.(bool)
	if !ok {
		return false, ErrUnknownValue
	}

	return val, nil
}

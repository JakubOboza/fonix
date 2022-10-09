package client

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Sync AV check,
// https://www.fonix.com/documentation/messaging-sms-billing-api/

// -d NUMBERS=447111222333
// -d NETWORKRETRY=no
// -d DUMMY=yes

/*
  Possible responses from docs

  {"verified":{
       "ifversion":"201001",
       "operator":"o2-uk",
       "guid":"f2e84610-534a-4e02-bfad-9ccbaa14e09a"
   }
 }

 {"not_verified":{
       "ifversion":"201001",
       "operator":"o2-uk",
       "guid":"f2e84610-534a-4e02-bfad-9ccbaa14e09a"
   }
 }

 {"unknown":{
       "ifversion":"201001",
       "operator":"o2-uk",
      "guid":"f2e84610-534a-4e02-bfad-9ccbaa14e09a"
   }
 }

 {"pending":{
       "ifversion":"201001",
       "operator":"o2-uk",
       "guid":"f2e84610-534a-4e02-bfad-9ccbaa14e09a"
   }
 }

*/

const (
	AV_VERIFIED     = "verified"
	AV_NOT_VERIFIED = "no_verified"
	AV_UNKNOWN      = "unknown"
	AV_PENDING      = "pending"
)

type SuccessAvSoloResponse struct {
	Status    string
	IfVersion string `json:"ifversion"`
	Operator  string `json:"operator"`
	Guid      string `json:"guid"`
}

type SuccessAvSoloResponseWrapper struct {
	VerifiedData    *SuccessAvSoloResponse `json:"verified"`
	NotVerifiedData *SuccessAvSoloResponse `json:"not_verified"`
	UnknownData     *SuccessAvSoloResponse `json:"unknown"`
	PendingData     *SuccessAvSoloResponse `json:"pending"`
}

func (client *Client) AvSolo(ctx context.Context, avParams *AvParams) (*SuccessAvSoloResponse, error) {

	apiUrl, err := client.apiBaseUrlAndUrlPath(client.baseURLAvSolo, V2_AVSOLO)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(avParams.ToParams()))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := SuccessAvSoloResponseWrapper{}

	if err = client.sendRequest(req, &response); err != nil {
		return nil, err
	}

	if response.VerifiedData != nil {
		response.VerifiedData.Status = AV_VERIFIED
		return response.VerifiedData, nil
	}

	if response.NotVerifiedData != nil {
		response.NotVerifiedData.Status = AV_NOT_VERIFIED
		return response.NotVerifiedData, nil
	}

	if response.PendingData != nil {
		response.PendingData.Status = AV_PENDING
		return response.PendingData, nil
	}

	if response.UnknownData != nil {
		response.UnknownData.Status = AV_UNKNOWN
		return response.UnknownData, nil
	}

	return nil, errors.New("empty response or unknown response")
}

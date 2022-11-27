# Fonix Client

This is a `golang` Fonix v2 API client for [fonix.com](fonix.com) sms gateway. 

# Table of Contents
1. [Usage (golang code)](#usagego)
2. [Usage (command line client)](#usagecli)
3. [Status](#status)
4. [Building & Compiling](#build)
5. [Testing](#test)

## Usage in Go <a name="usagego"></a>

Import the library package eg.
```
import "github.com/JakubOboza/fonix/client"
```

Send Sms using fonix client in golang:

example:
```
package main

import (
	"context"
	"fmt"
	fonix "github.com/JakubOboza/fonix/client"
)

func main() {
	client := fonix.New("MY-API_KEY")
	response, err := client.SendSms(context.Background(), &fonix.SmsParams{Originator: "889988", Numbers: "4474123456789", Body: "Hello!"})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}

```

Api Errors, if you get error from fonix api you can try casting it to `*ApiError` to get more info like `StatusCode` access to http response or the `ErrorData` object with parsed error response which might contain info about `Parameter` or `Failcode`

example:

```
response, err := client.SendSms(context.Background(), &fonix.SmsParams{Originator: "889988", Numbers: "4474123456789", Body: "Hello!"})

if err != nil {
    apiError, ok := err.(*ApiError)

    if ok {
        fmt.Println("http response status code:", apiError.StatusCode)
        fmt.Ptinln("issue with field: ",  apiError.ErrorData.Parameter)
    }
}
```

## MO/DR handling ~ integration

Example MO handler can be implemented like this:
```
func MoHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	ifversion := r.FormValue("IFVERSION")
	monumber := r.FormValue("MONUMBER")
	operator := r.FormValue("OPERATOR")
	destination := r.FormValue("DESTINATION")
	body := r.FormValue("BODY")
	receiveTime := r.FormValue("RECEIVETIME")
	guid := r.FormValue("GUID")
	price := r.FormValue("PRICE")
	requestID := r.FormValue("REQUESTID")

	fmt.Printf("IFVERSION: %s\nMONUMBER: %s\nOPERATOR: %s\nDESTINATION: %s\nBODY: %s\nRECEIVETIME: %s\nGUID: %s\nPRICE: %s\nREQUESTID: %s\n\n", ifversion, monumber, operator, destination, body, receiveTime, guid, price, requestID)

	w.WriteHeader(http.StatusOK)
}
```

Mock DR/MO handler is implemented [HERE](https://github.com/JakubOboza/fonix/blob/main/client/handlers.go)

CLI client has mock mo and dr handler in command `./bin/fonix modrmock --port=6688` that hosts handlers on paths:

```
MOs => /mock/mosms
DRs => /mock/drs
```

## Usage in command line <a name="usagecli"></a>

Library contains both command line client available via `release` and library to integrate in your code.

It can be compiled or downloaded as binary and use to interact with fonix API

To use cmd interface please download the release or build it from source `make`. This should produce `./bin/fonix` or in case of binary `fonix` executable.

To interact with client call
```
./fonix --help
```

This will show you all command you can use. 

To send bulk sms you can use the `sendsms` command. It requires API_KEY, body, originator and numbers parammeters.

example:
```
API_KEY=live_xyz ./fonix sendsms -b test -d no -n 447111222333 -o 84988
```

Api key can be set as ENVironment variable or as param to command `-k` or `--apikey=` eg. `--apikey=live:myKey123456XYZ`

Each other param can be checked by doing

```
./fonix sendsms --help
```

For dummy billing (mock api request) you use `--dummy=yes`

example:
```
API_KEY=live_xyz ./fonix sendsms -b test -n 447111222333 -o 84988 -d yes
or
API_KEY=live_xyz ./fonix sendsms --=body test --numbers=447111222333 --ooriginator=84988 --dummy=yes
```

To check cli client version just do `./fonix version`


# Compatibility Status <a name="status"></a>

Client Lib:

- [x] - v2/sendsms
- [x] - v2/adultverify
- [x] - v2/avsolo
- [x] - v2/chargesms 
- [x] - v2/sendbinsms 
- [x] - v2/sendwappush 
- [x] - v2/operator_lookup
- [x] - v2/kycsolo
- [x] - v2/chargemobile
- [x] - v2/refund


CLI Client:

- [x] - fonix sendsms
- [x] - fonix adultverify (sync/async)
- [x] - fonix chargesms 
- [x] - fonix sendbinsms 
- [x] - fonix sendwappush 
- [x] - fonix operator_lookup
- [x] - fonix chargemobile
- [x] - fonix refund

# Build <a name="build"></a>

To build development version just type `make`. 
To build release for all platforms `make release`

# Test <a name="test"></a>

To run entire suit of tests `make test`
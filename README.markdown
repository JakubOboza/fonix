# Fonix Client

This is a `golang` Fonix v2 API client for (fonix.com)[fonix.com] sms gateway. 

# Usage

WORK IN PROGRESS

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

# Status

Client Lib:

- [x] - v2/sendsms
- [x] - v2/adultverify
- [x] - v2/avsolo
- [x] - v2/chargesms 
- [x] - v2/sendbinsms 
- [x] - v2/sendwappush 
- [x] - v2/operator_lookup

CLI Client:

- [x] - fonix sendsms
- [x] - fonix adultverify (sync/async)
- [x] - fonix chargesms 
- [ ] - fonix sendbinsms 
- [ ] - fonix sendwappush 
- [x] - fonix operator_lookup

# Command line client / Release

Library contains both command line client available via `release` and library to integrate in your code.

# Build

To build development version just type `make`. 
To build release for all platforms `make release`

# Test

To run entire suit of tests `make test`
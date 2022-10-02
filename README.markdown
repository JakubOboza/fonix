# Fonix Client

This is a `golang` v2 API client for Fonix (fonix.com)[fonix.com] sms gateway. 

# Usage

WORK IN PROGRESS

example:
```
package main

import (
    "context"
    "github.com/JakubOboza/fonix"
)

func main() {
  client = fonix.New("MY-API_KEY")
  response, err := client.SendSms(context.Background(), &fonix.SmsParams{Originator: "889988", Numbers: "4474123456789", Body: "Hello!"})

  if err != nil {
    fmt.Println(err) 
    return
  }

 fmt.Println(response)
}

```

# Status

- [x] - v2/sendsms
- [ ] - v2/adultverify
- [ ] - v2/chargesms 
- [ ] - v2/sendbinsms 
- [ ] - v2/sendwappush 
- [ ] - v2/operator_lookup

# Command line client / Release

Library contains both command line client available via `release` and library to integrate in your code.

# Build

To build development version just type `make`. 
To build release for all platforms `make release`

# Test

To run entire suit of tests `make test`
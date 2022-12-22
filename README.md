# Duitku Payment Gateway Library for Go

----------

## Import
`go get github.com/mutashim/duitku-go`

## Usage

```go
package main()

import (
    "context"
    "fmt"

    "github.com/mutashim/duitku-go"
)

func main() {

	apiKey := "VERYSECRETAPIKEY"
	merchantCode := "thisIsMerchantCode"
	duitkuApi := duitku.New(apiKey, merchantCode, true, true)

	response, err := duitkuApi.GetPaymentMethod(context.Background(), 10000, "2022-01-25 16:23:08")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(response)
}

```
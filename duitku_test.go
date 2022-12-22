package duitku_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/mutashim/duitku-go"
)

func TestNew(t *testing.T) {

	apiKey := "VERYSECRETAPIKEY"
	merchantCode := "thisIsMerchantCode"
	duitkuApi := duitku.New(apiKey, merchantCode, true, true)

	response, err := duitkuApi.GetPaymentMethod(context.Background(), 10000, "2022-01-25 16:23:08")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(response)
}

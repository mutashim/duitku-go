package duitku

import (
	"crypto/sha256"
	"fmt"
)

func Sign(merchantCode string, amount int, date string, apikey string) string {
	return fmt.Sprint(sha256.Sum256([]byte(fmt.Sprint(merchantCode, amount, date, apikey))))
}

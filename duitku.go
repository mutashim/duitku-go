package duitku

import "net/http"

func New(apiKey, merchantCode string, isSandboxMode, isSanitizedMode bool) *Client {
	base := BASE_API_DEVELOPMENT
	if !isSandboxMode {
		base = BASE_API_PRODUCTION
	}
	return &Client{
		apiKey:          apiKey,
		merchantCode:    merchantCode,
		baseApi:         base,
		isSanitizedMode: isSanitizedMode,
		http:            http.DefaultClient,
	}
}

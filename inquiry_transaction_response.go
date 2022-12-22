package duitku

type InquiryTransactionResponse struct {
	MerchantCode string `json:"merchantCode"`
	Reference    string `json:"reference"`
	PaymentURL   string `json:"paymentUrl"`
	VANumber     string `json:"vaNumber"`
	Amount       int    `json:"amount"`
	QrString     string `json:"qrString"`
}

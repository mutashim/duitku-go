package duitku

type GetPaymentMethodResponse struct {
	BaseResponse
	PaymentFee []PaymentFee `json:"paymentFee"`
}

type PaymentFee struct {
	PaymentMethod string `json:"paymentMethod"`
	PaymentName   string `json:"paymentName"`
	PaymentImage  string `json:"paymentImage"`
	TotalFee      string `json:"totalFee"`
}

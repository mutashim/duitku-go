package duitku

// Read more: https://docs.duitku.com/api/id/#parameter-request-transaksi
type InquiryTransactionRequest struct {
	MerchantCode     string         `json:"merchantCode"`
	PaymentAmount    int            `json:"paymentAmount"`
	MerchantOrderId  string         `json:"merchantOrderId"`
	ProductDetails   string         `json:"productDetails"`
	Email            string         `json:"email"`
	AdditionalParam  string         `json:"additionalParam,omitempty"`
	PaymentMethod    string         `json:"paymentMethod"`
	MerchantUserInfo string         `json:"merchantUserInfo,omitempty"`
	CustomerVaName   string         `json:"customerVaName"`
	PhoneNumber      string         `json:"phoneNumber,omitempty"`
	ItemDetails      []ItemDetail   `json:"itemDetails"`
	CustomerDetail   CustomerDetail `json:"customerDetail,omitempty"`
	ReturnUrl        string         `json:"returnUrl"`
	CallbackUrl      string         `json:"callbackUrl"`
	Signature        string         `json:"signature"`

	// Read more: https://docs.duitku.com/api/id/#expiry-period
	ExpiryPeriod int         `json:"expiryPeriod,omitempty"`
	AccountLink  AccountLink `json:"accountLink,omitempty"`
}

// Read more: https://docs.duitku.com/api/id/#customer-detail
type CustomerDetail struct {
	Firstname       string  `json:"firstName,omitempty"`
	Lastname        string  `json:"lastName,omitempty"`
	Email           string  `json:"email,omitempty"`
	PhoneNumber     string  `json:"phoneNumber,omitempty"`
	BillingAddress  Address `json:"billingAddress,omitempty"`
	ShippingAddress Address `json:"shippingAddress,omitempty"`
}

// Read more: https://docs.duitku.com/api/id/#address
type Address struct {
	Firstname   string `json:"firstName,omitempty"`
	Lastname    string `json:"lastName,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	PostalCode  string `json:"postalCode,omitempty"`
	Phone       string `json:"phone,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

// Read more: https://docs.duitku.com/api/id/#item-details
type ItemDetail struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"integer"`
}

// Parameter khusus untuk metode pembayaran yang menggunakan
// OVO Account Link dan Shopee Account Link.
//
// Read more: https://docs.duitku.com/api/id/#account-link
type AccountLink struct {
	CredentialCode string `json:"credentialCode"`
	Ovo            OVO    `json:"ovo"`
	Shopee
}

// Read more: https://docs.duitku.com/api/id/#ovo-detail
type OVO struct {
	PaymentDetails []PaymentDetailObject `json:"paymentDetails"`
	PaymentType    string                `json:"paymentType"`
	Amount         int                   `json:"amount"`
}

// Read more: https://docs.duitku.com/api/id/#ovo-detail
type PaymentDetailObject struct {
	PaymentType string `json:"paymentType"`
	Amount      string `json:"amount"`
}

// Read more: https://docs.duitku.com/api/id/#shopee-detail
type Shopee struct {
	PromoIDs string `json:"promo_ids"`
	UseCoin  bool   `json:"useCoin"`
}

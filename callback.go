package duitku

// https://docs.duitku.com/api/id/#parameter-callback
type Callback struct {
	// Kode merchant, dikirimkan oleh server Duitku
	// untuk memberitahu kode proyek yang digunakan.
	MerchantCode string `form:"merchantCode"`

	// Jumlah nominal transaksi.
	Amount int `form:"amount"`

	// Nomor transaksi dari merchant.
	MerchantOrderID string `form:"merchantOrderId"`

	// Keterangan detail produk.
	ProductDetail string `form:"productDetail"`

	// Parameter tambahan yang anda kirimkan pada saat
	// awal request transaksi.
	AdditionalParam string `form:"additionalParam"`

	// Metode Pembayaran.
	PaymentCode string `form:"paymentCode"`

	// Pemberitahuan callback hasil transaksi.
	//
	// 00 - Success
	//
	// 01 - Failed
	ResultCode int `form:"resultCode"`

	// Username atau email pelanggan di situs anda.
	MerchantUserID string `form:"merchantUserId"`

	// Nomor referensi transaksi dari Duitku. Mohon
	// disimpan untuk keperluan pencatatan atau pelacakan
	//  transaksi.
	Reference string `form:"reference"`

	// Kode identifikasi transaksi. Berisikan parameter-
	// parameter transaksi yang di-hash menggunakan metode
	//  hashing MD5. Parameter keamanan sebagai acuan bahwa
	//  request yang diterima berasal dari server Duitku.
	// Formula: MD5(merchantcode + amount + merchantOrderId + apiKey)
	Signature string `form:"signature"`

	// Di kirim melalui callback jika pembayaran menggunakan
	// metode pembayaran ShopeePay(QRIS, App, dan Account Link).
	// Jika berisi string dengan kombinasi angka dan huruf,
	// maka menandakan pembayaran menggunakan Shopee itu sendiri.
	SpUserHash string `form:"spUserHash"`
}

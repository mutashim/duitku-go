package duitku

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	apiKey          string
	merchantCode    string
	isSanitizedMode bool
	baseApi         string
	http            *http.Client
}

func (c *Client) sendRequest(ctx context.Context, method, path string, data []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.baseApi, path), bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("client: could not create request: %s", err)
	}

	res, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client: error making http request: %s", err)
	}

	defer res.Body.Close()
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("client: could not read response body: %s", err)
	}
	log.Println(string(responseBody[:]))

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("client: error response: %s:%s", res.Status, res.Body)
	}

	return responseBody, nil
}

// Proses ini digunakan untuk mendapatkan metode pembayaran yang aktif
// dari proyek merchant (anda). API ini berisi nama metode pembayaran,
// biaya dan URL ke gambar metode pembayaran. Anda dapat menggunakan
// sebagai daftar channel pembayaran pada proyek anda dan anda akan
// mendapatkan paymentMethod yang berguna untuk diteruskan ke proses
// request transaksi. Proses ini opsional, anda dapat melewatinya menuju
// Permintaan Transaksi.

func (c *Client) GetPaymentMethod(ctx context.Context, amount int, datetime string) (*[]PaymentFee, error) {
	rawdata := map[string]interface{}{
		"merchantcode": c.merchantCode,
		"amount":       amount,
		"datetime":     datetime,
		"signature":    Sign(c.merchantCode, amount, datetime, c.apiKey),
	}

	data, err := json.Marshal(rawdata)
	if err != nil {
		return nil, fmt.Errorf("error parsing data into json: %w", err)
	}

	res, err := c.sendRequest(ctx, "POST", "webapi/api/merchant/paymentmethod/getpaymentmethod", data)
	if err != nil {
		return nil, err
	}

	respon := GetPaymentMethodResponse{}
	if err := json.Unmarshal(res, &respon); err != nil {
		return nil, err
	}

	return &respon.PaymentFee, nil
}

// Permintaan Transaksi
//
// Berikut ini adalah langkah utama pada proses transaksi
// diawali dengan melakukan request transaksi ke sistem Duitku.
// Proses ini akan diperuntukan untuk membuat pembayaran (melalui
// nomor virtual account, QRIS, e-wallet, dsb). Anda dapat membuat
// suatu halaman pembayaran yang berguna mengarahkan pelanggan
// membayar tagihan transaksinya kepada anda. Silahkan untuk membuat
// request transaksi dengan membuat HTTP request seperti berikut.
// Jika anda melewati Get Payment Method, anda dapat mengisi
// paymentMethod dengan referensi Metode Pembayaran.
func (c *Client) InquiryTransaction(ctx context.Context, r *InquiryTransactionRequest) (*InquiryTransactionResponse, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("error parsing data into json: %w", err)
	}

	res, err := c.sendRequest(ctx, "POST", "webapi/api/merchant/v2/inquiry", data)
	if err != nil {
		return nil, err
	}

	response := InquiryTransactionResponse{}
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

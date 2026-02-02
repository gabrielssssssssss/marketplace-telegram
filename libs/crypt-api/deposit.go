package cryptapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func (c *CryptAPI) NewPayment(ctx context.Context, request PaymentRequest) (*PaymentResponse, error) {
	params := "?" + url.Values{
		"callback":      {c.callbackUrl},
		"address":       {request.Address},
		"pending":       {"1"},
		"confirmations": {"1"},
	}.Encode()

	url, _ := url.JoinPath(c.url, request.Currency, "create")

	resp, err := http.Get(url + params)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var payment *PaymentResponse
	err = json.Unmarshal([]byte(body), &payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

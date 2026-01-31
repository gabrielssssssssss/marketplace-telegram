package cryptapi

type PaymentRequest struct {
	Address  string
	Currency string
}

type PaymentResponse struct {
	Status                 string `json:"status"`
	AddressIn              string `json:"address_in"`
	AddressOut             string `json:"address_out"`
	CallbackUrl            string `json:"callback_url"`
	Priority               string `json:"priority"`
	MinimumTransactionCoin string `json:"minimum_transaction_coin"`
}

package cryptapi

type OrderRequest struct {
	callback    string
	address     string
	post        int
	json        int
	pending     int
	multi_token int
	convert     int
}

type OrderResponse struct {
	AddressIn              string  `json:"address_in"`
	AddressOut             string  `json:"address_out"`
	CallbackUrl            string  `json:"callback_url"`
	Priority               string  `json:"priority"`
	MinimumTransactionCoin float64 `json:"minimum_transaction_coin"`
	Status                 string  `json:"status"`
}

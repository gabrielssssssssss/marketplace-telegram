package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type BinanceCurrency struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func CurrencyPrice(currency string) float64 {
	upperCurrency := strings.ToUpper(currency)

	if strings.Contains(upperCurrency, "USDT") {
		upperCurrency = "EUR" + upperCurrency
	} else {
		upperCurrency = upperCurrency + "EUR"
	}

	fmt.Println(upperCurrency)
	url := "https://api.binance.com/api/v3/ticker/price?symbol=" + upperCurrency

	resp, err := http.Get(url)
	if err != nil {
		return 0.0
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0
	}

	var data BinanceCurrency

	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0.0
	}

	price, err := strconv.ParseFloat(data.Price, 64)
	if err != nil {
		return 0.0
	}

	return price
}

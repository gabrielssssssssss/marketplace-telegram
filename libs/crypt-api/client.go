package cryptapi

import (
	"net/http"
	"time"
)

type CryptAPI struct {
	url         string
	address     string
	callbackUrl string
	client      *http.Client
}

func NewCryptAPI(url, address, callback string) *CryptAPI {
	return &CryptAPI{
		url:     url,
		address: address,
		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

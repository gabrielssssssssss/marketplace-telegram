package cryptapi

import (
	"net/http"
	"time"
)

type CryptAPI struct {
	url         string
	callbackUrl string
	client      *http.Client
}

func NewCryptAPI(url, callback string) *CryptAPI {
	return &CryptAPI{
		url:         url,
		callbackUrl: callback,
		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

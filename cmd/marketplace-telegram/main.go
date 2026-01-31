package main

import (
	"context"
	"fmt"
	"os"

	cryptapi "github.com/gabrielssssssssss/marketplace-telegram/libs/crypt-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	cfg := cryptapi.NewCryptAPI("https://api.cryptapi.io/", "https://mercy-broader-civilization-suggested.trycloudflare.com/callback")

	payload := cryptapi.PaymentRequest{
		Currency: "btc",
		Address:  "bc1q2dmygchykc9yk3qpme9c2la822c4qzelg5kfjq",
	}
	resp, err := cfg.CreatePayment(context.Background(), payload)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	// gotenv.Load(".env")
	// controller.Controller()
}

package controller

import (
	"context"
	"log"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/account"
	"github.com/go-telegram/bot"
)

func Controller() {
	app, err := bot.New(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	accountController := account.NewAccountController()
	accountController.Route(app)
	app.Start(context.TODO())
}

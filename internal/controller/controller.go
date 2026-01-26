package controller

import (
	"context"
	"log"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/account"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
)

func Controller() {
	app, err := bot.New(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	accountRepository := repository.NewAccountRepository()
	accountService := service.NewAccountService(accountRepository)
	accountController := account.NewAccountController(&accountService)

	accountController.Route(app)
	app.Start(context.TODO())
}

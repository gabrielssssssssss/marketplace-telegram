package controller

import (
	"context"
	"log"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/account"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/payment"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
)

func Controller() {
	app, err := bot.New(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	database, err := config.NewPostgresDatabase()
	if err != nil {
		log.Fatal(err)
	}

	accountRepository := repository.NewAccountRepository(database)
	paymentRepository := repository.NewPaymentRepository(database)

	accountService := service.NewAccountService(accountRepository)
	paymentService := service.NewPaymentService(paymentRepository)

	accountController := account.NewAccountController(&accountService)
	paymentController := payment.NewPaymentController(&paymentService)

	accountController.Route(app)
	paymentController.Route(app)

	app.Start(context.TODO())
}

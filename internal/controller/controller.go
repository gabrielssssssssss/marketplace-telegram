package controller

import (
	"context"
	"log"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/config"
	ah "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/account/handlers"
	ph "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/payment/handlers"

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

	accountController := ah.NewAccountController(&accountService)
	paymentHandlers := ph.NewPaymentController(&paymentService)

	accountController.Handlers(app)
	paymentHandlers.Handlers(app)

	// go func() {
	// 	http.HandleFunc("/callback", webhookApi.HandleCallback)
	// 	http.ListenAndServe(":5000", nil)
	// }()

	app.Start(context.TODO())
}

package controller

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/config"
	ah "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/account/handlers"
	ph "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/payment/handlers"
	wh "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/payment/webhooks"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
)

func Controller() {
	bot, err := bot.New(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	database, err := config.NewPostgresDatabase()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(database)
	paymentRepository := repository.NewPaymentRepository(database)

	accountService := service.NewAccountService(userRepository)
	paymentService := service.NewPaymentService(paymentRepository)

	accountController := ah.NewAccountHandler(&accountService)
	paymentHandlers := ph.NewPaymentHandler(&paymentService)
	paymentWebhooks := wh.NewPaymentWebhook(&paymentService, &accountService)

	accountController.Handlers(bot)
	paymentHandlers.Handlers(bot)

	go func() {
		http.HandleFunc("/callback", paymentWebhooks.WebhookPayment)
		http.ListenAndServe(os.Getenv("CALLBACK_PORT"), nil)
	}()

	bot.Start(context.TODO())
}

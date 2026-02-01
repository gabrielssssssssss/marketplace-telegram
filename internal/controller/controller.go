package controller

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/config"
	ah "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/account/handlers"
	ph "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/payment/handlers"
	wh "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/payment/webhooks"
	rh "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/restore/handlers"

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

	accountHandlers := ah.NewAccountHandler(&accountService)
	paymentHandlers := ph.NewPaymentHandler(&paymentService)
	restoreHandlers := rh.NewRestoreHandler(&accountService)

	paymentWebhooks := wh.NewPaymentWebhook(&paymentService, &accountService)

	go func() {
		http.HandleFunc("/callback", paymentWebhooks.WebhookPayment)
		http.ListenAndServe(os.Getenv("CALLBACK_PORT"), nil)
	}()

	accountHandlers.Handlers(bot)
	paymentHandlers.Handlers(bot)
	restoreHandlers.Handlers(bot)

	bot.Start(context.TODO())
}

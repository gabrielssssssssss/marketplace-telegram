package controller

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/carts"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/products"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/users"
	ah "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/account"
	ph "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/payment"
	rh "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/restore"
	wh "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/webhooks"
	"github.com/gin-gonic/gin"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
)

func Controller() {
	gin.SetMode(gin.TestMode)
	app := gin.Default()

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
	productRepository := repository.NewProductRepository(database)
	cartRepository := repository.NewCartRepository(database)

	userService := service.NewUserService(userRepository)
	paymentService := service.NewPaymentService(paymentRepository)
	productService := service.NewProductService(productRepository)
	cartService := service.NewCartService(cartRepository)

	accountHandlers := ah.NewAccountHandler(&userService)
	paymentHandlers := ph.NewPaymentHandler(&paymentService)
	restoreHandlers := rh.NewRestoreHandler(&userService)

	paymentWebhooks := wh.NewPaymentWebhook(&paymentService, &userService)

	usersRouter := users.NewUserController(&userService)
	productsRouter := products.NewProductController(&productService)
	cartsRouter := carts.NewCartController(&cartService)

	apiGroup := app.Group("/api/v1/")

	apiGroup.Use(middlewares.CORS(), middlewares.Authorization)
	usersRouter.Route(apiGroup)
	productsRouter.Route(apiGroup)
	cartsRouter.Route(apiGroup)

	accountHandlers.Handlers(bot)
	paymentHandlers.Handlers(bot)
	restoreHandlers.Handlers(bot)

	go func() {
		http.HandleFunc("/callback", paymentWebhooks.WebhookPayment)
		http.ListenAndServe(os.Getenv("CALLBACK_PORT"), nil)
	}()

	go func() {
		bot.Start(context.TODO())
	}()

	app.Run(":6060")
}

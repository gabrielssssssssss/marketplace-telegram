package controller

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gabrielssssssssss/marketplace-telegram/config"
	"github.com/gabrielssssssssss/marketplace-telegram/docs"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/carts"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/middlewares"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/orders"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/products"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/swagger"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/controller/api/v1/users"
	tgAccount "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/account"
	tgPayment "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/payment"
	tgRestore "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/restore"
	tgWebhooks "github.com/gabrielssssssssss/marketplace-telegram/internal/controller/telegram/webhooks"
	"github.com/gin-gonic/gin"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/repository"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func Controller() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	db, err := config.NewPostgresDatabase()
	if err != nil {
		log.Fatal(err)
	}

	tgBot, err := bot.New(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	userRepo := repository.NewUserRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	productRepo := repository.NewProductRepository(db)
	cartRepo := repository.NewCartRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	userService := service.NewUserService(userRepo)
	paymentService := service.NewPaymentService(paymentRepo)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo)
	orderService := service.NewOrderService(orderRepo)

	userCtrl := users.NewUserController(&userService)
	productCtrl := products.NewProductController(&productService)
	cartCtrl := carts.NewCartController(&cartService)
	orderCtrl := orders.NewOrderController(&orderService)

	accountHandler := tgAccount.NewAccountHandler(&userService)
	paymentHandler := tgPayment.NewPaymentHandler(&paymentService)
	restoreHandler := tgRestore.NewRestoreHandler(&userService)
	webhookHandler := tgWebhooks.NewPaymentWebhook(&paymentService, &userService)

	apiV1 := engine.Group("/api/v1/")
	apiV1.Use(middlewares.CORS())
	{
		swagger.Route(apiV1)
		userCtrl.Route(apiV1)
		productCtrl.Route(apiV1)
		cartCtrl.Route(apiV1)
		orderCtrl.Route(apiV1)
	}

	accountHandler.Handlers(tgBot)
	paymentHandler.Handlers(tgBot)
	restoreHandler.Handlers(tgBot)

	go func() {
		http.HandleFunc("/callback", webhookHandler.WebhookPayment)
		http.ListenAndServe(os.Getenv("CALLBACK_PORT"), nil)
	}()

	go func() {
		tgBot.Start(context.TODO())
	}()

	engine.Run(":6060")
}

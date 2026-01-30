package payment

import (
	"context"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type PaymentController struct {
	PaymentService service.PaymentService
}

func NewPaymentController(PaymentService *service.PaymentService) PaymentController {
	return PaymentController{PaymentService: *PaymentService}
}

func (controller *PaymentController) HandlerPayment(ctx context.Context, b *bot.Bot, update *models.Update) {
	controller.PaymentService.PaymentCallback(ctx, b, update)
}

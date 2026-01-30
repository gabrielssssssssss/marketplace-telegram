package payment

import (
	"context"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

type PaymentController struct {
	PaymentService service.PaymentService
}

func NewPaymentController(PaymentService *service.PaymentService) PaymentController {
	return PaymentController{PaymentService: *PaymentService}
}

func (controller *PaymentController) HandlerPayment(ctx context.Context, b *bot.Bot, update *models.Update) {
	err := controller.PaymentService.PaymentCallback(ctx, b, update)

	if err != nil {
		log.Error().
			Err(err).
			Str("component", "PaymentController.HandlerPayment").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process payment callback")
		return
	}

	log.Info().Msg("Payment processed successfully")
}

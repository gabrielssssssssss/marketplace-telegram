package webhooks

import (
	"net/http"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
)

type PaymentWebhook struct {
	PaymentService service.PaymentService
}

func NewPaymentWebhook(PaymentService *service.PaymentService) PaymentWebhook {
	return PaymentWebhook{PaymentService: *PaymentService}
}

func (webhook *PaymentWebhook) WebhookPayment(w http.ResponseWriter, r *http.Request) {
	// err := webhook.PaymentService.PaymentCallback(ctx, b, update)

	// if err != nil {
	// 	log.Error().
	// 		Err(err).
	// 		Str("component", "PaymentController.WebhookPayment").
	// 		Msg("Failed to process webhook payment")
	// 	return
	// }

	// log.Info().Msg("webhook payment successfully")
}

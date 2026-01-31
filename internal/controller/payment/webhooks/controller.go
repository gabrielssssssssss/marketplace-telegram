package webhooks

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/messages"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/rs/zerolog/log"
)

type PaymentWebhook struct {
	PaymentService service.PaymentService
}

func NewPaymentWebhook(service *service.PaymentService) PaymentWebhook {
	return PaymentWebhook{PaymentService: *service}
}

func (webhook *PaymentWebhook) WebhookPayment(w http.ResponseWriter, r *http.Request) {
	paymentCallback := entity.PaymentCallback{
		PaymentID:          r.URL.Query().Get("payment_id"),
		AddressIn:          r.URL.Query().Get("address_in"),
		AdddressOut:        r.URL.Query().Get("address_out"),
		ValueCoin:          r.URL.Query().Get("value_coin"),
		ValueForwardedCoin: r.URL.Query().Get("value_forwarded_coin"),
		TxidIn:             r.URL.Query().Get("txid_in"),
		TxidOut:            r.URL.Query().Get("txid_out"),
		Confirmations:      r.URL.Query().Get("confirmations"),
		Status:             r.URL.Query().Get("state"),
	}

	checkPayment, err := webhook.PaymentService.FindPayment(&paymentCallback)
	if err != nil || checkPayment.ID == "" {
		log.Error().
			Err(err).
			Str("component", "webhook.PaymentWebhook.WebhookPayment").
			Str("payment_id", paymentCallback.PaymentID).
			Msg("Failed to process payment validation")
		return
	}

	confirmPayment, err := webhook.PaymentService.ConfirmPayment(&paymentCallback)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "webhook.PaymentWebhook.WebhookPayment").
			Str("payment_id", paymentCallback.PaymentID).
			Msg("Failed to process payment validation")
		return
	}

	message := fmt.Sprintf(messages.MessagePaymentConfirmed,
		confirmPayment.ValueForwardedCoin,
		strings.ToUpper(confirmPayment.Currency),
		confirmPayment.ID,
		confirmPayment.ValueForwardedCoin,
		confirmPayment.CreatedAt,
		confirmPayment.ConfirmedAt,
	)

	err = helper.SendMessage(
		os.Getenv("TELEGRAM_TOKEN"),
		strconv.FormatInt(confirmPayment.UserID, 10),
		message,
	)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "webhook.PaymentWebhook.WebhookPayment").
			Str("payment_id", paymentCallback.PaymentID).
			Msg("Failed to process payment validation")
		return
	}
}

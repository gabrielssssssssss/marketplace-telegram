package webhooks

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/messages"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/rs/zerolog/log"
)

type PaymentWebhook struct {
	PaymentService service.PaymentService
	AccountService service.AccountService
}

func NewPaymentWebhook(paymentService *service.PaymentService, accountService *service.AccountService) PaymentWebhook {
	return PaymentWebhook{PaymentService: *paymentService, AccountService: *accountService}
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
			Msg("Failed to process payment finding")
		return
	}

	confirmPayment, err := webhook.PaymentService.ConfirmPayment(&paymentCallback)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "webhook.PaymentService.ConfirmPayment").
			Str("payment_id", paymentCallback.PaymentID).
			Msg("Failed to process payment confirmation")
		return
	}

	currencyPrice := helper.CurrencyPrice(confirmPayment.Currency) * confirmPayment.ValueForwardedCoin

	user, err := webhook.AccountService.FindUser(confirmPayment.UserID)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "webhook.AccountService.FindUser").
			Str("payment_id", paymentCallback.PaymentID).
			Msg("Failed to process update user finding")
		return
	}

	updatedUser := entity.Users{
		UserId:    confirmPayment.UserID,
		Balance:   *user.Balance + currencyPrice,
		UpdatedAt: time.Now(),
	}

	_, err = webhook.AccountService.UpdateUserBalance(&updatedUser)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "webhook.AccountService.UpdateUserBalance").
			Str("payment_id", paymentCallback.PaymentID).
			Msg("Failed to process update user balance")
		return
	}

	message := fmt.Sprintf(messages.MessagePaymentConfirmed,
		confirmPayment.ValueForwardedCoin,
		strings.ToUpper(confirmPayment.Currency),
		confirmPayment.ID,
		currencyPrice,
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

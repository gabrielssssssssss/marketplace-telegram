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
		Coin:               r.URL.Query().Get("coin"),
		ValueCoin:          r.URL.Query().Get("value_coin"),
		ValueForwardedCoin: r.URL.Query().Get("value_forwarded_coin"),
		TxidIn:             r.URL.Query().Get("txid_in"),
		TxidOut:            r.URL.Query().Get("txid_out"),
		Confirmations:      r.URL.Query().Get("confirmations"),
		Status:             r.URL.Query().Get("result"),
	}

	findPayment, err := webhook.PaymentService.FindPayment(&paymentCallback)
	if err != nil || findPayment.ID == "" {
		log.Error().
			Err(err).
			Str("component", "webhook.PaymentWebhook.WebhookPayment").
			Str("payment_id", paymentCallback.PaymentID).
			Msg("Failed to process payment finding")
		return
	}

	switch paymentCallback.Status {
	case "pending":
		message := fmt.Sprintf(messages.MessagePaymentPending,
			paymentCallback.ValueCoin,
			strings.ToUpper(paymentCallback.Coin),
			paymentCallback.PaymentID,
			paymentCallback.ValueCoin,
			paymentCallback.TxidIn,
			findPayment.CreatedAt,
		)

		helper.SendMessage(
			os.Getenv("TELEGRAM_TOKEN"),
			strconv.FormatInt(findPayment.UserID, 10),
			message,
		)

		log.Info().Msg("pending payment processed successfully")

	case "sent":
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

		user, err := webhook.AccountService.FindUserByID(confirmPayment.UserID)
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
			Balance:   user.Balance + currencyPrice,
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
			confirmPayment.TxidIn,
			confirmPayment.CreatedAt,
			confirmPayment.ConfirmedAt,
		)

		helper.SendMessage(
			os.Getenv("TELEGRAM_TOKEN"),
			strconv.FormatInt(confirmPayment.UserID, 10),
			message,
		)

		log.Info().Msg("sent payment processed successfully")
	}
}

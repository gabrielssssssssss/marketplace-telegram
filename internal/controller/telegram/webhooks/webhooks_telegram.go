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
	f := func(s string) float64 { v, _ := strconv.ParseFloat(s, 64); return v }

	payment := entity.Payment{
		ID:                 r.URL.Query().Get("payment_id"),
		AddressIn:          r.URL.Query().Get("address_in"),
		AddressOut:         r.URL.Query().Get("address_out"),
		Currency:           r.URL.Query().Get("coin"),
		ValueCoin:          f(r.URL.Query().Get("value_coin")),
		ValueForwardedCoin: f(r.URL.Query().Get("value_forwarded_coin")),
		TxidIn:             r.URL.Query().Get("txid_in"),
		TxidOut:            r.URL.Query().Get("txid_out"),
		Status:             r.URL.Query().Get("result"),
	}

	findPayment, err := webhook.PaymentService.FindPayment(&payment)
	if err != nil || findPayment.ID == "" {
		log.Error().
			Err(err).
			Str("component", "webhook.PaymentWebhook.WebhookPayment").
			Str("payment_id", payment.ID).
			Msg("Failed to process payment finding")
		return
	}

	switch payment.Status {
	case "pending":
		message := fmt.Sprintf(messages.MessagePaymentPending,
			payment.ValueCoin,
			strings.ToUpper(payment.Currency),
			payment.ID,
			payment.ValueCoin,
			payment.TxidIn,
			findPayment.CreatedAt,
		)

		helper.SendMessage(
			os.Getenv("TELEGRAM_TOKEN"),
			strconv.FormatInt(findPayment.UserID, 10),
			message,
		)

		log.Info().Msg("pending payment processed successfully")

	case "sent":
		updatePayment := entity.Payment{
			ID:                 payment.ID,
			AddressIn:          payment.AddressIn,
			AddressOut:         payment.AddressOut,
			ValueCoin:          payment.ValueCoin,
			ValueForwardedCoin: payment.ValueForwardedCoin,
			TxidIn:             payment.TxidIn,
			TxidOut:            payment.TxidOut,
			ConfirmedAt:        time.Now(),
			Status:             payment.Status,
		}
		confirmPayment, err := webhook.PaymentService.ConfirmPayment(&updatePayment)
		if err != nil {
			log.Error().
				Err(err).
				Str("component", "webhook.PaymentService.ConfirmPayment").
				Str("payment_id", payment.ID).
				Msg("Failed to process payment confirmation")
			return
		}

		currencyPrice := helper.CurrencyPrice(confirmPayment.Currency) * confirmPayment.ValueForwardedCoin

		user, err := webhook.AccountService.FindUserByID(&entity.User{UserId: confirmPayment.UserID})
		if err != nil {
			log.Error().
				Err(err).
				Str("component", "webhook.AccountService.FindUser").
				Str("payment_id", payment.ID).
				Msg("Failed to process update user finding")
			return
		}

		updatedUser := entity.User{
			UserId:    confirmPayment.UserID,
			Balance:   user.Balance + currencyPrice,
			UpdatedAt: time.Now(),
		}

		_, err = webhook.AccountService.ModifyUserByID(&updatedUser)
		if err != nil {
			log.Error().
				Err(err).
				Str("component", "webhook.AccountService.UpdateUserBalance").
				Str("payment_id", payment.ID).
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

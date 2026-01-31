package service

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/messages"
	cryptapi "github.com/gabrielssssssssss/marketplace-telegram/libs/crypt-api"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *paymentServiceImpl) PaymentCallback(ctx context.Context, b *bot.Bot, update *models.Update) error {
	cb := update.CallbackQuery
	if cb == nil || cb.Message.Message == nil {
		return fmt.Errorf("callback query or associated message is nil")
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: cb.ID,
	})

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    cb.Message.Message.Chat.ID,
		MessageID: cb.Message.Message.ID,
		ParseMode: "HTML",
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "BTC", CallbackData: "payment_currency_btc"},
					{Text: "ETH", CallbackData: "payment_currency_eth"},
					{Text: "LTC", CallbackData: "payment_currency_ltc"},
				},
				{
					{Text: "SOL", CallbackData: "payment_currency_sol"},
					{Text: "TRX", CallbackData: "payment_currency_trx"},
					{Text: "USDT (erc20)", CallbackData: "payment_currency_usdt_erc20"},
				},
				{
					{Text: "USDT (trc20)", CallbackData: "payment_currency_usdt_trc20"},
					{Text: "USDT (sol)", CallbackData: "payment_currency_usdt_sol"},
					{Text: "XMR", CallbackData: "payment_currency_xmr"},
				},
				{
					{Text: "👈 Retour", CallbackData: "account"},
				},
			},
		},
		Text: messages.MessagePayment,
	})

	return nil
}

func (s *paymentServiceImpl) PaymentCurrencyCallback(ctx context.Context, b *bot.Bot, update *models.Update) error {
	cb := update.CallbackQuery
	if cb == nil || cb.Message.Message == nil {
		return fmt.Errorf("callback query or associated message is nil")
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: cb.ID,
	})

	chatID := cb.Message.Message.Chat.ID
	currency := strings.Split(cb.Data, "_")[2]

	payment := entity.Payment{
		UserID:   chatID,
		Currency: currency,
	}

	transaction, err := s.repository.CreatePayment(&payment)
	if err != nil {
		return err
	}

	callbackUrl := fmt.Sprintf("%s?payment_id=%s", os.Getenv("CALLBACK_URL"), transaction.ID)
	client := cryptapi.NewCryptAPI("https://api.cryptapi.io/", callbackUrl)

	response, err := client.CreatePayment(ctx, cryptapi.PaymentRequest{
		Address:  os.Getenv(strings.ToUpper(currency)),
		Currency: currency,
	})
	if err != nil {
		return err
	}

	message := fmt.Sprintf(messages.MessagePaymentCurrency,
		strings.ToUpper(currency),
		transaction.ID,
		response.Status,
		response.AddressOut,
		response.MinimumTransactionCoin,
		response.Priority,
	)

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    chatID,
		MessageID: cb.Message.Message.ID,
		ParseMode: "HTML",
		Text:      message,
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{{Text: "👈 Retour", CallbackData: "payment"}},
			},
		},
	})

	return nil
}

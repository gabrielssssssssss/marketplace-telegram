package payment

import (
	"context"
	"fmt"
	"strings"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/messages"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

type PaymentHandler struct {
	PaymentService service.PaymentService
}

func NewPaymentHandler(PaymentService *service.PaymentService) PaymentHandler {
	return PaymentHandler{PaymentService: *PaymentService}
}

func (handler *PaymentHandler) HandlerPayment(ctx context.Context, b *bot.Bot, update *models.Update) {
	cb := update.CallbackQuery
	if cb == nil || cb.Message.Message == nil {
		log.Error().
			Err(fmt.Errorf("callback query or associated message is nil")).
			Str("component", "handler.HandlerPayment").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process payment callback")
		return
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

	log.Info().Msg("payment processed successfully")
}

func (handler *PaymentHandler) HandlerPaymentCurrency(ctx context.Context, b *bot.Bot, update *models.Update) {
	cb := update.CallbackQuery
	if cb == nil || cb.Message.Message == nil {
		log.Error().
			Err(fmt.Errorf("callback query or associated message is nil")).
			Str("component", "handler.HandlerPaymentCurrency").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process payment callback")
		return
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: cb.ID,
	})

	chatID := cb.Message.Message.Chat.ID

	createdPayment, providerResponse, err := handler.PaymentService.CreatePayment(ctx, cb)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "handler.PaymentService.PaymentCurrencyCallback").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process payment callback")
		return
	}

	message := fmt.Sprintf(messages.MessagePaymentCurrency,
		strings.ToUpper(createdPayment.Currency),
		createdPayment.ID,
		providerResponse.Status,
		providerResponse.AddressOut,
		providerResponse.MinimumTransactionCoin,
		providerResponse.Priority,
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

	log.Info().Msg("payment_currency processed successfully")
}

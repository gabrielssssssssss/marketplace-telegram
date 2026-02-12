package payment

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/messages"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	cryptapi "github.com/gabrielssssssssss/marketplace-telegram/libs/crypt-api"
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
			Str("component", "update.CallbackQuery").
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
			Str("component", "update.CallbackQuery").
			Int64("user_id", update.Message.From.ID).
			Msg("Failed to process payment callback")
		return
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: cb.ID,
	})

	currency := strings.Split(cb.Data, "_")[2]
	chatID := cb.From.ID

	payment, err := handler.PaymentService.Register(&entity.Payment{UserID: chatID, Currency: currency})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "handler.PaymentService.RegisterPayment").
			Int64("user_id", chatID).
			Msg("Failed to process payment callback")
		return
	}

	hook := fmt.Sprintf("%s?payment_id=%s", os.Getenv("CALLBACK_URL"), payment.ID)
	client := cryptapi.NewCryptAPI("https://api.cryptapi.io/", hook)

	providerResponse, err := client.NewPayment(ctx, cryptapi.PaymentRequest{
		Address:  os.Getenv(strings.ToUpper(currency)),
		Currency: currency,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "client.NewPayment").
			Int64("user_id", chatID).
			Msg("Failed to process payment callback")
		return
	}

	text := fmt.Sprintf(messages.MessagePaymentCurrency,
		strings.ToUpper(payment.Currency),
		payment.ID,
		providerResponse.Status,
		providerResponse.AddressIn,
		providerResponse.MinimumTransactionCoin,
		providerResponse.Priority,
	)

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    chatID,
		MessageID: cb.Message.Message.ID,
		ParseMode: "HTML",
		Text:      text,
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{{Text: "👈 Retour", CallbackData: "payment"}},
			},
		},
	})

	log.Info().Msg("payment_currency processed successfully")
}

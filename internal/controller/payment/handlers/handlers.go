package payment

import "github.com/go-telegram/bot"

func (handler *PaymentHandler) Handlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "payment", bot.MatchTypeExact, handler.HandlerPayment)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "payment_currency", bot.MatchTypeContains, handler.HandlerPaymentCurrency)
}

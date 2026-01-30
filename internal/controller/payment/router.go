package payment

import "github.com/go-telegram/bot"

func (controller *PaymentController) Route(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "payment", bot.MatchTypeExact, controller.HandlerPayment)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "payment_currency", bot.MatchTypeContains, controller.HandlerPaymentCurrency)

}

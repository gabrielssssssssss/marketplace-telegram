package payment

import "github.com/go-telegram/bot"

func (controller *PaymentController) Route(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "payment", bot.MatchTypePrefix, controller.HandlerPayment)
}

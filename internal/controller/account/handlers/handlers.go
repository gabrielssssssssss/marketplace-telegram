package account

import "github.com/go-telegram/bot"

func (controller *AccountController) Handlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, controller.HandlerStart)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "account", bot.MatchTypePrefix, controller.HandlerAccount)

}

package account

import "github.com/go-telegram/bot"

func (handler *AccountHandler) Handlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handler.HandlerStart)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "account", bot.MatchTypePrefix, handler.HandlerAccount)
}

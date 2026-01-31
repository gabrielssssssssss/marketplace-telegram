package handlers

import "github.com/go-telegram/bot"

func (handler *RestoreHandler) Handlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "restore", bot.MatchTypeContains, handler.HandlerRestore)

}

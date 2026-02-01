package handlers

import "github.com/go-telegram/bot"

func (handler *RestoreHandler) Handlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "restore", bot.MatchTypeExact, handler.HandlerRestore)
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, handler.ListenerRestore)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "restore_transfer_", bot.MatchTypeContains, handler.HandlerRestoreTransfer)
}

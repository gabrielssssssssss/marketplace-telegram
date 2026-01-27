package account

import "github.com/go-telegram/bot"

func (controller *AccountController) Route(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, controller.Start)
}

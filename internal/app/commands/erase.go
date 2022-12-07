package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Erase(inputMessage *tgbotapi.Message) {
	c.productService.Erase()

	outputMsg := "Все чейны были удалены. Можете начать заново"
	outputMsg += "\n"

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	c.bot.Send(msg)
}

package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Show(inputMessage *tgbotapi.Message) {
	coord := c.productService.Show()

	outputMsg := fmt.Sprintf("Сейчас воркер на чейне %d на расстоянии %d м от AC1", coord.ChainID, coord.Dist)
	outputMsg += "\n"

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	c.bot.Send(msg)
}

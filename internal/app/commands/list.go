package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {

	products := c.productService.List()
	var outputMsg string
	if len(products) == 0 {
		outputMsg = "Вы еще не добавили ни одного чейна\n"
	} else {
		outputMsg = "Список всех чейнов на горизонте:\n\n"
		for _, p := range products {
			outputMsg += fmt.Sprintf("%v\t%v\t%v\t%v", p.ID, p.AC1, p.AC2, p.Dist)
			outputMsg += "\n"
		}

		outputMsg += fmt.Sprintf("%v", c.productService.GetPoints())
		outputMsg += "\n"

	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	c.bot.Send(msg)
}

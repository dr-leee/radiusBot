package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

// функция добавления чейнов
func (c *Commander) Add(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx := strings.Fields(args)

	if len(idx) != 3 {
		log.Println("wrong args", args)
		return
	} else {
		ac1, _ := strconv.Atoi(idx[0])
		ac2, _ := strconv.Atoi(idx[1])
		dist, _ := strconv.Atoi(idx[2])
		c.productService.Add(ac1, ac2, dist)

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Чейн добавлен")

		c.bot.Send(msg)

		//пересчитываем таблицу соответствий
		c.productService.Calc()

	}

}

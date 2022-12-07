package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (c *Commander) Set(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx := strings.Fields(args)
	var mess string
	if len(idx) != 2 {
		log.Println("wrong args", args)
		return
	} else {
		chainID, _ := strconv.Atoi(idx[0])
		dist, _ := strconv.Atoi(idx[1])
		chains := c.productService.List()
		//проверяем, что переместили в существующий чейн
		if chainID >= len(chains) || chainID < 0 {
			mess = "Такого чейна не существует"
		} else {
			c.productService.Set(chainID, dist)
			mess = "Воркер перемещен. Наберите /show, чтобы увидеть его координаты"
		}

	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)

	c.bot.Send(msg)
}

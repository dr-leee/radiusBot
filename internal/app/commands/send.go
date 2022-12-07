package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (c *Commander) Send(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx := strings.Fields(args)
	var mess string

	if len(idx) != 2 {
		log.Println("wrong args", args)
		return
	} else {
		chainID, _ := strconv.Atoi(idx[0])
		dist, _ := strconv.Atoi(idx[1])
		chainID, dist, err := c.productService.Move(chainID, dist)
		if err != nil {
			log.Panic(err)
		}
		c.productService.Set(chainID, dist)
		mess = fmt.Sprintf("Воркер перемещен. Теперь он находится в чейне %d на расстоянии %d от AC1", chainID, dist)

	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)

	c.bot.Send(msg)
}

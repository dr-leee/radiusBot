package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help -- помощь\n"+
			"/list -- показать все чейны\n"+
			"/add 'ac1 ac2 distance' -- добавить новый чейн: 30 31 100\n"+
			"/show -- показать текущую координату воркера: ID чейна и расстояние от ac1\n"+
			"/set 'chainID distance' -- поставить воркера в позицию принудительно\n"+
			"/send 'chainID distance' -- отправить сигнал системе из новой точки\n"+
			"/erase -- удалить все чейны и начать заново\n",
	)

	c.bot.Send(msg)
}

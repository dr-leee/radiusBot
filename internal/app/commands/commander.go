package commands

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"radiusBot/internal/app/service/chain"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *chain.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *chain.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

type CommandData struct {
	Offset int `json:"offset"`
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Recover from panic: %v", panicValue)
		}
	}()
	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed: %+v\n", parsedData),
		)
		c.bot.Send(msg)
		return
	}

	if update.Message != nil { // If we got a message

		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "show":
			c.Show(update.Message)
		case "add":
			c.Add(update.Message)
		case "set":
			c.Set(update.Message)
		case "erase":
			c.Erase(update.Message)
		case "send":
			c.Send(update.Message)
		default:
			c.Default(update.Message)
		}

	}

}

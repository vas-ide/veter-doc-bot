package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/vas-ide/veter-doc-bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string] func(c *Commander, msg *tgbotapi.Message) {}

type Commander struct {
	bot *tgbotapi.BotAPI
    productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot: bot,
        productService: productService,
	}
}



func (c Commander) HandleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil{
			log.Printf("Recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := product.Animal{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, fmt.Sprintf("Data update - %v", parsedData))
		c.bot.Send(msg)
		return 
	}



	if update.Message == nil { // ignore any non-Message updates
		return
	}

	if !update.Message.IsCommand() { // ignore any non-command Messages
		c.Non(update.Message)
		return
	}

	command, ok := registeredCommands[update.Message.Command()]
	if ok {
		command(&c, update.Message)
	}else {
		c.Deffault(update.Message)
	}

	// switch update.Message.Command() {
	// case "help":
	// 	c.Help(update.Message)
	// case "dog":
	// 	c.Dog(update.Message)
	// case "cat":
	// 	c.Cat(update.Message)
	// case "other":
	// 	c.Other(update.Message)
	// case "product":
	// 	c.Products(update.Message)
	// case "pavel":
	// 	c.Pavel(update.Message)
	// default:
	// 	c.Deffault(update.Message)
	// }

}











package commands

import (

	"github.com/vas-ide/veter-doc-bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
	if update.Message == nil { // ignore any non-Message updates
		return
	}

	if !update.Message.IsCommand() { // ignore any non-command Messages
		c.Non(update.Message)
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "dog":
		c.Dog(update.Message)
	case "cat":
		c.Cat(update.Message)
	case "other":
		c.Other(update.Message)
	case "product":
		c.Products(update.Message)
	case "pavel":
		c.Pavel(update.Message)
	default:
		c.Deffault(update.Message)
	}

}











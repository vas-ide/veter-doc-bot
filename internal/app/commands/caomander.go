package commands

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
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

func (c Commander) Help (inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("What's you pet? \n/dog \n/cat \n/other"))
	c.bot.Send(msg)
}

func (c Commander) Dog (inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("DOGI"))
	c.bot.Send(msg)
}

func (c Commander) Cat (inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("CATI"))
	c.bot.Send(msg)
}

func (c Commander) Other (inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("NOT"))
	c.bot.Send(msg)
}

func (c Commander) Products (inputMessage *tgbotapi.Message) {
	products := c.productService.LstService()
	newStr := "All available products: \n"
	for _, v := range products {
		newStr = fmt.Sprintf(newStr+"\n%s", v.Title)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%s", newStr))
	c.bot.Send(msg)
}

func (c Commander) Pavel (inputMessage *tgbotapi.Message) {
    config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}
	// qrterminal.GenerateWithConfig("8-918-453-01-57 Kseniy best Veterinar in Eysk.", config)
	qrterminal.GenerateWithConfig("Pavel viebal TRUMP & MELANIA continue - 'https://gkovd.ru/branches/south-air-navigation/'.", config)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Terminal"))
	c.bot.Send(msg)
    
}

func (c Commander) Deffault (inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("We don't know that command. Pls give mo actual information for our assistance."))
	c.bot.Send(msg)
}

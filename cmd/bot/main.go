package main

import (
	"log"
	"os"

	"github.com/vas-ide/veter-doc-bot/internal/app/commands"
	"github.com/vas-ide/veter-doc-bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()
    commander := commands.NewCommander(bot,productService)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Info - "+update.Message.Text)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		}

		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "dog":
			commander.Dog(update.Message)
		case "cat":
			commander.Cat(update.Message)
		case "other":
			commander.Other(update.Message)
		case "product":
			commander.Products(update.Message)
        case "pavel":
			commander.Pavel(update.Message)
		default:
			commander.Deffault(update.Message)
		}

	}
}

// type Commander struct {
// 	bot *tgbotapi.BotAPI
//     productService *product.Service
// }

// func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
// 	return &Commander{
// 		bot: bot,
//         productService: productService,
// 	}
// }

// func (c Commander) Help (inputMessage *tgbotapi.Message) {
// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("What's you pet? \n/dog \n/cat \n/other"))
// 	c.bot.Send(msg)
// }

// func (c Commander) Dog (inputMessage *tgbotapi.Message) {
// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("DOGI"))
// 	c.bot.Send(msg)
// }

// func (c Commander) Cat (inputMessage *tgbotapi.Message) {
// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("CATI"))
// 	c.bot.Send(msg)
// }

// func (c Commander) Other (inputMessage *tgbotapi.Message) {
// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("NOT"))
// 	c.bot.Send(msg)
// }

// func (c Commander) Products (inputMessage *tgbotapi.Message) {
// 	products := c.productService.LstService()
// 	newStr := "All available products: \n"
// 	for _, v := range products {
// 		newStr = fmt.Sprintf(newStr+"\n%s", v.Title)
// 	}
// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%s", newStr))
// 	c.bot.Send(msg)
// }

// func (c Commander) Pavel (inputMessage *tgbotapi.Message) {
//     config := qrterminal.Config{
// 		Level:     qrterminal.M,
// 		Writer:    os.Stdout,
// 		BlackChar: qrterminal.WHITE,
// 		WhiteChar: qrterminal.BLACK,
// 		QuietZone: 1,
// 	}
// 	// qrterminal.GenerateWithConfig("8-918-453-01-57 Kseniy best Veterinar in Eysk.", config)
// 	qrterminal.GenerateWithConfig("Pavel viebal TRUMP & MELANIA continue - 'https://gkovd.ru/branches/south-air-navigation/'.", config)
// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Terminal"))
// 	c.bot.Send(msg)
    
// }

// func (c Commander) Deffault (inputMessage *tgbotapi.Message) {
// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("We don't know that command. Pls give mo actual information for our assistance."))
// 	c.bot.Send(msg)
// }

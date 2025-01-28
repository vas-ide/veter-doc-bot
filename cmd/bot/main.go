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


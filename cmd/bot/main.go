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
        commander.HandleUpdate(update)
    }
}


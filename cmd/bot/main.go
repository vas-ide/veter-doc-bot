package main

import (
	"log"
	"os"


    "github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)




func main() {

    godotenv.Load()


    bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true

    log.Printf("Authorized on account %s", bot.Self.UserName)


	u :=tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	
    

    updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	

    for update := range updates {
        if update.Message == nil { // ignore any non-Message updates
            continue
        }

        if !update.Message.IsCommand() { // ignore any non-command Messages
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Info - " + update.Message.Text)
            if _, err := bot.Send(msg); err != nil {
                log.Panic(err)
            }
            continue
        }


        // Create a new MessageConfig. We don't have text yet,
        // so we leave it empty.
        // msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Info - ")
        // msg.ReplyToMessageID = update.Message.MessageID
        // bot.Send(msg)

        // Extract the command from the Message.
        switch update.Message.Command() {
        case "help":
            helpCommand(bot, update.Message)
        case "dog":
            dogCommand(bot, update.Message)
        case "cat":
            catCommand(bot, update.Message)
        case "other":
            otherCommand(bot, update.Message)
        default:
            deffaultBehavior(bot, update.Message)
        }
   

        // if _, err := bot.Send(msg); err != nil {
        //     log.Panic(err)
        // }
    }
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "What's you pet? /dog /cat /other")
    bot.Send(msg)
}

func dogCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "DOGI")
    bot.Send(msg)
}

func catCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "CATI")
    bot.Send(msg)
}

func otherCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "NOT")
    bot.Send(msg)
}

func deffaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "We don't know that command. Pls give mo actual information for our assistance.")
    bot.Send(msg)
}




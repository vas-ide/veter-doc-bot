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

        // if !update.Message.IsCommand() { // ignore any non-command Messages
        //     continue
        // }


        // Create a new MessageConfig. We don't have text yet,
        // so we leave it empty.
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Info - " + update.Message.Text)
        msg.ReplyToMessageID = update.Message.MessageID
        bot.Send(msg)

        // Extract the command from the Message.
        // switch update.Message.Command() {
        // case "help":
        //     msg.Text = "I understand /sayhi and /status."
        // case "sayhi":
        //     msg.Text = "Hi :)"
        // case "status":
        //     msg.Text = "I'm ok."
        // default:
        //     msg.Text = "I don't know that command"
        // }
   

        // if _, err := bot.Send(msg); err != nil {
        //     log.Panic(err)
        // }
    }
}





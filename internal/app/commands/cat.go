package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)


func (c Commander) Cat (inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("CATI"))
	c.bot.Send(msg)
}



func init() {
	registeredCommands["cat"] = (*Commander).Cat 
}

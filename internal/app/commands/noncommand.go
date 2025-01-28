package commands


import (
	"fmt"


	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)


func (c Commander) Non (inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("'%v' --- Message", inputMessage.Text))
	c.bot.Send(msg)
}


func init() {
	registeredCommands["non"] = (*Commander).Non
}

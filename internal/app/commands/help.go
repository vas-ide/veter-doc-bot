package commands


import (
	"fmt"
	
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)



func (c Commander) Help (inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("What's you pet? \n/dog \n/cat \n/other"))
	c.bot.Send(msg)
}



func init() {
	registeredCommands["help"] = (*Commander).Help 
}



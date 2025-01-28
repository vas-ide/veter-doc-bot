package commands



import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)


func (c Commander) Deffault (inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("'%v' --- We don't know that command. Pls give mo actual information for our assistance.", inputMessage.Text))
	c.bot.Send(msg)
}

func init() {
	registeredCommands["deffault"] = (*Commander).Deffault 
}

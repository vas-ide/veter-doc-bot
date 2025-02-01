package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)


func (c Commander) Cat (inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("CATI"))


	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Cat-M","Кот"),
			tgbotapi.NewInlineKeyboardButtonData("Cat-W","Кошка"),
		),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData("Dog-M","Кабель"),
		// 	tgbotapi.NewInlineKeyboardButtonData("Dog-w","Сука"),
		// ),
	)

	c.bot.Send(msg)



}



func init() {
	registeredCommands["cat"] = (*Commander).Cat 
}

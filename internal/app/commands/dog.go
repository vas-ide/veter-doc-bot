package commands

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	// "github.com/google/uuid"
	"github.com/vas-ide/veter-doc-bot/internal/service/product"
)



func (c Commander) Dog (inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("DOGI"))

	serializedData, _ := json.Marshal(product.Animal{
		Name: "Gerda",
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData("Cat-M","Кот"),
		// 	tgbotapi.NewInlineKeyboardButtonData("Cat-W","Кошка"),
		// ),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Dog-M","Кабель"),
			// tgbotapi.NewInlineKeyboardButtonData("Dog-W",fmt.Sprint("Сука %v", string(serializedData))),
			tgbotapi.NewInlineKeyboardButtonData("Dog-W",string(serializedData))),
	)
	c.bot.Send(msg)
}


func init() {
	registeredCommands["dog"] = (*Commander).Dog 
}


package commands

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"


	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)


func (c Commander) Getqr (inputMessage *tgbotapi.Message) {
    config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}
	// qrterminal.GenerateWithConfig("8-918-453-01-57 Kseniy best Veterinar in Eysk.", config)
	qrterminal.GenerateWithConfig("Pavel viebal TRUMP & MELANIA continue - 'https://gkovd.ru/branches/south-air-navigation/'.", config)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Terminal"))
	c.bot.Send(msg)
	msgQR := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("QR"))
	c.bot.Send(msgQR)
    
}



func init() {
	registeredCommands["getqr"] = (*Commander).Getqr 
}

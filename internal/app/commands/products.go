package commands



import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)


func (c Commander) Products (inputMessage *tgbotapi.Message) {
	products := c.productService.LstService()
	newStr := "All available products: \n"
	for _, v := range products {
		newStr = fmt.Sprintf(newStr+"\n%s", v.Title)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%s", newStr))
	c.bot.Send(msg)
}


func init() {
	registeredCommands["product"] = (*Commander).Products
}

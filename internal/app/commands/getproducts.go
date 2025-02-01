package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func (c Commander) GetProducts (inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Wrong args ", args)
		return
	}

	product, err := c.productService.GetProducts(idx)
	if err != nil {
		log.Println("Fail to get product ", idx, err)
		return
	}


	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Successfully pars argument: %v", product.Title))

	c.bot.Send(msg)
}


func init() {
	registeredCommands["getpr"] = (*Commander).GetProducts
}

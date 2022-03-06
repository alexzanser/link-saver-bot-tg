package telegram

import "github.com/alexzanser/telegramBotGo.git/clients/telegram"

type Processor struct {
	tg *telegram.Client
	offset int
	//storage
}

func New(client *telegram )
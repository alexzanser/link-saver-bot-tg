package main

import (
	"flag"
	"log"

	eventconsumer "github.com/alexzanser/telegramBotGo.git/clients/consumer/event-consumer"
	"github.com/alexzanser/telegramBotGo.git/events/telegram"
	tgClient "github.com/alexzanser/telegramBotGo.git/clients/telegram"
	"github.com/alexzanser/telegramBotGo.git/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files/storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

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
	tgBotHost = "api.Telegram.org"
	storagePath = "storage"
	batchSize = 100
)
func main() {
	tgClient := tgClient.New(tgBotHost, mustToken())
	eventsProcessor := telegram.New(&tgClient, files.New(storagePath))
	log.Printf("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func	mustToken() string {
	token := flag.String(
				"token-bot-token", 
				"", 
				"token for access to telegram bot ")
	flag.Parse()
	
	if *token == "" {
		log.Fatal(" token is invalid")
	}
	return ""
}

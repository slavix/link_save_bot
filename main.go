package main

import (
	"flag"
	"linkSaveBot/clients/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	_ = telegram.New(tgBotHost, mustToken())
	// fetcher = fetcher.New()
	// processor = processor.New()
	// comsumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

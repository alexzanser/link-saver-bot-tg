package telegram

import (
	"log"
	"net/url"
	"strings"

	"github.com/alexzanser/telegramBotGo.git/lib/e"
	"github.com/alexzanser/telegramBotGo.git/storage"
)

const (
	RndCmd = "/rnd"
	HelpCmd = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	//add page: 
	//rnd page : /rnd
	// help : /help
	//start /start : hi + help

	if isAddCmd() {
		//TODO: AddPage()

	}
	switch text {
	case RndCmd:
	case HelpCmd:
	case StartCmd:
	default:

	}
	
}

func (p *Processor) savePage(chatID int, pageURL string, username string) (err error) {
	defer func() {err = e.WrapIfErr("can't do command: save page", err)} ()

	page := &storage.Page{
		URL: pageURL,
		UserName: username,
	}

	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}

	if isExists {
		return p.tg.SendMessage(chatID, "")
	}
}

func isAddCmd(text string) bool {
	return isUrl(text)
}

func isUrl(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}
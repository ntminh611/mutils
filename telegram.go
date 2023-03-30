package mutils

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	Bot      tgbotapi.BotAPI
	token    string
	Received tgbotapi.UpdatesChannel
	mutex    *sync.Mutex
}

func NewTelegramBot(token string) *Telegram {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	bot.Buffer = 1000
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	rec := bot.GetUpdatesChan(u)
	return &Telegram{Bot: *bot, Received: rec, token: token, mutex: &sync.Mutex{}}
}

func (tl *Telegram) Send(msg tgbotapi.MessageConfig) (tgbotapi.Message, error) {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()
	msg.DisableWebPagePreview = true
	return tl.send(msg)
}

func (tl *Telegram) send(msg tgbotapi.Chattable) (tgbotapi.Message, error) {
	rMsg, e := tl.Bot.Send(msg)
	if e != nil {
		log.Println("Error TG by Error: ", e.Error(), tl.token)
		tooMany := "Too Many Requests: retry after "
		var tgErr *tgbotapi.Error
		if errors.As(e, &tgErr) {
			if strings.Contains(tgErr.Message, "Too Many Requests") && tgErr.RetryAfter != 0 {
				log.Println("Telegram: Too Many Requests, retry after", tgErr.RetryAfter, tl.token)
				time.Sleep(time.Second * time.Duration(tgErr.RetryAfter))
				return tl.send(msg)
			} else {
				log.Println("Error TG by Error: ", tgErr.Message, tgErr.RetryAfter)
			}
		} else if strings.Contains(e.Error(), tooMany) {
			valStr := strings.ReplaceAll(e.Error(), tooMany, "")
			sec, err := strconv.Atoi(strings.TrimSpace(valStr))
			if err == nil {
				log.Println("Telegram: Too Many Requests, retry after", sec, tl.token)
				time.Sleep(time.Second * time.Duration(sec))
				return tl.send(msg)
			} else {
				log.Println("Error TG by Error: ", e.Error(), sec, tl.token)
			}
		} else {
			log.Println("Error TG by Error: ", e.Error())
		}
	}
	return rMsg, e
}

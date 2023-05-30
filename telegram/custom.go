package telegram

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramMessageOptions struct {
	ChatID  int64
	TopicID int
	Message string
}

type CustomTgMessage interface {
	params() (tgbotapi.Params, error)
	method() string
}

func (options TelegramMessageOptions) params() (tgbotapi.Params, error) {
	params := make(tgbotapi.Params)
	params.AddFirstValid("chat_id", options.ChatID)
	params.AddFirstValid("message_thread_id", options.TopicID)
	params.AddNonEmpty("text", options.Message)
	params.AddNonEmpty("parse_mode", "HTML") // refer to "formatting-options" of telegram API
	//params.AddNonEmpty("parse_mode", "markdown")     // refer to "formatting-options" of telegram API
	params.AddBool("disable_web_page_preview", true) // I didn't need link previews, leaving here for the sake of example

	return params, nil
}

func (options TelegramMessageOptions) method() string {
	return "sendMessage"
}

func (tl *Telegram) sendCustom(c CustomTgMessage) (tgbotapi.Message, error) {
	resp, err := tl.request(c)
	if err != nil {
		return tgbotapi.Message{}, err
	}

	var message tgbotapi.Message
	err = json.Unmarshal(resp.Result, &message)

	return message, err
}

func (tl *Telegram) request(c CustomTgMessage) (*tgbotapi.APIResponse, error) {
	params, err := c.params()
	if err != nil {
		return nil, err
	}

	return tl.Bot.MakeRequest(c.method(), params)
}

func (tl *Telegram) SendTelegramTopic(options TelegramMessageOptions) error {
	_, err := tl.sendCustom(options)
	if err != nil {
		return fmt.Errorf("ERROR WHILE SENDING MESSAGE TO TELEGRAM: %v", err)
	}
	return nil
}

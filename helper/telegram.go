package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func SendMessage(botToken string, chatID string, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	body := Message{
		ChatID:    chatID,
		Text:      text,
		ParseMode: "HTML",
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return err
	}

	return nil
}

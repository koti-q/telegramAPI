package telegramApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BotAPI struct {
	Token string
	URL   string
}

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageID int    `json:"message_id"`
	Text      string `json:"text"`
	Chat      Chat   `json:"chat"`
}

type Chat struct {
	ID string `json:"id"`
}

func NewBotAPI(token string) *BotAPI {
	return &BotAPI{
		Token: token,
		URL:   fmt.Sprintf("https://api.telegram.org/bot%s/", token),
	}
}

func (bot *BotAPI) GetUpdates(offset int) ([]Update, error) {
	response, err := http.Get(fmt.Sprintf("%sgetUpdates?offset=%d", bot.URL, offset))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Ok      bool     `json:"ok"`
		Updates []Update `json:"result"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Updates, nil
}

func (bot *BotAPI) SendMessange(chatID string, text string) (bool, error) {
	body, _ := json.Marshal(map[string]string{
		"chat_id": chatID,
		"text":    text,
	})

	response, err := http.Post(
		fmt.Sprintf("%ssendMessage", bot.URL),
		"application/json",
		bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	return true, nil
}

package slack

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Webhook struct {
	Url string
}

type Message struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

func (webhook Webhook) Send(message Message) bool {
	jsonBytes, _ := json.Marshal(message)

	res, err := http.PostForm(
		webhook.Url,
		url.Values{"payload": {string(jsonBytes)}},
	)
	if err != nil {
		log.Print(err)
		return false
	}
	defer res.Body.Close()
	return true
}

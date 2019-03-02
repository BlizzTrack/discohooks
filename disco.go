package discohooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Send lets you send a message to a given webhook URL
func Send(url string, payload *WebhookParams) error {
	return request(url, payload)
}

// SendByID lets you send to the webhook if you know have it's guild and token
func SendByID(guildID, token string, payload *WebhookParams) error {
	return request(fmt.Sprintf("https://discordapp.com/api/webhooks/%s/%s", guildID, token), payload)
}

func request(url string, payload *WebhookParams) error {
	jsonValue, _ := json.Marshal(payload)

	_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

	return err
}

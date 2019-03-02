package main

import (
	"fmt"
	"log"

	"github.com/blizztrack/discohooks"
)

func main() {
	payload := &discohooks.WebhookParams{
		Content:  "**Test Payload**",
		Username: "McHookingFace",
		Embeds:   make([]*discohooks.MessageEmbed, 0),
	}

	for i := 0; i < 5; i++ {
		payload.Embeds = append(payload.Embeds, &discohooks.MessageEmbed{
			Title:       fmt.Sprintf("%v - title", i),
			Description: fmt.Sprintf("I am detail for %v", i),
		})
	}

	log.Println(discohooks.SendByID("551192923836907521", "zYUtKkE-FJJ98Mcsnzye9-GfCcHKuXorcmegt2sfuRGN4B-V273--JCOZcbKsqn_L-SK", payload))
}

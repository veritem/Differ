package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func main() {

	api := slack.New("xoxb-1765374764241-1758951736772-8hRkefHN5ZzVZ8dWXBRS5I2y")

	attachment := slack.Attachment{
		Pretext: "Happy Birth day Makuza Mugabo Verite",
		Text:    "Thanks for being cool",
	}

	channelID, timeStamp, err := api.PostMessage("#general", slack.MsgOptionText("Hello @channel", false), slack.MsgOptionAttachments(attachment), slack.MsgOptionAsUser(true))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(channelID, timeStamp)

	fmt.Println("Hello from birthday Bot")
}

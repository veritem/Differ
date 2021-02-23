package lib

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

//SchedulesMessage message types to send
type SchedulesMessage struct {
	text      string
	postAt    int64
	channelID string
}

var schedulesMessages = map[string]SchedulesMessage{
	"MessageOne": {
		text:      "message here",
		postAt:    time.Date(2021, time.February, 23, 10, 50, 0, 0, time.Local).Unix(),
		channelID: "CG3A11ZNG",
	},
}

//HandleScheduled hander
func HandleScheduled() {
	tokenErr := godotenv.Load()

	// names := "Makuza Mugabo Verite"

	if tokenErr != nil {
		log.Fatal("Error while loading .env file")
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionLog(log.New(os.Stdout, "slack-bot:", log.Lshortfile|log.LstdFlags)))

	for _, element := range schedulesMessages {
		_, _, err := api.ScheduleMessage(element.channelID, fmt.Sprint(element.postAt),
			slack.MsgOptionBlocks(
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("It's his Birth Day By The Way!"), false, false), nil, nil)))
		if err != nil {
			panic(err)
		}
	}

	for _, element := range schedulesMessages {
		scheduledMsg, _, err := api.GetScheduledMessages(&slack.GetScheduledMessagesParameters{
			Channel: element.channelID,
			Limit:   10,
		})

		fmt.Println("Schedules message", scheduledMsg)

		if err != nil {
			fmt.Println("Error while scheduling", err)
		}
	}
}

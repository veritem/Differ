package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

//SchedulesMessage message types to send
type SchedulesMessage struct {
	text      string
	postAt    int64
	channelID string
}

var schedulesMessages = map[string]SchedulesMessage{
	"MessageOne": {
		text:      "Scheduled at 20:13",
		postAt:    time.Date(2021, time.February, 20, 20, 17, 0, 0, time.Local).Unix(),
		channelID: "C01NUH9UBDW",
	},
	// "MessageTwo": {
	// 	text:      "Hello 2",
	// 	postAt:    time.Date(2021, time.February, 23, 8, 0, 0, 0, time.Local).Unix(),
	// 	channelID: "C01MYDFT51D",
	// },
}

func main() {

	// bdTime :=

	// fmt.Println(bdTime)

	// MessageOne := SchedulesMessage{
	// 	text:      "Happy Birth day <@%s>mugaboverite :tada:",
	// 	postAt:    time.Now().Local().Add(time.Second * 30).Unix(),
	// 	channelID: "C01MYDFT51D",
	// }

	tokenErr := godotenv.Load()

	if tokenErr != nil {
		log.Fatal("Error while loading .env file")
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionLog(log.New(os.Stdout, "slack-bot:", log.Lshortfile|log.LstdFlags)))

	// handles message scheduling

	// return key and elements
	for _, element := range schedulesMessages {

		// fmt.Sprintf("Happy Birth day <@%s> :tada:", "mugaboverite")

		_, _, err := api.ScheduleMessage(element.channelID, fmt.Sprint(element.postAt), slack.MsgOptionText(element.text, false))

		if err != nil {
			fmt.Println("Scheduling")
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Hello from our bot")
	})

	http.HandleFunc("/slack/events", func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sToken, err := slack.NewSecretsVerifier(r.Header, os.Getenv("SIGNING_SCRET"))

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if _, err := sToken.Write(body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err := sToken.Ensure(); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		eventsAPI, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if eventsAPI.Type == slackevents.URLVerification {
			var r *slackevents.ChallengeResponse
			err := json.Unmarshal([]byte(body), &r)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "text")
			w.Write([]byte(r.Challenge))
		}

		if eventsAPI.Type == slackevents.Message {
			var r *slackevents.ChallengeResponse

			err := json.Unmarshal([]byte(body), &r)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		}

		if eventsAPI.Type == slackevents.CallbackEvent {
			innerEvent := eventsAPI.InnerEvent

			// handle all events here
			switch ev := innerEvent.Data.(type) {

			case *slackevents.AppMentionEvent:
				api.PostMessage(ev.Channel, slack.MsgOptionText("Hello @Makuza Mugabo Verite", false))
			case *slackevents.MessageEvent:
				_, _, err := api.PostMessage("#tests", slack.MsgOptionText("Hello @verite", false))
				if err != nil {
					fmt.Println(err)
				}
				//  := ev.User
			default:
				// handle defaults
			}
		}
	})

	fmt.Println("[INFO] Server started listerning on port 3000")
	http.ListenAndServe(":3000", nil)
}

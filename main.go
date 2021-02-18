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
	postAt    string
	channelID string
}

func main() {

	MessageOne := SchedulesMessage{
		text: "Scheduled message",
		postAt: time.Now().Local().Add(time.Hour*time.Duration(1) +
			time.Minute*time.Duration(1) +
			time.Second*time.Duration(1)).String(),
		channelID: "C01MYDFT51D",
	}

	tokenErr := godotenv.Load()

	if tokenErr != nil {
		log.Fatal("Error while loading .env file")
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionLog(log.New(os.Stdout, "slack-bot:", log.Lshortfile|log.LstdFlags)))

	// handles message scheduling
	_, _, err := api.ScheduleMessage(MessageOne.channelID, MessageOne.postAt, slack.MsgOptionText(MessageOne.text, false))

	if err != nil {
		fmt.Println("Ërror while scheduling", err)
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

				userID := ev.User

				fmt.Print("USER ID", userID)

			default:
				// handle defaults
			}
		}

	})

	fmt.Println("[INFO] Server started listerning on port 3000")
	http.ListenAndServe(":3000", nil)
}

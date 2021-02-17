package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func main() {

	tokenErr := godotenv.Load()

	if tokenErr != nil {
		log.Fatal("Error while loading .env file")
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionLog(log.New(os.Stdout, "slack-bot:", log.Lshortfile|log.LstdFlags)))

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

		// fmt.Println("Event type ", eventsAPI.Type)

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

			// api.PostMessage("#general", slack.MsgOptionMeMessage())

			fmt.Println("Message event ", body)
		}

		if eventsAPI.Type == slackevents.CallbackEvent {
			innerEvent := eventsAPI.InnerEvent

			// handle all events here
			switch ev := innerEvent.Data.(type) {
			case *slackevents.AppMentionEvent:
				api.PostMessage(ev.Channel, slack.MsgOptionText("Hello @Makuza Mugabo Verite", false))
			case *slackevents.MessageEvent:

				// if a nwe message is posted send a message to general channel

				_, _, err := api.PostMessage("#general", slack.MsgOptionText("Hello @verite", false))

				if err != nil {
					fmt.Println(err)
				}

				// fmt.Printf("chanelId: %s, timeStamp: %s", channelID, timestamp)
			default:
				// fmt.Println(ev)
			}
		}

	})

	fmt.Println("[INFO] Server started listerning on port 3000")
	http.ListenAndServe(":3000", nil)
}

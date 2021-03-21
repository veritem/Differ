package lib

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

//HandleEvents handler for all of our events
func HandleEvents(w http.ResponseWriter, r *http.Request) {

	tokenErr := godotenv.Load()

	if tokenErr != nil {
		panic(tokenErr)
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionLog(log.New(os.Stdout, "slack-bot:", log.Lshortfile|log.LstdFlags)))

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sToken, err := slack.NewSecretsVerifier(r.Header, os.Getenv("SIGNING_SECRET"))

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
		// w.Write([]byte(r.Challenge))
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

		//handle all events here
		switch ev := innerEvent.Data.(type) {

		case *slackevents.AppMentionEvent:
			api.PostMessage(ev.Channel, slack.MsgOptionText("Hello @Makuza Mugabo Verite", false))
			break
		case *slackevents.MessageEvent:
			_, _, err := api.PostMessage("#tests", slack.MsgOptionText("Hello @verite", false))
			if err != nil {
				fmt.Println(err)
			}

			fmt.Print("Hello")
			//  := ev.User
			break
		default:
			// handle defaults
			fmt.Print("Event here")
		}

	}
}

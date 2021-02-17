package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

// var apiToken string = "xoxb-1765374764241-1758951736772-8hRkefHN5ZzVZ8dWXBRS5I2y"

// var apiToken string = "xapp-1-A01NL6V5X8R-1761778692548-f5bcc36cadceac0b8e4dfc61d855edb1ba4e49c9e5f8ce94d2a05692a07bc4c9"

func main() {

	// api := slack.New(apiToken, slack.OptionDebug(true), slack.OptionLog(log.New(os.Stdout, "slack-bot:", log.Lshortfile|log.LstdFlags)))

	// rtm := api.NewRTM()

	// go rtm.ManageConnection()

	// for msg := range rtm.IncomingEvents {
	// 	fmt.Println("Event recieved", msg.Data)
	// 	switch ev := msg.Data.(type) {
	// 	case *slack.HelloEvent:
	// 		fmt.Print("Hello")
	// 	case *slack.ConnectedEvent:
	// 		fmt.Println("Infos:", ev.Info)
	// 		fmt.Println("Connection counter: ", ev.ConnectionCount)
	// 		rtm.SendMessage(rtm.NewOutgoingMessage("Hello World", "#general"))
	// 	case *slack.MessageEvent:
	// 		fmt.Printf("Message: %v\n", ev)
	// 	case *slack.PresenceChangeEvent:
	// 		fmt.Printf("Presence changed: %v", ev)
	// 	case *slack.RTMError:
	// 		fmt.Printf("Error: %s", ev.Error())
	// 	default:
	// 		// fmt.Println("Something went wrong!")
	// 	}
	// }

	// sends message to and api

	// attachment := slack.Attachment{
	// 	Pretext: "Happy Birth day Makuza Mugabo Verite",
	// 	Text:    "Thanks for being cool",
	// 	Color:   "#666666",
	// }

	// channelID, timeStamp, err := api.PostMessage("#general", slack.MsgOptionText("Hello @channel", false), slack.MsgOptionAttachments(attachment), slack.MsgOptionAsUser(true))

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	body, err := ioutil.ReadAll(r.Body)

	// 	if err != nil {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}

	// 	sToken, err := slack.NewSecretsVerifier(r.Header, apiToken)

	// 	if err != nil {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}

	// 	if _, err := sToken.Write(body); err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 	}

	// 	if err := sToken.Ensure(); err != nil {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}

	// 	eventsAPI, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())

	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}

	// 	if eventsAPI.Type == slackevents.URLVerification {
	// 		var r *slackevents.ChallengeResponse
	// 		err := json.Unmarshal([]byte(body), &r)

	// 		if err != nil {
	// 			w.WriteHeader(http.StatusInternalServerError)
	// 			return
	// 		}

	// 		w.Header().Set("Content-Type", "text")
	// 		w.Write([]byte(r.Challenge))
	// 	}

	// 	if eventsAPI.Type == slackevents.CallbackEvent {
	// 		innerEvent := eventsAPI.InnerEvent

	// 		switch ev := innerEvent.Data.(type) {
	// 		case *slackevents.AppMentionEvent:
	// 			api.PostMessage(ev.Channel, slack.MsgOptionText("Hello @Makuza Mugabo Verite", false))
	// 		}
	// 	}

	// })
	// fmt.Println("[INFO] Server started listerning on port 8080")
	// http.ListenAndServe(":8080", nil)

	tokenErr := godotenv.Load()

	if tokenErr != nil {
		log.Fatal("Error while loading .env file")
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"))

	user, err := api.GetUserByEmail("mugaboverite@gmail.com")

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("ID: %s,Fullname: %s,Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}

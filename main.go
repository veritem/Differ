package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/makuzaverite/bd-reminder-bot/lib"
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
}

func main() {

	http.HandleFunc("/slack/events", lib.HandleEvents)

	fmt.Println("[INFO] Server started listerning on port 3000")
	http.ListenAndServe(":3000", nil)
}

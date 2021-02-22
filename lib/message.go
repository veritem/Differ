package lib

// tokenErr := godotenv.Load()

// 	if tokenErr != nil {
// 		log.Fatal("Error while loading .env file")
// 	}

// 	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionLog(log.New(os.Stdout, "slack-bot:", log.Lshortfile|log.LstdFlags)))

// 	// handles message scheduling

// 	for _, element := range schedulesMessages {

// 		_, _, err := api.ScheduleMessage(element.channelID, fmt.Sprint(element.postAt), slack.MsgOptionText(element.text, false))

// 		if err != nil {
// 			fmt.Println("Scheduling")
// 		}

// 	}

// 	for _, element := range schedulesMessages {
// 		scheduledMsg, _, err := api.GetScheduledMessages(&slack.GetScheduledMessagesParameters{
// 			Channel: element.channelID,
// 			Limit:   10,
// 		})

// 		fmt.Println("Schedules message", scheduledMsg)

// 		if err != nil {
// 			fmt.Println("Error while scheduling", err)
// 		}
// 	}

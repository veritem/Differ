package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func main() {

	api := slack.New("xoxb-1765374764241-1758951736772-8hRkefHN5ZzVZ8dWXBRS5I2y")

	groups, err := api.GetGroups(true)

	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	fmt.Println(groups)

	fmt.Println("Hello from birthday Bot")
}

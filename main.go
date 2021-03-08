package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/makuzaverite/bd-reminder-bot/lib"
)

func main() {

	tokenErr := godotenv.Load()

	if tokenErr != nil {
		log.Fatal("Error while loading .env file")
	}

	lib.HandleScheduled()
	http.HandleFunc("/slack/events", lib.HandleEvents)
	http.HandleFunc("/login", lib.HandleLogin)

	fmt.Println("[INFO] Server started listerning on port 3000")
	http.ListenAndServe(":3000", nil)
}

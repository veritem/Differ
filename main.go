package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/makuzaverite/Differ/lib"
)

func main() {

	tokenErr := godotenv.Load()

	if tokenErr != nil {
		log.Fatal("Error while loading .env file")
	}

	// lib.HandleScheduled()
	http.HandleFunc("/slack/events", lib.HandleEvents)
	http.HandleFunc("/login", lib.HandleLogin)
	http.HandleFunc("/install", lib.HandleInstall)

	fmt.Println("[INFO] Server started listerning on port 3000")
	http.ListenAndServe(":3000", nil)
}

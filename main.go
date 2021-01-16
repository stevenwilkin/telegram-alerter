package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/stevenwilkin/treasury/telegram"
)

var (
	port    int
	alerter telegram.Telegram
)

func init() {
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()
}

func alertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("POST /")
		alerter.Notify("Error in logs")
	} else {
		log.Println("GET /")
	}
}

func main() {
	chatId, err := strconv.Atoi(os.Getenv("TELEGRAM_CHAT_ID"))
	if err != nil {
		log.Fatal(err.Error())
	}

	alerter = telegram.Telegram{
		ApiToken: os.Getenv("TELEGRAM_API_TOKEN"),
		ChatId:   chatId}

	log.Printf("Starting on http://0.0.0.0:%d\n", port)
	http.HandleFunc("/", alertHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

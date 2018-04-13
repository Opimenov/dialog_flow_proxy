package main

import (
	"log"
	"net/http"
	. "leo/listener"
	"leo/chat"
)

const ENABLE_CHAT = true
const SHOW_JSON = true

func main() {
	http.HandleFunc("/askleo/", AskLeoHandler)
	if ENABLE_CHAT {
		chat.Start_chat()
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

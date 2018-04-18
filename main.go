package main

import (
	"log"
	"net/http"
	. "leo/listener"
	"leo/chat"
	"flag"
	"fmt"
)

func main() {
	//declare few flags to be used with command line args
	enableChat := flag.Bool("chat", false,"used to enable chat endpoint")
	flag.Parse()
	fmt.Println(*enableChat)
	http.HandleFunc("/askleo/", AskLeoHandler)
	if *enableChat {
		chat.Start_chat()
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

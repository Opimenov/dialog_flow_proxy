//entry point of the dialog flow proxy
package main

import (
	"log"
	"net/http"
	. "leo/listener"
	"leo/chat"
	"flag"
	"fmt"
)

//checks if the flag passed on command line matches "chat".
//If it does enables chat endpoint.
//It is not necessary to have -chat flag in order get a response from Leo agent.
//To get a simple text response from leo agent use:
// <base_url>/askleo/<user_entered_text>
// Note: <user_entered_text> will be sent using GET method.
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

package main

import (
	"fmt"
	"log"
	."leo/models"
	."leo/callers"
	"encoding/json"
	."leo/models/caller_options"
	."leo/models/responses"

	"bufio"
	"os"
)

func main() {
	for {


	//wait for user to enter text query
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("you say : ")
	text, _ := reader.ReadString('\n')

	err, client := NewDialogFlowClient(Options{})
	if err != nil {
		log.Fatal(err)
	}

	//let's start with a simple request
	request := NewDialogFlowRequest(
		client,
		RequestOptions{
			URI:
				client.GetBaseUrl() +
				"query?v=" +
				client.GetApiVersion() +
				"&query="+text+"&lang=" +
				client.GetApiLang() +
				"&sessionId=7413f2c4-2b90-4c43-97a9-b692c6ee2ee5&" +
				"timezone=America/New_York",
			Method: "GET",
			Body:   nil,
		}, )
	//make a request
	data, err := request.Perform()
	//check if error occured
	if err != nil {
		log.Fatal(err)
		fmt.Println("performing request failed")
	}
	var res QueryResponse

	err = json.Unmarshal(data, &res)
	fmt.Println("LEO says " + res.Result.Fulfillment.Speech)
	}
}


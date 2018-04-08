package main

import (
	"fmt"
	"log"
	."leo/callers"
	"encoding/json"
	."leo/models/caller_options"
	."leo/models/responses"

//	"bufio"
//	"os"
	"net/http"
)

func askLeoHandler(w http.ResponseWriter, r *http.Request)  {
	user_text := r.URL.Path[len("/askleo/"):]
	fmt.Fprintf(w,askAgent(user_text))
}



func askAgent(text string) (answer string) {
	err, client := NewDialogFlowClient(AgentClientOptions{})
	if err != nil {
		log.Fatal(err)
	}

	//let's start with a simple request
	request := NewDialogFlowRequest(
		client,
		RequestOptions{
			URI:
			client.GetAgentBaseUrl() +
				"query?v=" +
				client.GetAgentApiVersion() +
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
	var res AgentQueryResponse
	err = json.Unmarshal(data, &res)
	return res.Result.Fulfillment.Speech
}

func main() {
	http.HandleFunc("/askleo/", askLeoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


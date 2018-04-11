package main

import (
	"fmt"
	"log"
	. "leo/callers"
	"encoding/json"
	. "leo/models/caller_options"
	. "leo/models/responses"
	"net/http"
)

const CREATE_PROJECT string = "project.creating"//agent defined value

func askLeoHandler(w http.ResponseWriter, r *http.Request) {
	//get what user said from the query. Works only with GET keyword
	user_text := r.URL.Path[len("/askleo/"):]
	//response data contains json data
	response := askAgent(user_text)

	//if action in result field equals to project.creating
	//and parameters contains { Project : "project name" }
	//call engineering.com api to create a project with the given name
	if CREATE_PROJECT == response.Result.Action &&
		response.Result.Parameters["Project"] != "" {
			//here is an explanation of the following line
			// https://golang.org/ref/spec#Type_assertions
			proj_name := response.Result.Parameters["Project"].(string)
		createProject(proj_name)
	}

	//this is what needs to send to the user
	fmt.Fprintf(w, response.Result.Fulfillment.Speech)
}

func createProject(proj_name string) (answer EngineeringQueryResponse) {
	fmt.Println("creating project with name " + proj_name)
	return
}

func askAgent(text string) (answer AgentQueryResponse) {
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
				"&query=" + text + "&lang=" +
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
	return res
}

func main() {
	http.HandleFunc("/askleo/", askLeoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

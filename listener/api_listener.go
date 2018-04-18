package listener

import (
	"fmt"
	"log"
	. "leo/callers"
	"encoding/json"
	. "leo/models/caller_options"
	. "leo/models/responses"
	"net/http"
	"github.com/ajanicij/goduckgo/goduckgo"
	"os"
)

//agent defined values
const CREATE_PROJECT string = "project.creating"
const SEARCH_WEB string = "web.search"

func AskLeoHandler(w http.ResponseWriter, r *http.Request) {
	//get what user said from the query. Works only with GET keyword
	user_text := r.URL.Path[len("/askleo/"):]
	//response data contains json data
	response := AskAgent(user_text)

	//if action in result field equals to project.creating
	//and parameters contains { Project : "project name" }
	//call engineering.com api to create a project with the given name
	if CREATE_PROJECT == response.Result.Action &&
		response.Result.Parameters["Project"] != "" {
		//here is an explanation of the following line
		// https://golang.org/ref/spec#Type_assertions
		proj_name := response.Result.Parameters["Project"].(string)
		CreateProject(proj_name)
	} else if SEARCH_WEB == response.Result.Action {
		//if action in result field equals to web.search
		//extract query from parameters { "q" : <what to search for> }
		//and do a web search
		searchFor := response.Result.Parameters["q"].(string)
		webRes := DoWebSearch(searchFor)
		fmt.Fprintf(w, webRes)
		 //webRes
	}

	//this is what needs to send to the user
	fmt.Fprintf(w, response.Result.Fulfillment.Speech)
	//return response.Result.Fulfillment.Speech
}

func DoWebSearch(query string) string {
	//q := "https://api.duckduckgo.com/?q="+query+
	//	"&format=json&pretty=1"
	m, err := goduckgo.Query(query)
	fmt.Println("message " + m.AbstractText)
	CheckError(err)
	if len(m.RelatedTopics) != 0 {
		return "Here is what I found on the internet::  " + m.RelatedTopics[0].Text +
			".  Check out this link for more info -> "+
				m.AbstractURL
	}
	return "it seems that wikipedia doesn't know anything about this"
}

func AskAgent(text string) AgentQueryResponse {
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

func CheckError(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(-1)
	}
}

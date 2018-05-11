//Contains an api listener, its job is to process http requests.
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
	"leo/util"
)

const (
	//A string describing an action of creating a project as it is defined by DialogFlow agent LEO.
	CREATE_PROJECT string = "project.creating"
	//A string describing an action of doing a web search as it is defined by DialogFlow agent LEO.
	SEARCH_WEB string = "web.search"
)

//Handles http GET requests received at <askleo> endpoint.
//Parses user url to extract user query.
//Sends a request to DialogFlow Leo Agent.
//Looks for predefined action strings in the Agent's response.
//Either "project.creating" or "web.search".
//Depending on obtained Action object makes necessary API calls.
//Writes a response back based on the Agent's response.
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
	}
	//Send text response to the user
	fmt.Fprintf(w, response.Result.Fulfillment.Speech)
	//return response.Result.Fulfillment.Speech
}

//Makes an api call using duckduckgo search engine api. Easy to use, doesn't
//collect information, never uses third party cookies to target ads.
//To make a request it uses github.com/ajanicij/goduckgo library.
//From the returned result it extracts first related topic and returns it in a text
//form along with the url with full information for further readings.
func DoWebSearch(query string) string {
	//q := "https://api.duckduckgo.com/?q="+query+
	//	"&format=json&pretty=1"
	m, err := goduckgo.Query(query)
	//fmt.Println("here is what you search for :" + query)
	util.CheckError(err)
	if len(m.RelatedTopics) != 0 {
		return "Here is what I found on the internet::  " + m.RelatedTopics[0].Text +
			".  Check out this link for more info -> " +
			m.AbstractURL
	}
	return "it seems that wikipedia doesn't know anything about this"
}

//Makes a GET request to DialogFlow agent using provided text.
//Creates DialogFlowClient object using default client options.
//Populates response model struct using response data.
//NOTE: Uses the same sessionID for every request.
//POSSIBLE IMPROVEMENTS:
// 1.Pass DialogFlowClient object as a method parameter (easier to test).
// 2.Establish a better practice for handling sessionId. Needs to be carefully designed.
// 3.Might want to switch to POST method if privacy is important.
func AskAgent(text string) AgentQueryResponse {
	//TODO pass in NewDialogFlowClient as a method parameter.
	err, client := NewDialogFlowClient(AgentClientOptions{})
	if err != nil {
		log.Fatal(err)
	}

	request := NewDialogFlowRequest(
		client,
		RequestOptions{
			URI:
			client.GetAgentBaseUrl() +
				"query?v=" +
				client.GetAgentApiVersion() +
				"&query=" + text + "&lang=" +
				client.GetApiLang() +
					//TODO improve sessionId generation and handling
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
	//populate response object with obtained data.
	var res AgentQueryResponse
	err = json.Unmarshal(data, &res)
	return res
}


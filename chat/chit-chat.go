//Very basic chat. Only for testing and demo purposes. It has a single window
//for message display and a text entry field along with a send button.
package chat

import (
	"fmt"
	"net/http"
	"strings"
	"leo/listener"
	"leo/callers"
)

type Handler struct {
	wait          chan bool
	message       string
	agentResponse string
}

//Handler function for </chat> endpoint.
//If GET used calls polling method, else broadcast.
func (h *Handler) chatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.polling(w, r)
	case "POST":
		h.broadcast(w, r)
	}
}

//Gets called when </chat> endpoint received a query with a POST method.
func (h *Handler) broadcast(w http.ResponseWriter, r *http.Request) {
	fmt.Println("broadcasting")
	fmt.Println(r.FormValue("body"))
	body := r.FormValue("body")
	h.message = body
	//send true to wait channel so that polling can proceed.
	h.wait <- true
	//and close it
	//please read more here
	// https://golang.org/pkg/builtin/#close
	// http://guzalexander.com/2013/12/06/golang-channels-tutorial.html
	//closing the channel does not block reading operations,
	// After value is read channel will return default value for channel type.
	//in our case value will be false.
	close(h.wait)
	//and make a new empty channel to synchronize broadcast and polling methods.
	h.wait = make(chan bool)
}

func (h *Handler) polling(w http.ResponseWriter, r *http.Request) {
	//Execution will not go forward until something is written in <wait> communication channel.
	//It is happening because <wait> is not buffered and will block if it has no value.
	//Meaning that if there is nothing to read just wait till there is.
	<-h.wait

	//exctract only what user entered
	user_text := strings.TrimSpace(strings.Split(h.message, ":")[1])

	//ask agent and get only the text response
	response := (listener.AskAgent(user_text))

	if "project.creating" == response.Result.Action &&
		response.Result.Parameters["Project"] != "" {
		//check if what we have is a string. Explanation:
		// https://golang.org/ref/spec#Type_assertions
		proj_name := response.Result.Parameters["Project"].(string)
		//call Engineering.com api for creating project
		callers.CreateProject(proj_name)
		//writes a response to the screen
		fmt.Fprintf(w, "<p align=" + "right"+
			"><br><br><b>LEO</b> : "+
			response.Result.Fulfillment.Speech+ "</p>"+
			h.message)

	} else if "web.search" == response.Result.Action {
		//if action in result field equals to web.search
		//extract query from parameters { "q" : <what to search for> }
		//and do a web search
		searchFor := response.Result.Parameters["q"].(string)
		//perform search
		webRes := listener.DoWebSearch(searchFor)
		//writes a response to the screen
		fmt.Fprintf(w, "<p align=" + "right" + "><br><br><b>LEO</b> : "+
			webRes+ "</p>"+ h.message)

	} else {
		fmt.Fprintf(w, "<p align=" + "right"+
			"><br><br><b>LEO</b> : "+
			response.Result.Fulfillment.Speech+ "</p>"+
			h.message)
	}
}

//If -chat option was provided as a command line argument when calling main.go this function
//start chat app at:  http://localhost:8080
func Start_chat() {
	ch := make(chan bool)
	h := Handler{wait: ch}
	http.HandleFunc("/chat", h.chatHandler)
	http.Handle("/", http.FileServer(http.Dir("./chat/static")))
}

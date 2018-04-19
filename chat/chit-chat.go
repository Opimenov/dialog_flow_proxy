package chat

import (
	"fmt"
	"net/http"
	"strings"
	"leo/listener"
	"leo/callers"
)

type Handler struct {
	wait           chan bool
	message        string
	agent_response string
}

func (h *Handler) chatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.polling(w, r)
	case "POST":
		h.broadcast(w, r)
	}
}

func (h *Handler) broadcast(w http.ResponseWriter, r *http.Request) {
	// fires after user pressed "Send"
	fmt.Println("broadcasting")
	fmt.Println(r.FormValue("body"))
	body := r.FormValue("body")
	h.message = body

	h.wait <- true
	close(h.wait)
	h.wait = make(chan bool)
}

func (h *Handler) polling(w http.ResponseWriter, r *http.Request) {
	<-h.wait

	//exctract only what user entered
	user_text := strings.TrimSpace(strings.Split(h.message, ":")[1])
	//fmt.Println(user_text)
	//ask agent and get only the text response
	response := (listener.AskAgent(user_text))

	if "project.creating" == response.Result.Action &&
		response.Result.Parameters["Project"] != "" {
		//here is an explanation of the following line
		// https://golang.org/ref/spec#Type_assertions
		proj_name := response.Result.Parameters["Project"].(string)
		callers.CreateProject(proj_name)
		//writes a response to the screen
		fmt.Fprintf(w, h.message+
			"<p align="+ "right"+
				"><br><br><b>LEO</b> : "+
					response.Result.Fulfillment.Speech+ "</p>")


	} else if "web.search" == response.Result.Action {
		//if action in result field equals to web.search
		//extract query from parameters { "q" : <what to search for> }
		//and do a web search
		searchFor := response.Result.Parameters["q"].(string)
		//fmt.Println(searchFor)
		webRes := listener.DoWebSearch(searchFor)
		fmt.Fprintf(w, h.message+
			"<p align="+ "right"+ "><br><br><b>LEO</b> : "+
				webRes+ "</p>")
		//writes a response to the screen

	} else {
		fmt.Fprintf(w, h.message+
			"<p align="+ "right"+
			"><br><br><b>LEO</b> : "+
			response.Result.Fulfillment.Speech+ "</p>")
	}

}

func Start_chat() {
	ch := make(chan bool)
	h := Handler{wait: ch}
	http.HandleFunc("/chat", h.chatHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//http.ListenAndServe(":9000", nil)
}

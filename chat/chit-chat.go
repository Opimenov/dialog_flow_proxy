package chat

import (
"fmt"
"net/http"
	"strings"
	"leo/listener"
)

type handler struct {
	wait    chan bool
	message string
	agent_response string
}

func (h *handler) chatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.polling(w, r)
	case "POST":
		h.broadcast(w, r)
	}
}

func (h *handler) broadcast(w http.ResponseWriter, r *http.Request) {
	// fires after user pressed "Send"
	fmt.Println("broadcasting")
	fmt.Println(r.FormValue("body"))
	body := r.FormValue("body")
	h.message = body



	h.wait <- true
	close(h.wait)
	h.wait = make(chan bool)
}

func (h *handler) polling(w http.ResponseWriter, r *http.Request) {
	<-h.wait

	//exctract only what user entered
	user_text := strings.TrimSpace(strings.Split(h.message, ":")[1])
	//fmt.Println(user_text)
	//ask agent and get only the text response
	response := (listener.AskAgent(user_text)).Result.Fulfillment.Speech
	//fmt.Println(response)
	h.agent_response = response


	//fmt.Println("polling")
	fmt.Fprintf(w, h.message+
	"<p align="+"right"+"><br><br><b>LEO</b> : "+h.agent_response+"</p>") //writes a response to the screen
	//fmt.Fprintf(w, "<b>LEO</b> : "+h.agent_response)

}



func Start_chat() {
	ch := make(chan bool)
	h := handler{wait: ch}
	http.HandleFunc("/chat", h.chatHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//http.ListenAndServe(":9000", nil)
}
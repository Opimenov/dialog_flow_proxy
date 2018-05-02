//Execution starting point.
// Contains a main method that starts web service and
// assigns handlers to different endpoints.
package main

import (
	"log"
	"net/http"
	. "leo/listener"
	"leo/chat"
	"flag"
	"fmt"
)

//Gets called when >>> go run  <path_to_go>/go/src/leo/main.go is ran.
// To enable demo chat at <base_url>:8080 use flag: -chat
// Example >>> go run  <path_to_go>/go/src/leo/main.go -chat.
// If you run this on your local machine, go to:
// localhost:8080 to test.
// To enable <askleo> endpoint use flag: -leo
// Example >>> go run  <path_to_go>/go/src/leo/main.go -leo.
// After service initialization, to get a simple text response from leo agent use:
// http://<base_url>:8080/askleo/<text_to_be_processed>
// Note: <user_entered_text> will be sent using GET method.
func main() {
	//declare few flags to be used with command line args
	enableChat := flag.Bool("chat", false, "used to enable chat endpoint")
	enableLeoEndPoint := flag.Bool("leo", false, "used to enable leo agent query endpoint")
	flag.Parse()
	fmt.Println("enabling chat? --> ", *enableChat)
	if *enableChat {
		chat.Start_chat()
	}
	fmt.Println("enabling askleo endpoint? --> ", *enableLeoEndPoint)
	if *enableLeoEndPoint {
		http.HandleFunc("/askleo/", AskLeoHandler)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

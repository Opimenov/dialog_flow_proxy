package main

import (
	"log"
	"net/http"
	. "leo/listener"
)




func main() {
	http.HandleFunc("/askleo/", AskLeoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

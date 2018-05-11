//File: api_listener_test.go. The idea is to hold all tests in this
//package. In this way we can generate documentation and have tests
//reside in one place.
package test_suite

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"leo/listener"
)

//Test makes a request to perform a web search and checks the response text.
func TestAskLeoHandlerSearchWeb(t *testing.T) {
	//since there is a context feature we need to do a few dummy
	//requests to reset the context
	req, err := http.NewRequest("GET",
		"/askleo/search%20rat",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization",
		"Bearer a7cd00ab945249428a7f2f5841213fb3")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(listener.AskLeoHandler)

	handler.ServeHTTP(responseRecorder, req)

	for i := 0; i < 10; i++ {
		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(listener.AskLeoHandler)
		handler.ServeHTTP(responseRecorder, req)
	}

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: "+
			"got %v want %v", status, http.StatusOK)
	}

	if "Here is what I found on the internet::  " +
		"Rat Rats are various medium-sized, " +
		"long-tailed rodents of the superfamily Muroidea..  " +
		"Check out this link for more info -> " +
		"https://en.wikipedia.org/wiki/Rat_(disambiguation)" !=
		responseRecorder.Body.String() {
		t.Errorf("handler returned unexpected body: "+
			"got %v want %v", responseRecorder.Body.String(),
			"...")
	}
}

//Test makes an http request using /askleo endpoint and simply checks if a
//response status code is 200
func TestAskLeoHandlerStatusCodeCheck(t *testing.T) {
	req, err := http.NewRequest("GET",
		"/askleo/what%20is%20going%20on",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization",
		"Bearer a7cd00ab945249428a7f2f5841213fb3")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(listener.AskLeoHandler)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: "+
			"got %v want %v", status, http.StatusOK)
	}
}

//Test makes an http request asking an agent to create a project,
//and checks if any of the follow up intent questions has been returned
func TestAskLeoHandlerCreatingProjectAction(t *testing.T) {
	req, err := http.NewRequest("GET",
		"/askleo/can%20you%20create%20project",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization",
		"Bearer a7cd00ab945249428a7f2f5841213fb3")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(listener.AskLeoHandler)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: "+
			"got %v want %v", status, http.StatusOK)
	}

	if "What is the name of the project?" != responseRecorder.Body.String() &&
		"How should I name your project?" != responseRecorder.Body.String() &&
		"How should I name your new project?" != responseRecorder.Body.String() {
		t.Errorf("handler returned unexpected body: "+
			"got %v want %v", responseRecorder.Body.String(),
			"What is the name of the project?")
	}
}

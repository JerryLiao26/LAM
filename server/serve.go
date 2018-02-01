package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type message struct {
	Text  string `json:"text"`
	Token string `json:"token"`
}

type serverRespond struct {
	Code   int    `json:"code"`
	Text   string `json:"text"`
	Method string `json:"method"`
}

type openSendRespond struct {
	Code   int    `json:"code"`
	Text   string `json:"text"`
	Sender string `json:"sender"`
}

type trustedSendRespond struct {
	Code   int    `json:"code"`
	Tag    string `json:"tag"`
	Text   string `json:"text"`
	Sender string `json:"sender"`
}

func enterLog(w http.ResponseWriter, r *http.Request) {
	serverLogger("From", r.RemoteAddr, INFO)
	serverLogger(r.Method, r.URL.Path, INFO)
	serverLogger("Scheme", r.URL.Scheme, INFO)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	enterLog(w, r)
	sr := serverRespond{
		Code:   200,
		Method: r.Method,
		Text:   "LAM Server here",
	}
	output, err := json.Marshal(sr)
	if err != nil {
		serverLogger("JSON build error", err.Error(), ERROR)
	}
	fmt.Fprintf(w, string(output))
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	enterLog(w, r)
	if r.Method == "POST" {
		var m message
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&m)
		if err != nil {
			serverLogger("JSON parse error", err.Error(), ERROR)
		}
		// TODO: checking token

		// TODO: storing database

		//  Response to client
		user := "jerryliao" // Test use only
		sr := serverRespond{
			Code:   200,
			Method: r.Method,
			Text:   "Message from " + user + " saved",
		}
		output, err := json.Marshal(sr)
		if err != nil {
			serverLogger("JSON build error", err.Error(), ERROR)
		}
		fmt.Fprintf(w, string(output))
		serverLogger("Message from", user, INFO)
	} else {
		sr := serverRespond{
			Code:   400,
			Method: r.Method,
			Text:   "Method not allowed",
		}
		output, err := json.Marshal(sr)
		if err != nil {
			serverLogger("JSON build error", err.Error(), ERROR)
		}
		fmt.Fprintf(w, string(output))
		serverLogger("Send warning", r.Method+" method is not allowed for /send, abandoned", WARN)
	}
}

func main() {
	// Handlers
	http.HandleFunc("/send", sendHandler)
	http.HandleFunc("/hello", helloHandler)
	// Server start
	serverLogger("Starting", "Serve at http://127.0.0.1:12306", INFO)
	err := http.ListenAndServe("127.0.0.1:12306", nil)
	// Error
	if err != nil {
		serverLogger("Cannot start", err.Error(), ERROR)
	}
}

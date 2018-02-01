package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func serverLogger(title string, content string, level string) {
	levelString := "[" + level + "]"
	timeString := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	fmt.Println(levelString, timeString, title, content)
}

func helloPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	serverLogger("Path:", r.URL.Path, "Log")
	serverLogger("Scheme:", r.URL.Scheme, "Log")

	fmt.Fprintf(w, "LAM Server Here")
}

func main() {
	http.HandleFunc("/hello", helloPage)
	err := http.ListenAndServe(":12580", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

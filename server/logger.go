package main

import (
	"fmt"
	"time"
)

func serverLogger(title string, content string, level LogLevel) {
	titleString := title + ":"
	levelString := "[" + level + "]"
	timeString := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	fmt.Println(timeString, levelString, titleString, content)
}

func cliLogger(content string) {
	logString := "[LAM]" + content
	fmt.Println(logString)
}

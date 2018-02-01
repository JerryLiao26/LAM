package main

import (
	"fmt"
	"time"
)

// LogLevel is to defined logging level
type LogLevel string

// INFO level
var INFO LogLevel = "INFO"

// WARN level
var WARN LogLevel = "WARN"

// ERROR level
var ERROR LogLevel = "ERROR"

func serverLogger(title string, content string, level LogLevel) {
	titleString := title + ":"
	levelString := "[" + level + "]"
	timeString := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	fmt.Println(timeString, levelString, titleString, content)
}

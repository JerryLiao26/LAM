package main

// Command is to defined supported command
type Command string

// SERVE command
var SERVE Command = "serve"

// START command
var START Command = "start"

// GEN command
var GEN Command = "gen"

// SET command
var SET Command = "set"

// HELP command
var HELP Command = "help"

// Supported commands
var supportedCommand = [...]Command{SERVE, START, GEN, SET, HELP}

// LogLevel is to defined logging level
type LogLevel string

// INFO level
var INFO LogLevel = "INFO"

// WARN level
var WARN LogLevel = "WARN"

// ERROR level
var ERROR LogLevel = "ERROR"

// Config path
type confPath struct {
	homepath string
	confFile string
	confDir  string
}

// Datbase config
type dbConf struct {
	username string
	password string
	addr     string
	port     string
}

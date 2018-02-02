package main

// Command is to defined supported command
type Command string

// SERVE command
const SERVE Command = "serve"

// START command
const START Command = "start"

// GEN command
const GEN Command = "gen"

// SET command
const SET Command = "set"

// HELP command
const HELP Command = "help"

// Supported commands
var supportedCommand = [...]Command{SERVE, START, GEN, SET, HELP}

// LogLevel is to defined logging level
type LogLevel string

// INFO level
const INFO LogLevel = "INFO"

// WARN level
const WARN LogLevel = "WARN"

// ERROR level
const ERROR LogLevel = "ERROR"

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

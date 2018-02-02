package main

import (
	"fmt"
	"os"
)

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

func str2comm(str string) Command {
	return Command(str)
}

func serveHandler() {

}

func startHandler() {
	if len(os.Args) >= 3 {
		str := os.Args[2]
		serve(str)
	} else if len(os.Args) >= 2 {
		serve("127.0.0.1:12306")
	}
}

func genHandler() {

}

func setHandler() {

}

func helpHandler() {
	fmt.Println("Usage:")
	fmt.Println("lam-cli [command] [args...]")
	fmt.Println("Commands:")
	// Print help text
	fmt.Println("help                       Print help text")
	fmt.Println("start [address]:[port]     Start server")
	fmt.Println("serve [address]:[port]     Start server deamon")
	fmt.Println("gen [tag]                  Generate token for provided tag")
	fmt.Println("set [username]:[password]  Set database connection info")
}

func main() {
	if len(os.Args) >= 2 {
		// Get command
		comm := str2comm(os.Args[1])
		// Handlers for commands
		switch comm {
		case SERVE:
			serveHandler()
		case START:
			startHandler()
		case GEN:
			genHandler()
		case SET:
			setHandler()
		default:
			helpHandler()
		}
	} else {
		helpHandler()
	}
}

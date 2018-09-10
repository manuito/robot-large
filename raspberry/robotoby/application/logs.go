package application

import (
	"log"
)

var debug = State.Config.LogLevel == "DEBUG"
var info = State.Config.LogLevel == "INFO"

// Debug : simple debug entrypoint
func Debug(v string) {
	if debug {
		log.Println(v)
	}
}

// Info : simple debug entrypoint
func Info(v string) {
	if debug || info {
		log.Println(v)
	}
}

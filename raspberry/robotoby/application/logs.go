package application

import (
	"log"
)

// Debug : simple debug entrypoint
func Debug(v string) {
	if State.Config.Debug {
		log.Println(v)
	}
}

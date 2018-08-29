package application

import (
	"net"
	"robotoby/tools"
)

/*
 * Global states / configuration access for the current pilote application
 */

// State : Global application environment
var State = initState()

// ApplicationState : global state model for application
type ApplicationState struct {
	Config     Configuration
	OutBountIP net.IP
}

// Get preferred outbound ip of this machine
func initState() ApplicationState {
	return ApplicationState{loadConfiguration(), tools.GetOutboundIP()}
}

package main

import (
	"fmt"
	"robotoby/connect"
	"robotoby/pilot"
)

func main() {
	fmt.Println("Starting server")
	connect.StartGetMonitoring()
	pilot.StartStateUpdate()
	pilot.StartServer()
}

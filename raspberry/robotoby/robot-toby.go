package main

import (
	"fmt"
	"robotoby/connect"
	"robotoby/pilote"
)

func main() {
	fmt.Println("Starting server")
	connect.StartGetMonitoring()
	pilote.StartStateUpdate()
	pilote.StartServer()
}

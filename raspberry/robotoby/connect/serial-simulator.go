package connect

import (
	"encoding/json"
	"fmt"
	"os"
	"robotoby/application"
)

/*
 * Some test values generator for simulated robot connect
 */

// See RobotCommand comment for details
var simuCommands = loadSource()

type simulation struct {
	Serial []string
}

var current = 0

func simuSendSerial(value string) {
	application.Debug("Sending : " + value)
}

func simuReceiveSerial() string {

	if current >= len(simuCommands) {
		current = 0
	}
	val := simuCommands[current]
	current++
	return val
}

func loadSource() []string {
	file, _ := os.Open(application.State.Config.SerialSimulation)
	defer file.Close()
	decoder := json.NewDecoder(file)
	simulation := simulation{}
	err := decoder.Decode(&simulation)
	if err != nil {
		fmt.Println("error:", err)
	}
	return simulation.Serial
}

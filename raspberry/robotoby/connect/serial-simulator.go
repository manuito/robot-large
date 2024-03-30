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
var simuCommands []string

type simulation struct {
	Serial []string
}

var current = 0

func startSimulation() {
	application.Info("Start with simulation")
	simuCommands = loadSource()
}

func simuSendSerial(value string) {
	application.Debug(fmt.Sprintf("Sending simulated Serial CMD : %v", value))
}

func simuReceiveSerial() string {

	if simuCommands == nil {
		startSimulation()
	}

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
		fmt.Println("Simulation load error:", err)
	}
	return simulation.Serial
}

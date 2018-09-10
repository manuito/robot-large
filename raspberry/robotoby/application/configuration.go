package application

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
 * Pilote app configuration management.
 * Uses file "conf.json" for configuration definition
 */

// Configuration : Robot configuration holder
type Configuration struct {
	RobotName            string
	DebounceBufferSize   int
	DebounceMaxVariation int
	StandardMove         int
	StandardSpeed        int
	StandardDegree       int
	CentimerToSpeedRatio int
	DegreeToSpeedRatio   int
	SpeedChangeSteps     int
	LogLevel             string
	MinimalStopRange     int
	MonitoringDelayMs    int
	SerialSimulation     string
	SerialDevice         string
	SerialSpeed          int
}

func loadConfiguration() Configuration {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

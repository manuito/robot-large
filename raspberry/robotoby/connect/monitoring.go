package connect

import (
	"log"
	"time"
)

/*
 * Features for using raw monitoring values from robot
 * Monitoring is a continuous feedback from the Mega MCU. It is published
 * unchanged in a channel "Monitors" for conveniance
 */

// Monitors : all the monitoring elements in a channel
var Monitors = make(chan RobotMonitoring)
var stop = false

// StartGetMonitoring :
func StartGetMonitoring() {
	log.Println("Starting auto read of monitoring")
	go func() {
		for !stop {
			time.Sleep(100 * time.Millisecond)
			Monitors <- GetMonitoring()
		}
	}()
}

// StopGetMonitoring :
func StopGetMonitoring() {
	log.Println("Stopping monitoring auto read")
	stop = true
}

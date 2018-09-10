package connect

import (
	"robotoby/application"
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
var delay = time.Duration(application.State.Config.MonitoringDelayMs)

// StartGetMonitoring :
func StartGetMonitoring() {
	application.Info("Starting auto read of monitoring")
	go func() {
		for !stop {
			time.Sleep(delay * time.Millisecond)
			mon := GetMonitoring()
			if mon != nil {
				Monitors <- *mon
			}
		}
	}()
}

// StopGetMonitoring :
func StopGetMonitoring() {
	application.Info("Stopping monitoring auto read")
	stop = true
}

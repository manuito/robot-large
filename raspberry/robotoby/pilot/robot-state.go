package pilot

import (
	"fmt"
	"robotoby/application"
	"robotoby/connect"
)

/*
 * Here the access to clean debounced robot state (built from raw monitoring)
 */

var positionLeftDeb, positionRightDeb, temperatureDeb debouncedValue
var leftSpeed, rightSpeed, bandState, faceState, beeperState int

// RobotState : global state model with usefull values
type RobotState struct {
	LeftDistance  int
	RightDistance int
	LeftSpeed     int
	RightSpeed    int
	Temperature   int
	Action        RobotAction
}

// RobotAction : current action processing
type RobotAction struct {
	BandState   int
	FaceState   int
	BeeperState int
}

// StartStateUpdate : Start to get monitoring and clean positions
func StartStateUpdate() {
	go func() {
		for {
			mon := <-connect.Monitors
			positionLeftDeb.storeIfClean(mon.LeftDistance)
			positionRightDeb.storeIfClean(mon.RightDistance)
			temperatureDeb.storeIfClean(mon.Temperature)
			leftSpeed = mon.LeftMotorState
			rightSpeed = mon.RightMotorState
			bandState = mon.BandState
			faceState = mon.FaceState
			beeperState = mon.BeeperState
			application.Info(fmt.Sprintf("Get state : Left = %v, Right = %v, Temperature = %v", positionLeftDeb.getCurrent(), positionRightDeb.getCurrent(), temperatureDeb.getCurrent()))
		}
	}()
	application.Info("Monitoring linked to state update")
}

// GetCurrentRobotState : get current clean position, left / right
func GetCurrentRobotState() RobotState {

	return RobotState{
		positionLeftDeb.getCurrent(),
		positionRightDeb.getCurrent(),
		leftSpeed,
		rightSpeed,
		temperatureDeb.getCurrent(),
		RobotAction{
			bandState,
			faceState,
			beeperState,
		},
	}
}

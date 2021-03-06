package pilote

import (
	"robotoby/application"
	"robotoby/connect"
	"strconv"
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
			application.Info("Get state : Left = " + strconv.Itoa(positionLeftDeb.getCurrent()) + ", Right = " + strconv.Itoa(positionRightDeb.getCurrent()) + ", Temperature = " + strconv.Itoa(temperatureDeb.getCurrent()))
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

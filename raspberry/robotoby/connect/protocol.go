package connect

/*
 * Here all details on connect protocol (stack used for serial exchange with robot mega MCU) :
 * -> Monitoring protocol (continuous feedback from sensors / MCU core)
 * -> Command protocol (serial command from PI to Mega for driving)
 */

import (
	"fmt"
	"strconv"
	"strings"
)

// RobotCommand : Definition of a command
//
//  Command syntax :
//
//  M/m => Enable monitoring state
//
//  B/b => beeper
//    B1 => beep
//    B2 => mario theme
//    B3 => mario underworld
//
//  S/s => Straight
//    Sxy => move speed
//
//  L/l => Left
//    Lxy => turn speed
//
//  R/r => Right
//    Rxy => turn speed
//
//  H/h => Head
//    H1 => face 1
//    H2 => face 2
//
//  F/f => Band light
//    F1 => Police
//    F2 => K2000
//    F3 => Dim yellow
//    F4 => Car light
type RobotCommand struct {
	Action rune
	Dim    int
}

// RobotMonitoring : Current monitoring raw data from the robot
// Serial monitoring values are in form "ABxy|CDxy|EFxy ..."
// Fixed order for attributes :
// 0]  => MOx  => Monitoring state (always on = 1)
// 1]  => MLxy => Left motor state
// 2]  => MRxy => Right motor state
// 3]  => DLxy => Left distance value
// 4]  => DRxy => Right distance value
// 5]  => AXxy => Accelerometer X
// 6]  => AYxy => Accelerometer Y
// 7]  => AZxy => Accelerometer Z
// 8]  => GXxy => Gyroscop X
// 9]  => GYxy => Gyroscop Y
// 10] => GZxy => Gyroscop Z
// 11] => TPxy => Temperature
// 12] => BAxy => RGB Band state
// 13] => FAxy => Face state
// 14] => BExy => Beeper state
type RobotMonitoring struct {
	MonitoringState int
	LeftMotorState  int
	RightMotorState int
	LeftDistance    int
	RightDistance   int
	AccelX          int
	AccelY          int
	AccelZ          int
	GyroX           int
	GyroY           int
	GyroZ           int
	Temperature     int
	BandState       int
	FaceState       int
	BeeperState     int
}

// Internal convert for a command
func (c *RobotCommand) convertSerial() string {
	return string(c.Action) + strconv.Itoa(c.Dim)
}

// Send : sending the command with serial connect
func (c *RobotCommand) Send() error {
	return sendSerial(c.convertSerial())
}

// Converter for one monitoring value in raw serial values
func monitoringValue(pos int, codes []string) int {

	v, err := strconv.Atoi((codes[pos])[2:])

	if err != nil {
		fmt.Print(err)
	}

	return int(v)
}

func readMonitoring(rawValue string) RobotMonitoring {
	codes := strings.Split(rawValue, "|")

	return RobotMonitoring{
		MonitoringState: monitoringValue(0, codes),
		LeftMotorState:  monitoringValue(1, codes),
		RightMotorState: monitoringValue(2, codes),
		LeftDistance:    monitoringValue(3, codes),
		RightDistance:   monitoringValue(4, codes),
		AccelX:          monitoringValue(5, codes),
		AccelY:          monitoringValue(6, codes),
		AccelZ:          monitoringValue(7, codes),
		GyroX:           monitoringValue(8, codes),
		GyroY:           monitoringValue(9, codes),
		GyroZ:           monitoringValue(10, codes),
		Temperature:     monitoringValue(11, codes),
		BandState:       monitoringValue(12, codes),
		FaceState:       monitoringValue(13, codes),
		BeeperState:     monitoringValue(14, codes)}
}

// GetMonitoring : get the monitoring value available in serial RX buffer
func GetMonitoring() RobotMonitoring {

	v, err := receiveSerial()

	if err != nil {
		fmt.Print(err)
	}

	// log.Println("Received monitoring : " + v)

	return readMonitoring(v)
}

package pilote

import (
	"log"
	"robotoby/application"
	"robotoby/connect"
	"strconv"
	"strings"
	"time"
)

/*
 * Command model for pilote. Mostly process move change,
 * going from a low level speed-based actions to high level distance / degree actions
 */

var supportedOptions = []string{
	"AUTO_AVOIDANCE",
}

// Command : Model for a prepared command for pilote / auto-pilote
type Command struct {
	Action     string
	Dimensions []int
}

// ChainedCommand : For multiple commands to process
type ChainedCommand struct {
	Commands []Command
	Options  []string
}

/*
 * Actions :
 *
 * ####### Move actions ########
 *
 *  - scoped move -
 *  MCM - no dim => Move standard Centimeters (1m) at standard speed
 *  MCM - dim0 => Move dim0 Centimeter at standard speed
 *  MCM - dim0 - dim1 => Move dim0 Centimeter at dim1 speed
 *  MCM - dim0 - dim1 - dim2 => Move dim0 Centimeter, starting at dim1 speed, finishing at dim2 speed
 *
 *  - scoped turn left -
 *  LDE => Turn left for standard degrees at standard speed
 *  LDE - dim0 => Turn left for dim0 degrees at standard speed
 *  LDE - dim0 - dim1 => Turn left for dim0 degrees as dim1 speed
 *
 *  - scoped turn right -
 *  RDE => Turn right for dim0 degrees at standard speed
 *  RDE - dim0 => Turn right for dim0 degrees at standard speed
 *  RDE - dim0 => Turn right for dim0 degrees as dim1 speed
 *
 *  - unscoped move -
 *  MOV - no dim => Move continuously at standard speed
 *  MOV - dim0 => Move continuously at dim0 speed
 *
 *  - unscoped turn left -
 *  LEF => Turn left continuously at standard speed
 *  LEF - dim0 => Turn left continuously at dim0 speed
 *
 *  - unscoped turn right -
 *  RIG => Turn right continuously at standard speed
 *  RIG - dim0 => Turn right continuously at dim0 speed
 */

// Process : On specified command, run (in go routine)
func (com Command) Process() error {
	go com.getProcess()()
	return nil
}

func (com Command) getProcess() func() {

	log.Println("Start a command for action " + com.Action)

	switch com.Action {
	case "MCM":
		return scopedMoveRun(com.Dimensions)
	case "MOV":
		return unscopedMoveRun(com.Dimensions)
	case "LDE":
		return scopedTurnRun('L', com.Dimensions)
	case "LEF":
		return unscopedTurnRun('L', com.Dimensions)
	case "RDE":
		return scopedTurnRun('R', com.Dimensions)
	case "RIG":
		return unscopedTurnRun('R', com.Dimensions)
	default:
		log.Println("Cannot process command " + com.Action)
	}
	return func() {
		// Empty
	}
}

// Process : run all the chained command, in one single routine
func (chain ChainedCommand) Process() error {
	go chain.getProcess()()
	return nil
}

func (chain ChainedCommand) getProcess() func() {
	return func() {
		log.Println("Start a command chain with options " + strings.Join(chain.Options, ", "))
		for pos, com := range chain.Commands {
			log.Println(" --- Command " + strconv.Itoa(pos) + " ---")
			com.getProcess()()
		}
	}
}

/*
 * ############################################
 * ######           ACTION RUNS          ######
 * ############################################
 */

// MCM
func scopedMoveRun(dim []int) func() {

	switch len(dim) {
	case 0:
		// Standard speed / standard distance
		return prepareScopedMoveConstantSpeedRun(application.State.Config.StandardSpeed, application.State.Config.StandardMove)
	case 1:
		// Standard speed / given distance
		return prepareScopedMoveConstantSpeedRun(application.State.Config.StandardSpeed, dim[0])
	case 2:
		// Given speed / given distance
		return prepareScopedMoveConstantSpeedRun(dim[1], dim[0])
	case 3:
		// changing speed / given distance
		return prepareScopedMoveVariableSpeedStepsRun(dim[1], dim[2], dim[0])
	}

	return func() {
		log.Println("Wrong dim for scoped move command")
	}
}

// LDE / RDE
func scopedTurnRun(action rune, dim []int) func() {

	switch len(dim) {
	case 0:
		// Standard speed / standard degree
		return prepareScopedTurnConstantSpeedRun(action, application.State.Config.StandardSpeed, application.State.Config.StandardDegree)
	case 1:
		// Standard speed / given degree
		return prepareScopedTurnConstantSpeedRun(action, application.State.Config.StandardSpeed, dim[0])
	case 2:
		// Given speed / given distance
		return prepareScopedTurnConstantSpeedRun(action, dim[1], dim[0])
	}

	return func() {
		log.Println("Wrong dim for scoped turn command")
	}
}

// MOV
func unscopedMoveRun(dim []int) func() {

	switch len(dim) {
	case 0:
		// Standard speed / standard distance
		return prepareUnScopedMoveConstantSpeedRun(application.State.Config.StandardSpeed)
	case 1:
		// Standard speed / given distance
		return prepareUnScopedMoveConstantSpeedRun(dim[0])
	}

	return func() {
		log.Println("Wrong dim for unscoped move command")
	}
}

// LEF / RIG
func unscopedTurnRun(action rune, dim []int) func() {

	switch len(dim) {
	case 0:
		// Standard speed
		return prepareUnscopedTurnConstantSpeedRun(action, application.State.Config.StandardSpeed)
	case 1:
		// Given speed
		return prepareUnscopedTurnConstantSpeedRun(action, dim[0])
	}

	return func() {
		log.Println("Wrong dim for scoped turn command")
	}
}

/*
 * ############################################
 * ######           INNER FUNC           ######
 * ############################################
 */

func prepareScopedMoveConstantSpeedRun(speed, centimeter int) func() {

	duration := prepareDurationForDistance(speed, centimeter)

	return func() {
		// Process direct speed
		processMoveCentimeterAction(speed, duration)

		// And stop
		processStopAction()
	}
}

func prepareScopedTurnConstantSpeedRun(action rune, speed, degree int) func() {

	duration := prepareDurationForTurnDegree(speed, degree)

	return func() {
		// Process direct speed
		processTurnDegreeAction(action, speed, duration)

		// And stop
		processStopAction()
	}
}

func prepareUnScopedMoveConstantSpeedRun(speed int) func() {

	return func() {

		log.Println(" => start unscoped move at speed " + strconv.Itoa(speed))

		// Process direct speed
		command := connect.RobotCommand{Action: 'S', Dim: speed}
		command.Send()
	}
}

func prepareUnscopedTurnConstantSpeedRun(action rune, speed int) func() {

	return func() {

		log.Println(" => start unscoped turn " + strconv.QuoteRune(action) + " at speed " + strconv.Itoa(speed))

		// Process direct speed
		command := connect.RobotCommand{Action: action, Dim: speed}
		command.Send()
	}
}

func prepareScopedMoveVariableSpeedStepsRun(startSpeed, endSpeed, centimeter int) func() {

	steps := application.State.Config.SpeedChangeSteps
	stepSpeedChange := (endSpeed - startSpeed) / (steps - 1)
	stepCentimer := centimeter / steps

	log.Println("Variable speed prepare with " + strconv.Itoa(steps) + " steps, of " + strconv.Itoa(stepCentimer) + " cm, adding " + strconv.Itoa(stepSpeedChange) + " to speed")

	return func() {
		currentSpeed := startSpeed

		log.Println("[Start move steps for variable speed]")

		// Process required steps
		for index := 0; index < steps-1; index++ {
			duration := prepareDurationForDistance(currentSpeed, stepCentimer)
			processMoveCentimeterAction(currentSpeed, duration)
			currentSpeed += stepSpeedChange
		}

		// End stop
		processStopAction()
	}
}

func processStopAction() {

	log.Println(" => stop move")

	command := connect.RobotCommand{Action: 's', Dim: 0}
	command.Send()
}

func processMoveCentimeterAction(speed int, duration time.Duration) {

	log.Println(" => scoped move at speed " + strconv.Itoa(speed) + " and sleep " + duration.String())

	command := connect.RobotCommand{Action: 'S', Dim: speed}
	command.Send()
	if speed > 0 {
		// Auto set duration before next step
		time.Sleep(duration)
	}
}

func processTurnDegreeAction(action rune, speed int, duration time.Duration) {

	log.Println(" => scoped turn " + strconv.QuoteRune(action) + " at speed " + strconv.Itoa(speed) + " and sleep " + duration.String())

	command := connect.RobotCommand{Action: action, Dim: speed}
	command.Send()
	if speed > 0 {
		// Auto set duration before next step
		time.Sleep(duration)
	}
}

/*
 * ############################################
 * ######         DIMENSION CALC         ######
 * ############################################
 */

func prepareDurationForDistance(speed, centimer int) time.Duration {
	return time.Duration((centimer*application.State.Config.CentimerToSpeedRatio)/speed) * time.Millisecond
}

func prepareDurationForTurnDegree(speed, degree int) time.Duration {
	return time.Duration((degree*application.State.Config.DegreeToSpeedRatio)/speed) * time.Millisecond
}

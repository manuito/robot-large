package pilote

import (
	"robotoby/application"
	"time"
)

/*
 * Here code for auto-pilote features
 */

const contMoveAction string = "MOV"
const turnLeftAction string = "LDE"
const turnRightAction string = "RDE"

var interruptAutoPilote = false
var autoPiloteRunning = false

// Check on monitoring
var brainCheckDelay = time.Duration(application.State.Config.MonitoringDelayMs)

func stopAutoPilote() {

	if autoPiloteRunning {
		application.Info("[AUTO]-[Interrupting auto-pilote]")
	}

	interruptAutoPilote = true
}

func startFullAutoPilote() {

	// Do not start auto pilote is already running
	if !autoPiloteRunning {

		application.Info("[AUTO]-[Starting auto-pilote]")

		// Reset interrupt
		interruptAutoPilote = false
		autoPiloteRunning = true

		mov := Command{
			Action: contMoveAction,
		}

		// Start moving
		mov.Process()

		go func() {

			for !interruptAutoPilote {
				state := GetCurrentRobotState()

				application.Debug("[AUTO]-[Check distance]")

				// Basic check on combined space
				if state.LeftDistance+state.RightDistance < application.State.Config.MinimalStopRange || state.LeftDistance+state.LeftDistance < application.State.Config.MinimalStopRange {
					// Change direction by checking left or right space
					if state.LeftDistance+state.RightDistance > state.LeftDistance+state.LeftDistance {
						application.Info("[AUTO]-[Found an obstacle, changing direction - turn right]")
						changeDirection(turnRightAction).Process()

					} else {
						application.Info("[AUTO]-[Found an obstacle, changing direction - turn left]")
						changeDirection(turnLeftAction).Process()
					}
				}

				time.Sleep(brainCheckDelay * time.Millisecond)
			}

			// If interrupted, do not stop : let the new interrupting command continue

			// Just update state
			autoPiloteRunning = false
		}()
	} else {
		application.Info("[AUTO]-[Auto-pilote is already running]")
	}
}

// prepareDirectionChange : in auto-pilote, build a chain of turn + move continuously
func changeDirection(action string) ChainedCommand {

	return ChainedCommand{
		Commands: []Command{
			Command{
				Action:     action,
				Dimensions: []int{}, // No dim = Standard degree, standard speed
			},
			Command{
				Action:     contMoveAction,
				Dimensions: []int{}, // No dim = standard speed
			},
		},
		Options: []string{"FLUID_MOVE"},
	}
}

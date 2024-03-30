package pilot

import (
	"log"
	"robotoby/application"
	"time"
)

/*
 * Here code for auto-pilot features
 */

const contMoveAction string = "MOV"
const turnLeftAction string = "LDE"
const turnRightAction string = "RDE"

var interruptAutoPilot = false
var autoPilotRunning = false

// Check on monitoring
var brainCheckDelay = time.Duration(application.State.Config.MonitoringDelayMs)

func stopAutoPilot() {

	if autoPilotRunning {
		application.Info("[AUTO]-[Interrupting autopilot]")
	}

	interruptAutoPilot = true
}

func startFullAutoPilot() {

	// Do not start autopilot is already running
	if !autoPilotRunning {

		application.Info("[AUTO]-[Starting autopilot]")

		// Reset interrupt
		interruptAutoPilot = false
		autoPilotRunning = true

		mov := Command{
			Action: contMoveAction,
		}

		// Start moving
		err := mov.Process()
		if err != nil {
			log.Printf("Command error : %v\n", err)
			return
		}

		go func() {

			for !interruptAutoPilot {
				state := GetCurrentRobotState()

				application.Debug("[AUTO]-[Check distance]")

				// Basic check on combined space
				if state.LeftDistance+state.RightDistance < application.State.Config.MinimalStopRange || state.LeftDistance+state.LeftDistance < application.State.Config.MinimalStopRange {
					// Change direction by checking left or right space
					if state.LeftDistance+state.RightDistance > state.LeftDistance+state.LeftDistance {
						application.Info("[AUTO]-[Found an obstacle, changing direction - turn right]")
						err := changeDirection(turnRightAction).Process()
						if err != nil {
							log.Printf("Command error : %v\n", err)
							return
						}

					} else {
						application.Info("[AUTO]-[Found an obstacle, changing direction - turn left]")
						err := changeDirection(turnLeftAction).Process()
						if err != nil {
							log.Printf("Command error : %v\n", err)
							return
						}
					}
				}

				time.Sleep(brainCheckDelay * time.Millisecond)
			}

			// If interrupted, do not stop : let the new interrupting command continue

			// Just update state
			autoPilotRunning = false
		}()
	} else {
		application.Info("[AUTO]-[Auto-pilot is already running]")
	}
}

// prepareDirectionChange : in auto-pilot, build a chain of turn + move continuously
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

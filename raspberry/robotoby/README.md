# Pilot application

The pilot application is a GO driving app running in large-robot Raspberry PI W. It manages serial communication with the MEGA MCU, provides a REST service (over wifi) for remote driving, and provides a limited "autopilot" mode.

The roles between Raspberry PI Pilot app and MEGA MCU code are defined as this :

**Arduino Mega MCU :**

 * Manages low level connect with modules / motor
 * Provides support for all "physical" features. From a simple command, it can drive motors, start RGB Band action, beep ... All actions are processed in small "steps" for simulated multitasking support
 * Provides a continuous feedback ("monitoring") with serial. Gives access to sensors output (converted raw values but not debounced)
 * Implements a basic Serial protocol for command management

 **Raspberry Pi pilot :**

 * Communicate with mega MCU using serial : implements both command and monitoring protocols
 * Provides outside-world interface (a REST service)
 * Debounces monitoring values
 * Manages a clean up-to-date state of the robot (including sensor based values)
 * From the sensor values, try to add basic autopilot features. 

## Autopilot vs Command mode

The application provides 1 remote command mode, 1 fully automatic driving mode and 1 mixed command / autopilot mode


## Configuration

Edit file conf.json

Default :
```json
{
  "RobotName": "Robot-Toby",
  "DebounceBufferSize": 10,
  "DebounceMaxVariation": 10,
  "StandardMove": 100,
  "StandardSpeed": 10,
  "StandardDegree": 90,
  "CentimerToSpeedRatio": 350,
  "DegreeToSpeedRatio": 4,
  "SpeedChangeSteps": 10,
  "LogLevel": "INFO",
  "MinimalStopRange": 10,
  "MonitoringDelayMs": 200,
  "SerialSimulation": "simulations/serial-move.json",
  "SerialDevice": "/dev/serial1",
  "SerialSpeed": 115200,
  "PilotServerPort": 8080,
  "CameraDevice": "/dev/video0"
}

```

## How to build :

Install code into your $GOHOME, then

     go build robot-toby.go
     ./robot-toby.go

**Tested onto :**

* FriendlyElec UbuntuCore 16.3 - NanoPi NEO2 (my testing environment for ARM dev)
* Raspian - RaspberryPi Zero W (robot core platform)

**References :**

Makefile copied from https://github.com/genuinetools projects. Thanks to [Jessie Frazelle](https://github.com/jessfraz)
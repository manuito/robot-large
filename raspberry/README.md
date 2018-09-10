# Pilote application (Golang code for the raspberry Pi W Zero) 

The pilote application is a GO driving app running in large-robot Raspberry PI W. It manages serial communication with the MEGA MCU, provides a REST service (over wifi) for remote driving, and provides a limited "auto-pilote" mode.

The roles between Raspberry PI Pilote app and MEGA MCU code are defined as this :

**Arduino Mega MCU :**

* Manages low level connect with modules / motor
* Provides support for all "physical" features. From a simple command, it can drive motors, start RGB Band action, beep ... All actions are processed in small "steps" for simulated multitasking support
* Provides a continuous feedback ("monitoring") with serial. Gives access to sensors output (converted raw values but not debounced)
* Implements a basic Serial protocol for command management

 **Raspberry Pi pilote :**

* Communicate with mega MCU using serial : implements both command and monitoring protocols
* Provides outside-world interface (a REST service)
* Debounces monitoring values
* Manages a clean up-to-date state of the robot (including sensor based values)
* From the sensor values, try to add basic auto-pilote features. 

The GO project is named "robotoby". The code needs to be checkouted at *$GOHOME/src/robotoby*

## Build / run

I will add a Makefile. For now, use : 

    go build robot-toby.go

For missing libs, run : 

    go get -d ./...

And then run :

    ./robot-toby

## Testing 

A Swagger ui is integrated for rest control. Check log outputs for url

**Tested onto :**

* FriendlyElec UbuntuCore 16.3 - NanoPi NEO2 (my testing environment for ARM dev)
* Raspian - RaspberryPi Zero W (robot core platform)

## Configuration

Use file *conf.json* :

    {
        "RobotName": "Robot-Toby",          // ==> Keep it
        "DebounceBufferSize" : 10,          // ==> For debouncing. Nbr of items controled for AVG value check
        "DebounceMaxVariation" : 10,        // ==> For debouncing. Range of difference before eliminating a value
        "StandardMove" : 100,               // ==> Standard distance moved for 1 move action without arg (in centimeter)
        "StandardSpeed" : 10,               // ==> Standard speed for 1 move action without arg (in centimeter/s)
        "StandardDegree" : 90,              // ==> Standard angle moved for 1 turn action without arg (in degree)
        "CentimerToSpeedRatio" : 350,       // ==> For speed calculation. Need to be tested
        "DegreeToSpeedRatio" : 4,           // ==> For angle calculation. Need to be tested
        "SpeedChangeSteps" : 10,            // ==> For speed variation in 1 command, nbr of intermediate speed
        "LogLevel" : "INFO",                // ==> Internal log level (basic log system)
        "MinimalStopRange" : 10,            // ==> Minimal distance (in centimeters) for obstacle avoidance in auto-drive
        "MonitoringDelayMs" : 200,          // ==> Interval in ms between each sensor serial read
        "SerialSimulation" : "NONE",        // ==> Can enable some test values for serial. Use files into /simulations folder
        "SerialDevice" : "/dev/ttyACM0",    // ==> Serial device in "real" serial connection
        "SerialSpeed" : 115200              // ==> Serial speed (bauds)
    }

The project is organized in 4 packages : 

* **tools** : some internal utils
* **application** : Basic management of application. logging + config + app state management
* **connect** : Serial connection + similation + robot protocol
* **pilote** : Everything for driving the robot. auto pilote, wifi pilote (includes web server with swagger-ui interface) + Monitored state

## Pilote over REST services

The rest interface publish 2 endpoints : 

* **/pilote** : for driving the robot
* **/state** : access to some status of the robot

The application provides 1 remote command mode, 1 fully automatic driving mode and 1 mixed command / autopilote mode

**References :**

Makefile copied from https://github.com/genuinetools projects. Thanks to [Jessie Frazelle](https://github.com/jessfraz)

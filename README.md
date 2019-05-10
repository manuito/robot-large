# My first Robot !

Here I will publish everything I did to build my first "large" robot, using :
* A Pololu Romi base kit 
* The polulu Romi power mgmt / motor driver board
* An Arduina Mega 2560 (Elegoo clone) for electronical support
* Many sensors / outputs integrated through a mega Proto shield
  * 2 ultrasonic sensors
  * 1 50 cm RGB Band (15 addressable RGB leds from Makebloc)
  * 1 8x8 LED matrice
  * Motor encoders (from Pololu)
  * 1 Buzzer
  * 1 I2C 6 axis Gy
* A RaspberryPi W Zero for management / wifi support 
  * With Raspberry Pi camera
* A level shifter 5v-3.3v for UART link between Raspberry and Arduino

![An ugly boy](docs/robot.jpg?raw=true "An ugly boy")

## Building it

I will publish more details on the used components later. On the protoboard their is only socket for the other stuffs, so everything can be re-used for other projects.

Everything is build using standard modules, integrated with a perma-proto board. The perma proto board is an official Mega perma proto, which is the easier to use actually with Mega2560 (I tryed some other + custom perma proto). 

**Update** : robot has now a raspberryPi camera on board. For now it is used only through standard RaspberryPu camera streaming

## Programming it

Two main components to program :
* The arduino handle "low level" features with its numerous GPIOs... 
* The raspberry is "the brain" : auto-pilote, remote control through wifi ...

They talk to each other through a serial connection. A very basic 4 bytes command system is used, with the ability to provides continiously monitoring from the arduino to the raspberry to check speed, location, obstacles ...

### Arduino Mega code

WIP - I'm working on 3 different robot projects on the same time + job + familly = It's improving very slowly ...

See arduino sub-folder. Done with Arduino IDE and various standard libs. 

I use a time splited process, where execution is processed in repeated step, with support for one step action on each device at each repeated step. Currently the step are "ok" but the rate is still it little bit low. And for sensors without prefixed response time (like UltraSonic sensors) it's not constant. I will try other solution during the "improving" steps of the project.

In the last updates, the startup process has been simplified to lower power consumption. Until Raspberry app is started, the Mega board just light up one loading LED.

### Raspberry code

WIP 

See raspberry sub-folder. Done with Golang. Details on code are given in internal README

Started when the raspberry power up (takes ~ 30s to boot). Raspberry uses a minimal Raspian distribution, console only, with SSH enabled. So it is always possible to SSH into the raspberry. 

WIFI uses a dedicated Gateway (A basic Dragino LoRa "Gateway" reserved for my Arduino projects). Access to internet is granted.

I will try to make the pilote code easy to adapt for other robotic project based on Linux ARM32/ARM64 (raspberry pi 3 / pi Zero W / clones like NanoPi Neo2 or Neo Air...)
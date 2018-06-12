# My first Robot !

Here I will publish everything I did to build my first "large" robot, using :
* A Pololu Romi base kit 
* The polulu Romi power mgmt / motor driver board
* An Arduina Mega 2560 (Elegoo) for electronical support
* Many sensors / outputs integrated through a mega Proto shield
  * 2 ultrasonic captors
  * 1 50 cm RGB Band (15 addressable RGB leds from Makebloc)
  * 1 8x8 LED matrice
  * Motor encoders (from Pololu)
  * 1 Buzzer
  * 1 I2C Gy
* A RaspberryPi W Zero for management / wifi support 
* A level shifter 5v-3.3v for UART link between Raspberry and Arduino

## Building it

I will publish more details on the used components later. On the protoboard their is only socket for the other stuffs, so everything can be re-used for other projects.

## Programming it

Two main components to program :
* The arduino handle "low level" features with its numerous GPIOs... 
* The raspberry is "the brain" : auto-pilote, remote control through wifi ...

They talk to each other through a serial connection. A very basic 4 bytes command system is used, with the ability to provides continiously monitoring from the arduino to the raspberry to check speed, location, obstacles ...

A "simplified" command system is also available, using only 1 char commands for testing with a basic serial console.

### Arduino Mega code

WIP

See arduino sub-folder. Done with Arduino IDE and various standard libs

### Raspberry code

WIP 

See raspberry sub-folder. Done with Golang. 

3 parts : 
* Commands => Integration with Arduino (management of serial connection)
* Brain => Auto pilote
* Control => Remote pilote

Started when the raspberry power up (takes ~ 30s to boot). Raspberry uses a minimal Raspian distribution, console only, with SSH enabled. So it is always possible to SSH into the raspberry. 

WIFI uses a dedicated Gateway (A basic Dragino LoRa Gateway reserved for my Arduino projects). Access to internet is granted.
# Robot-large : Serial simulated with usb (leonardo)

The arduino code here is a basic simulation created with a Leonardo compatible board (actually, I use a basic ATmega32u4 badusb ...) : this simply push continuously some serial data, simulating the monitoring flow from the "real" robot Mega board.

I will improve it to make it interactive (in a limited way) : for example, when asked to move, the left/right sensors have new values, the speed increase ...

The values contains some "wrong" lines to check also debouncing and value filtering processed by raspberry core.

To use the usb serial, for example for Ubuntu you can specify in config for SerialDevice : "/dev/ttyACM0"

Don't forget to add current user to *dialout* group :

    sudo usermod -a -G dialout my_user

and logout / login again


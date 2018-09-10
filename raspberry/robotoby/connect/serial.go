package connect

import (
	"bufio"
	"log"
	"robotoby/application"

	"github.com/tarm/serial"
)

/*
 * Low-level connect with Mega core using serial protocol
 */

var useSimulator = application.State.Config.SerialSimulation != "NONE"
var port *serial.Port

// Init Serial
func openSerial() *serial.Port {
	c := &serial.Config{Name: application.State.Config.SerialDevice, Baud: application.State.Config.SerialSpeed}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	application.Info("Serial connection up on " + application.State.Config.SerialDevice)

	return s
}

// sendSerial : Send the given raw value using serial
func sendSerial(value string) error {

	if useSimulator {
		simuSendSerial(value)
		return nil
	}

	if port == nil {
		port = openSerial()
	}

	_, err := port.Write([]byte(value))

	return err
}

// receiveSerial : Get the last received value in serial connect
func receiveSerial() (string, error) {

	if useSimulator {
		return simuReceiveSerial(), nil
	}

	if port == nil {
		port = openSerial()
	}

	// Here receive

	buf := make([]byte, 128)
	_, err := port.Read(buf)

	reader := bufio.NewReader(port)
	reply, err := reader.ReadBytes('\r')
	result := string(reply[:])

	application.Debug("Serial get => " + result)

	return result, err
}

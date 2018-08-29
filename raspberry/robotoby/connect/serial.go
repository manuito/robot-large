package connect

/*
 * Low-level connect with Mega core using serial protocol
 */

const useSimulator bool = true

// sendSerial : Send the given raw value using serial
func sendSerial(value string) error {

	if useSimulator {
		simuSendSerial(value)
		return nil
	}

	// Here send

	return nil
}

// receiveSerial : Get the last received value in serial connect
func receiveSerial() (string, error) {

	if useSimulator {
		return simuReceiveSerial(), nil
	}

	// Here receive

	return "value", nil
}

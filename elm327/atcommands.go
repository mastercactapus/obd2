package elm327

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Description will return the device description
func (c *ELM327) Description() (string, error) {
	return c.SendAT("@1")
}

// ID will return the device identifier
func (c *ELM327) ID() (string, error) {
	return c.SendAT("@2")
}

// ReadVoltage will read the input voltage
func (c *ELM327) ReadVoltage() (float64, error) {
	res, err := c.SendAT("RV")
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(strings.TrimSuffix(res, "V"), 32)
}

// CalibrateVoltage will calibrate the voltage (limit 2 decimal places)
// if set to 0, factory default will be restored
func (c *ELM327) CalibrateVoltage(v float64) error {
	if v < 0 || v > 99.99 {
		return errors.New("voltage must be between 0 and 99.99")
	}
	str := strings.Replace(fmt.Sprintf("%05.2f", v), ".", "", 1)
	_, err := c.SendAT("CV " + str)
	return err
}

func (c *ELM327) GetIgnition() (bool, error) {
	res, err := c.SendAT("IGN")
	return res == "ON", err
}

func (c *ELM327) SendAT(command string) (string, error) {
	command = "AT" + command + "\r"
	if c.last == command {
		command = "\r"
	} else {
		c.last = command
	}
	_, err := io.WriteString(c.w, command)
	if err != nil {
		return "", err
	}
	response, err := c.r.ReadString('>')
	if err != nil {
		return "", err
	}
	response = strings.TrimSpace(response[len(command) : len(response)-1])
	// according to docs, this can happen and should be ignored...
	response = strings.Replace(response, "\x00", "", -1)
	if response == "?" {
		return response, ErrUnknownCommand
	}
	return response, nil
}
func (c *ELM327) Reset() error {
	_, err := c.SendAT("Z")
	return err
}

package obd2

const (
	// PIDSupport1 will return supported PIDs from 0x01 to 0x20
	PIDSupport1 byte = 0x00

	// PIDMonitorStatus is used to monitor status since DTCs were cleared
	PIDMonitorStatus byte = 0x01

	// PIDFreeze will freeze DTC values
	PIDFreeze byte = 0x02

	// PIDFuelSystemStatus will request the loop status of the fuel system(s)
	PIDFuelSystemStatus byte = 0x03

	// PIDEngineLoad will request the calculated engine load
	PIDEngineLoad byte = 0x04

	// PIDECT will request the engine coolant temperature
	PIDECT byte = 0x05

	// PIDSTFTBank1 is for short term fuel trim (bank 1)
	PIDSTFTBank1 byte = 0x06

	// PIDLTFTBank1 is for long term fuel trim (bank 1)
	PIDLTFTBank1 byte = 0x07

	// PIDSTFTBank2 is for short term fuel trim (bank 2)
	PIDSTFTBank2 byte = 0x08

	// PIDLTFTBank2 is for long term fuel trim (bank 2)
	PIDLTFTBank2 byte = 0x09

	// PIDFuelPressure is for Fuel pressure (guage pressure)
	PIDFuelPressure byte = 0x0a

	// PIDIntakeMAP is for intake manifold absolute pressure
	PIDIntakeMAP byte = 0x0b

	// PIDEngineRPM is for engine RPM
	PIDEngineRPM byte = 0x0c

	// PIDVehicleSpeed is for the vehicle speed
	PIDVehicleSpeed byte = 0x0d

	// PIDTimingAdvance is the timing advance of the engine
	PIDTimingAdvance byte = 0x0e

	// PIDIntakeAirTemp is the temperature of the air (at intake)
	PIDIntakeAirTemp byte = 0x0f
	// PIDMAFRate is the rate of air flow through the Mass Air Flow sensor
	PIDMAFRate byte = 0x10

	// PIDThrottlePos gives the current position of the throttle
	PIDThrottlePos byte = 0x11

	// PIDComAirStatus gives the commanded secondary air status
	PIDComAirStatus byte = 0x12

	// Mode1O2Present will determine what O2 sensors are present
	Mode1O2Present byte = 0x13
)

const (
	// FuelSystemStatusOpenTemp means running in open loop due to insufficient engine temperature
	FuelSystemStatusOpenTemp byte = 1

	//FuelSystemStatusClosed means running in closed loop, using oxygen sensor feedback to determine fuel mix
	FuelSystemStatusClosed byte = 2

	//FuelSystemStatusOpenLoad means running in open loop due to engine load, or fuel cut due to deceleration
	FuelSystemStatusOpenLoad byte = 4

	// FuelSystemStatusOpenFailure means running in open loop due to system failure
	FuelSystemStatusOpenFailure byte = 8

	// FuelSystemStatusClosedFault means running in closed loop (at least 1 O2 sensor) but there is a fault in the feedback system
	FuelSystemStatusClosedFault byte = 16
)

// IsSupported will determine if a particular PID is supported given the response slice. The slice must be long enough to contain
// the PID or it will panic. (1 bit per PID, 8 bits per byte; so to check PID 10/0x0a, res must have a length of at least 2)
func IsSupported(res []byte, pid byte) bool {
	byteIndex := pid / 8
	bitIndex := pid % 8
	return (res[byteIndex] & (1 << bitIndex)) == 1
}

// DecodeFuelTrim will return the fuel trim value as a percentage of rich or lean (-1 to 1, respectively)
func DecodeFuelTrim(v byte) float64 {
	return float64(v)/128.0 - 100.0
}

// DecodeEngineLoad will decode the engine load as a percentage (between 0 and 1)
func DecodeEngineLoad(v byte) float64 {
	return float64(v) / 255.0
}

// DecodeECT will convert the response for PIDECT to degrees Celsius
func DecodeECT(v byte) int {
	return int(v) - 40
}

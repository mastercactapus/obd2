package mode1

const (
	// ID is the identifier to use all mode1 commands
	ID byte = 0x01

	// PIDSupport1 will return supported PIDs from 0x01 to 0x20
	PIDSupport1 byte = 0x00

	// PIDMonitorStatus is used to monitor status since DTCs were cleared. Use with DecodeMonitorStatus
	PIDMonitorStatus byte = 0x01

	// PIDFreeze will freeze DTC values
	PIDFreeze byte = 0x02

	// PIDFuelSystemStatus will request the loop status of the fuel system(s). Response should be two bytes, one for each possible fuel system. The values can be matched against FuelSystemStatus
	PIDFuelSystemStatus byte = 0x03

	// PIDEngineLoad will request the calculated engine load. Use with DecodeEngineLoad
	PIDEngineLoad byte = 0x04

	// PIDECT will request the engine coolant temperature. Use with DecodeECT
	PIDECT byte = 0x05

	// PIDSTFTBank1 is for short term fuel trim (bank 1). Use with DecodeFuelTrim
	PIDSTFTBank1 byte = 0x06

	// PIDLTFTBank1 is for long term fuel trim (bank 1). Use with DecodeFuelTrim
	PIDLTFTBank1 byte = 0x07

	// PIDSTFTBank2 is for short term fuel trim (bank 2). Use with DecodeFuelTrim
	PIDSTFTBank2 byte = 0x08

	// PIDLTFTBank2 is for long term fuel trim (bank 2). Use with DecodeFuelTrim
	PIDLTFTBank2 byte = 0x09

	// PIDFuelPressure is for Fuel pressure (guage pressure). Use with DecodeFuelPressure
	PIDFuelPressure byte = 0x0a

	// PIDIntakeMAP is for intake manifold absolute pressure. Response is in kPa
	PIDIntakeMAP byte = 0x0b

	// PIDEngineRPM is for engine RPM. Use with DecodeEngineRPM
	PIDEngineRPM byte = 0x0c

	// PIDVehicleSpeed is for the vehicle speed. Response is in km/h
	PIDVehicleSpeed byte = 0x0d

	// PIDTimingAdvance is the timing advance of the engine. Use with DecodeTimingAdvance
	PIDTimingAdvance byte = 0x0e

	// PIDIAT is the temperature of the air (at intake)
	PIDIAT byte = 0x0f

	// PIDMAFRate is the rate of air flow through the Mass Air Flow sensor
	PIDMAFRate byte = 0x10

	// PIDThrottlePos gives the current position of the throttle
	PIDThrottlePos byte = 0x11

	// PIDComAirStatus gives the commanded secondary air status
	PIDComAirStatus byte = 0x12

	// Mode1O2Present will determine what O2 sensors are present
	Mode1O2Present byte = 0x13
)

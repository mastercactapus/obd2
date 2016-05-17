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

	// PIDFuelSystemStatus will request the loop status of the fuel system(s). Response should be two bytes, one for each possible fuel system. The values can be matched against FuelSystemStatus values
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

	// PIDIAT is the temperature of the air (at intake). Use with DecodeIAT
	PIDIAT byte = 0x0f

	// PIDMAFRate is the rate of air flow through the Mass Air Flow sensor. Use with DecodeMAFRate
	PIDMAFRate byte = 0x10

	// PIDThrottlePos gives the current position of the throttle. Use with DecodeThrottlePos
	PIDThrottlePos byte = 0x11

	// PIDComAirStatus gives the commanded secondary air status. Can be matched against SecondaryAirStatus values
	PIDComAirStatus byte = 0x12

	// PIDO2Present will determine what O2 sensors are present. Use with DecodeO2Present
	PIDO2Present byte = 0x13

	// PIDO2STFT1 requests the short term fuel trim and voltage for oxygen sensor 1. Use with DecodeO2STFT
	PIDO2STFT1 byte = 0x14

	// PIDO2STFT2 requests the short term fuel trim and voltage for oxygen sensor 2. Use with DecodeO2STFT
	PIDO2STFT2 byte = 0x15

	// PIDO2STFT3 requests the short term fuel trim and voltage for oxygen sensor 3. Use with DecodeO2STFT
	PIDO2STFT3 byte = 0x16

	// PIDO2STFT4 requests the short term fuel trim and voltage for oxygen sensor 4. Use with DecodeO2STFT
	PIDO2STFT4 byte = 0x17

	// PIDO2STFT5 requests the short term fuel trim and voltage for oxygen sensor 5. Use with DecodeO2STFT
	PIDO2STFT5 byte = 0x18

	// PIDO2STFT6 requests the short term fuel trim and voltage for oxygen sensor 6. Use with DecodeO2STFT
	PIDO2STFT6 byte = 0x19

	// PIDO2STFT7 requests the short term fuel trim and voltage for oxygen sensor 7. Use with DecodeO2STFT
	PIDO2STFT7 byte = 0x1a

	// PIDO2STFT8 requests the short term fuel trim and voltage for oxygen sensor 8. Use with DecodeO2STFT
	PIDO2STFT8 byte = 0x1b

	// PIDOBDStandard requests what OBD (onboard diagnostic) standards the vehicle conforms to. Can be matched against OBDStandard values
	PIDOBDStandard byte = 0x1c

	// PIDO2PresentExt will determine what O2 sensors are present. This is an extended version for up to 4 banks. Use with DecodeO2Present for each byte (1 byte = 2 banks)
	PIDO2PresentExt byte = 0x1d

	// PIDAuxInput will request the auxilliary input status. Use with DecodeAuxInput
	PIDAuxInput byte = 0x1e

	// PIDRunTime will request the run time since engine start in seconds. Use with DecodeRunTime
	PIDRunTime byte = 0x1f
)

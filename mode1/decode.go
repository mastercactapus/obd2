package mode1

import (
	"encoding/binary"
	"time"
)

// FuelSystemStatus is the status returned for the fuel system
type FuelSystemStatus byte

const (
	// FuelSystemStatusOpenTemp means running in open loop due to insufficient engine temperature
	FuelSystemStatusOpenTemp FuelSystemStatus = 1

	//FuelSystemStatusClosed means running in closed loop, using oxygen sensor feedback to determine fuel mix
	FuelSystemStatusClosed FuelSystemStatus = 2

	//FuelSystemStatusOpenLoad means running in open loop due to engine load, or fuel cut due to deceleration
	FuelSystemStatusOpenLoad FuelSystemStatus = 4

	// FuelSystemStatusOpenFailure means running in open loop due to system failure
	FuelSystemStatusOpenFailure FuelSystemStatus = 8

	// FuelSystemStatusClosedFault means running in closed loop (at least 1 O2 sensor) but there is a fault in the feedback system
	FuelSystemStatusClosedFault FuelSystemStatus = 16
)

// ComAirStatus is the status returned for the secondary air system
type ComAirStatus byte

const (
	// ComAirStatusUpstream means secondary air is on from upstream
	ComAirStatusUpstream ComAirStatus = 1

	// ComAirStatusDownstream means secondary air is on, downstream of catalytic converter
	ComAirStatusDownstream ComAirStatus = 2

	// ComAirStatusOff means secondary air is off, or coming from the outside atmosphere
	ComAirStatusOff ComAirStatus = 4

	// ComAirStatusDiag means the secondary air pump has been commanded on for diagnostics
	ComAirStatusDiag ComAirStatus = 8
)

// MonitorStatusSpark contains monitoring status for spark-ignition engines
type MonitorStatusSpark struct {
	// Catalyst contains status and availability of the catalyst test
	Catalyst TestStatus

	// HeatedCatalyst contains status and availability of the heated catalyst test
	HeatedCatalyst TestStatus

	// EvapSystem contains status and availability of the evaporative system test
	EvapSystem TestStatus

	// SecondaryAir contains status and availability of the secondary air system test
	SecondaryAir TestStatus

	// ACRefrigerant contains status and availability of the A/C refrigerant test
	ACRefrigerant TestStatus

	// O2Sensor contains status and availability of the oxygen sensor test
	O2Sensor TestStatus

	// O2SensorHeater contains status and availability of the oxygen sensor heater test
	O2SensorHeater TestStatus

	// EGRSystem contains status and availability of the EGR (exhaust gas relfow) System test
	EGRSystem TestStatus
}

// MonitorStatusCompression contains monitoring status for compression-ignition engines
type MonitorStatusCompression struct {
	// NMHCCatalyst contains status and availability of the non-methane hydrocarbons catalyst test
	NMHCCatalyst TestStatus

	// NOxSCRMonitor contains status and availability of the NOx/SCR monitor test
	NOxSCRMonitor TestStatus

	// BoostPressure contains status and availability of the boost pressure test
	BoostPressure TestStatus

	// ExhauseGasSensor contains status and availability of the exhaust gas sensor test
	ExhauseGasSensor TestStatus

	// PMFilter contains status and availability of the PM filter monitor test
	PMFilter TestStatus

	// EGRVTT contains status and availability of the EGS and/or VTT system test
	EGRVTT TestStatus
}

// MonitorStatus represents the monitoring systems status
type MonitorStatus struct {
	// MIL indicates if the CEL/MIL is on (or should be)
	MIL bool

	//DTCCount indicates the number of emissions-related DTCs available/set
	DTCCount int

	// Misfire contains status and availability of the misfire test
	Misfire TestStatus

	// FuelSystem contains status and availability of the fuel system test
	FuelSystem TestStatus

	// Components contains status and availability of the components test
	Components TestStatus

	// Spark will contain data if this is a spark-ignition system. It will be nil otherwise
	Spark *MonitorStatusSpark

	// Compression will contain data if this is a compression-ignition system. It will be nil otherwise
	Compression *MonitorStatusCompression
}

// TestStatus indicates the status of a monitoring test
type TestStatus struct {

	// Available will be set if the test is available and enabled
	Available bool

	// Complete will be set if the test has completed
	Complete bool
}

// DecodeMonitorStatus will decode the response of a PIDMonitorStatus request. res must be at least 4 bytes
func DecodeMonitorStatus(res []byte) MonitorStatus {
	var s MonitorStatus
	s.MIL = (1<<7)&res[0] == 1
	s.DTCCount = int((1 << 7) ^ res[0])

	// FYI B7 should be 0 -- reserved, but we don't care
	s.Misfire.Available = (res[1] & 1) == 1
	s.Misfire.Complete = (res[1] & (1 << 4)) == 0
	s.FuelSystem.Available = (res[1] & (1 << 1)) == 1
	s.FuelSystem.Complete = (res[1] & (1 << 5)) == 0
	s.Components.Available = (res[1] & (1 << 2)) == 1
	s.Components.Complete = (res[1] & (1 << 6)) == 0

	comp := res[1]&(1<<3) == 1

	if comp {
		c := new(MonitorStatusCompression)
		c.NMHCCatalyst.Available = (res[2] & 1) == 1
		c.NMHCCatalyst.Complete = (res[3] & 1) == 0
		c.NOxSCRMonitor.Available = (res[2] & (1 << 1)) == 1
		c.NOxSCRMonitor.Complete = (res[3] & (1 << 1)) == 0
		c.BoostPressure.Available = (res[2] & (1 << 3)) == 1
		c.BoostPressure.Complete = (res[3] & (1 << 3)) == 0
		c.ExhauseGasSensor.Available = (res[2] & (1 << 5)) == 1
		c.ExhauseGasSensor.Complete = (res[3] & (1 << 5)) == 0
		c.PMFilter.Available = (res[2] & (1 << 6)) == 1
		c.PMFilter.Complete = (res[3] & (1 << 6)) == 0
		c.EGRVTT.Available = (res[2] & (1 << 7)) == 1
		c.EGRVTT.Complete = (res[3] & (1 << 7)) == 0
		s.Compression = c
	} else {
		sp := new(MonitorStatusSpark)
		sp.Catalyst.Available = (res[2] & 1) == 1
		sp.Catalyst.Complete = (res[3] & 1) == 0
		sp.HeatedCatalyst.Available = (res[2] & (1 << 1)) == 1
		sp.HeatedCatalyst.Complete = (res[3] & (1 << 1)) == 0
		sp.EvapSystem.Available = (res[2] & (1 << 2)) == 1
		sp.EvapSystem.Complete = (res[3] & (1 << 2)) == 0
		sp.SecondaryAir.Available = (res[2] & (1 << 3)) == 1
		sp.SecondaryAir.Complete = (res[3] & (1 << 3)) == 0
		sp.ACRefrigerant.Available = (res[2] & (1 << 4)) == 1
		sp.ACRefrigerant.Complete = (res[3] & (1 << 4)) == 0
		sp.O2Sensor.Available = (res[2] & (1 << 5)) == 1
		sp.O2Sensor.Complete = (res[3] & (1 << 5)) == 0
		sp.O2SensorHeater.Available = (res[2] & (1 << 6)) == 1
		sp.O2SensorHeater.Complete = (res[3] & (1 << 6)) == 0
		sp.EGRSystem.Available = (res[2] & (1 << 7)) == 1
		sp.EGRSystem.Complete = (res[3] & (1 << 7)) == 0
		s.Spark = sp
	}
	return s
}

// DecodeFuelPressure will return the fuel pressure in kPa
func DecodeFuelPressure(v byte) int {
	return int(v) * 3
}

// IsSupported will determine if a particular PID is supported given the response slice. The slice must be long enough to contain
// the PID or it will panic. (1 bit per PID, 8 bits per byte; so to check PID 10/0x0a, res must have a length of at least 2)
func IsSupported(res []byte, pid byte) bool {
	byteIndex := pid / 8
	bitIndex := pid % 8
	return (res[byteIndex] & (1 << bitIndex)) == 1
}

// DecodeFuelTrim will return the fuel trim value as a percentage of rich or lean (-1 to 1, respectively)
func DecodeFuelTrim(v byte) float64 {
	return float64(v)/128 - 100
}

// DecodeEngineLoad will decode the engine load as a percentage (between 0 and 1)
func DecodeEngineLoad(v byte) float64 {
	return float64(v) / 255
}

// DecodeECT will convert the response for PIDECT to degrees Celsius
func DecodeECT(v byte) int {
	return int(v) - 40
}

// DecodeEngineRPM will convert the response to RPM (rotations per minute). res must be at least 2 bytes
func DecodeEngineRPM(res []byte) float64 {
	return (float64(res[0])*256 + float64(res[1])) / 4
}

// DecodeTimingAdvance will convert the response to degrees before TDC (top dead center)
func DecodeTimingAdvance(v byte) float64 {
	return float64(v)/2 - 64
}

// DecodeIAT will convert the response for PIDIAT to degrees Celsius
func DecodeIAT(v byte) int {
	return int(v) - 40
}

// DecodeMAFRate will convert the response to grams/sec. res must be at least 2 bytes
func DecodeMAFRate(res []byte) float64 {
	return (float64(res[0])*256 + float64(res[1])) / 100
}

// DecodeThrottlePos will decode the throttle position as a percentage (between 0 and 1)
func DecodeThrottlePos(v byte) float64 {
	return float64(v) / 255
}

// O2Present is the decoded result of the O2Present command
type O2Present struct {
	Bank1 [4]bool
	Bank2 [4]bool
}

// DecodeO2Present will decode what oxygen sensors are present. Must have 2 or 4 bytes
func DecodeO2Present(v byte) O2Present {
	var p O2Present
	p.Bank1[0] = v&1 == 1
	p.Bank1[1] = v&(1<<1) == 1
	p.Bank1[2] = v&(1<<2) == 1
	p.Bank1[3] = v&(1<<3) == 1
	p.Bank2[0] = v&(1<<4) == 1
	p.Bank2[1] = v&(1<<5) == 1
	p.Bank2[2] = v&(1<<6) == 1
	p.Bank2[3] = v&(1<<7) == 1
	return p
}

// O2STFT is the oxygen sensor short term fuel trim stats
type O2STFT struct {
	// STFT represents the short-term fuel trim as a percentage of rich or lean (-1 to 1, respectively)
	STFT float64

	// Voltage is the voltage measured
	Voltage float64

	// if SensorUsed is false, this sensor is not used in fuel trim calculation
	SensorUsed bool
}

// DecodeO2STFT will decode the response for a PIDO2STFTx request. res must be 2 bytes
func DecodeO2STFT(res []byte) O2STFT {
	var o O2STFT
	o.Voltage = float64(res[0]) / 200
	o.STFT = 100*float64(res[1])/128 - 100
	o.SensorUsed = res[1] != 0xff
	return o
}

// AuxInputStatus represents auxilliary input status
type AuxInputStatus struct {
	// PTO represents the PTO (Power Take Off) status, if true PTO is active
	PTO bool
}

// DecodeAuxInput will decode the response to PIDAuxInput
func DecodeAuxInput(v byte) AuxInputStatus {
	var s AuxInputStatus
	s.PTO = v&1 == 1
	return s
}

// DecodeRunTime will decode engine runtime from a response. res must be 2 bytes
func DecodeRunTime(res []byte) time.Duration {
	return time.Duration(binary.BigEndian.Uint16(res))
}

// OBDStandard represents an OBD standard as returned by PIDOBDStandard
type OBDStandard byte

const (
	// OBDStandardOBD2CARB means OBD-II as defined by the CARB
	OBDStandardOBD2CARB OBDStandard = 1

	// OBDStandardOBD1EPA means OBD-I as defined by the EPA
	OBDStandardOBD1EPA OBDStandard = 2

	// OBDStandardOBD1OBD2 means OBD-I and OBD-II
	OBDStandardOBD1OBD2 OBDStandard = 3

	// OBDStandardOBD1 means OBD-I
	OBDStandardOBD1 OBDStandard = 4

	// OBDStandardNone means the vehicle is not OBD compliant
	OBDStandardNone OBDStandard = 5

	// OBDStandardEOBD means EOBD (Europe)
	OBDStandardEOBD OBDStandard = 6

	// OBDStandardEOBDOBD2 means EOBD and OBD-II
	OBDStandardEOBDOBD2 OBDStandard = 7

	// OBDStandardEOBDOBD2 means EOBD and OBD-I
	OBDStandardEOBDOBD1 OBDStandard = 8

	// OBDStandardEOBDOBD1OBD2 means EOBD, OBD-I and OBD-II
	OBDStandardEOBDOBD1OBD2 OBDStandard = 9

	// OBDStandardJOBD means JOBD (Japan)
	OBDStandardJOBD OBDStandard = 10

	// OBDStandardJOBD means JOBD and OBD-II
	OBDStandardJOBD OBDStandard = 11

	// OBDStandardJOBDEOBD means JOBD and EOBD
	OBDStandardJOBDEOBD OBDStandard = 12

	// OBDStandardJOBDEOBDOBD2 means JOBD, EOBD, and OBD-II
	OBDStandardJOBDEOBDOBD2 OBDStandard = 13

	/* 14-16 Reserved */

	// OBDStandardEMD means Engine Manufacturer Diagnostics (EMD)
	OBDStandardEMD OBDStandard = 17

	// OBDStandardEMD means Engine Manufacturer Diagnostics Enhanced (EMD+)
	OBDStandardEMDPlus OBDStandard = 18

	// OBDStandardHDOBDC means Heavy Duty On-Board Diagnostics (Child/Partial) (HD OBD-C)
	OBDStandardHDOBDC OBDStandard = 19

	// OBDStandardHDOBD means Heavy Duty On-Board Diagnostics (HD OBD)
	OBDStandardHDOBD OBDStandard = 20

	// OBDStandardWWHOBD means World Wide Harmonized OBD (WWH OBD)
	OBDStandardWWHOBD OBDStandard = 21

	/* 22 Reserved */

	// OBDStandardHDEOBD1 means Heavy Duty Euro OBD Stage I without NOx control (HD EOBD-I)
	OBDStandardHDEOBD1

	// OBDStandardHDEOBD1N means Heavy Duty Euro OBD Stage I with NOx control (HD EOBD-I N)
	OBDStandardHDEOBD1N

	// OBDStandardHDEOBD2 means Heavy Duty Euro OBD Stage II without NOx control (HD EOBD-II)
	OBDStandardHDEOBD2

	// OBDStandardHDEOBD2N means Heavy Duty Euro OBD Stage II with NOx control (HD EOBD-II N)
	OBDStandardHDEOBD2N

	/* 27 Reserved */

	// OBDStandardOBDBr1 means Brazil OBD Phase 1 (OBDBr-1)
	OBDStandardOBDBr1

	// OBDStandardOBDBr2 means Brazil OBD Phase 2 (OBDBr-2)
	OBDStandardOBDBr2

	// OBDStandardKOBD means Korean OBD (KOBD)
	OBDStandardKOBD

	// OBDStandardIOBD1 means India OBD I (IOBD I)
	OBDStandardIOBD1

	// OBDStandardIOBD2 means India OBD II (IOBD II)
	OBDStandardIOBD2

	// OBDStandardHDEOBD6 means Heavy Duty Euro OBD Stage VI (HD EOBD-IV)
	OBDStandardHDEOBD6

	/* 34-250 Reserved */

	/* 251-255 Not Available */
)

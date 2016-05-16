package mode1

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

// DecodeEngineRPM will convert the response to RPM (rotations per minute). res must be at least 2 bytes
func DecodeEngineRPM(res []byte) float64 {
	return (float64(res[0])*256.0 + float64(res[1])) / 4.0
}

// DecodeTimingAdvance will convert the response to degrees before TDC (top dead center)
func DecodeTimingAdvance(v byte) float64 {
	return float64(v)/2.0 - 64.0
}

//DecodeIntakeAirTemp will convert the response for PIDIAT to degrees Celsius
func DecodeIntakeAirTemp(v byte) int {
	return int(v) - 40
}

// DecodeMAFRate will convert the response to grams/sec. res must be at least 2 bytes
func DecodeMAFRate(res []byte) float64 {
	return (float64(res[0])*256.0 + float64(res[1])) / 100.0
}

// DecodeThrottlePos will decode the throttle position as a percentage (between 0 and 1)
func DecodeThrottlePos(v byte) float64 {
	return float64(v) / 255.0
}

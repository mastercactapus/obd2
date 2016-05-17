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

	// OBDStandardJOBDOBD2 means JOBD and OBD-II
	OBDStandardJOBDOBD2 OBDStandard = 11

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
	OBDStandardHDEOBD1 OBDStandard = 23

	// OBDStandardHDEOBD1N means Heavy Duty Euro OBD Stage I with NOx control (HD EOBD-I N)
	OBDStandardHDEOBD1N OBDStandard = 24

	// OBDStandardHDEOBD2 means Heavy Duty Euro OBD Stage II without NOx control (HD EOBD-II)
	OBDStandardHDEOBD2 OBDStandard = 25

	// OBDStandardHDEOBD2N means Heavy Duty Euro OBD Stage II with NOx control (HD EOBD-II N)
	OBDStandardHDEOBD2N OBDStandard = 26

	/* 27 Reserved */

	// OBDStandardOBDBr1 means Brazil OBD Phase 1 (OBDBr-1)
	OBDStandardOBDBr1 OBDStandard = 28

	// OBDStandardOBDBr2 means Brazil OBD Phase 2 (OBDBr-2)
	OBDStandardOBDBr2 OBDStandard = 29

	// OBDStandardKOBD means Korean OBD (KOBD)
	OBDStandardKOBD OBDStandard = 30

	// OBDStandardIOBD1 means India OBD I (IOBD I)
	OBDStandardIOBD1 OBDStandard = 31

	// OBDStandardIOBD2 means India OBD II (IOBD II)
	OBDStandardIOBD2 OBDStandard = 32

	// OBDStandardHDEOBD6 means Heavy Duty Euro OBD Stage VI (HD EOBD-IV)
	OBDStandardHDEOBD6 OBDStandard = 33

	/* 34-250 Reserved */

	/* 251-255 Not Available */
)

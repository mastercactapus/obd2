package obd2

import (
	"errors"
	"strconv"
)

// DTC represents a single trouble code
type DTC struct {
	// Type is the type of error code
	Type DTCType

	// Category tells if this is a manufacturer-specific code or not
	Category DTCCategory

	// System represents where the trouble code came from
	System DTCSystem

	// Fault is the actual fault-index of this code
	Fault int
}
type DTCType rune
type DTCCategory rune
type DTCSystem rune

const (
	DTCTypePowertrain DTCType = 'P'
	DTCTypeBody       DTCType = 'B'
	DTCTypeChassis    DTCType = 'C'
	DTCTypeNetwork    DTCType = 'U'
)

const (
	DTCCategorySAE          DTCCategory = '0'
	DTCCategoryManufacturer DTCCategory = '1'
)

const (
	DTCSystemAirFuel         DTCSystem = '1'
	DTCSystemAirFuelInjector DTCSystem = '2'
	DTCSystemIgnition        DTCSystem = '3'
	DTCSystemEmissions       DTCSystem = '4'
	DTCSystemSpeedIdle       DTCSystem = '5'
	DTCSystemComputer        DTCSystem = '6'
	DTCSystemTransimission1  DTCSystem = '7'
	DTCSystemTransimission2  DTCSystem = '8'
)

// String returns the string representation of the trouble code, as it would appear on a scanner
func (d DTC) String() string {
	if d.Fault < 0 || d.Fault > 100 {
		panic("fault value out of bounds")
	}
	return string(d.Type) + string(d.Category) + string(d.System) + strconv.Itoa(d.Fault)
}

// ParseDTC will parse the trouble code into it's relevant parts. The code must be 5 characters long.
func ParseDTC(s string) (*DTC, error) {
	if len(s) != 5 {
		return nil, errors.New("invalid length")
	}
	f, err := strconv.Atoi(s[3:])
	if err != nil {
		return nil, err
	}
	d := &DTC{Fault: f}
	switch s[0] {
	case 'P', 'B', 'C', 'U':
		d.Type = DTCType(s[0])
	default:
		return nil, errors.New("bad type specifier")
	}
	switch s[1] {
	case '0', '1':
		d.Category = DTCCategory(s[1])
	default:
		return nil, errors.New("bad category specifier")
	}
	switch s[3] {
	case '1', '2', '3', '4', '5', '6', '7', '8':
		d.System = DTCSystem(s[3])
	default:
		return nil, errors.New("invalid system specifier")
	}

	return d, nil
}

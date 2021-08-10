// Package lenconv performs distance (meter and foot) conversions
package lenconv

import (
	"fmt"
	"math"
)

// Meters is a dedicated type for a length in meters
type Meters float64

// Feet is a dedicated type for a length in feet
type Feet float64

// Some well known constants
const (
	OneFootM      Meters = 0.3048 // Number of meters per foot
	OneMeterF     Feet   = 3.2809 // Number of feet per meter
	InchesPerFoot int    = 12     // Number of inches per foot
)

// Format a distance in meters
func (m Meters) String() string { return fmt.Sprintf("%g m", m) }

// Format a distance in feet/inches
func (f Feet) String() string {
	return fmt.Sprintf("%g' %g\"", math.Floor(float64(f)), math.Mod(float64(f), 1.0)*float64(InchesPerFoot))
}

// Package tempconv performs temperature (Celsius, Fahrenheit and Kelvin) conversions.
package tempconv

import "fmt"

// Celsius is a dedicated type for a temperature in Celsius degrees
type Celsius float64

// Fahrenheit is a dedicated type for a temperature in Fahrenheit degrees
type Fahrenheit float64

// Some well known constants
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// Format a Celcius temperature
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// Format a Fahrenheuit temperature
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// Package tempconv performs Celsius, Fahrenheit and Kelvin conversions.
package tempconv

import "fmt"

// Celsius is a dedicated type for a temperature in Celsius degrees
type Celsius float64

// Fahrenheit is a dedicated type for a temperature in Fahrenheit degrees
type Fahrenheit float64

// Kelvin is a dedicated type for a temperature in Kelvin degrees
type Kelvin float64

// Some well known constants
const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingK     Kelvin     = 273.15
	FreezingC     Celsius    = 0
	FreezingF     Fahrenheit = 32
	BoilingK      Kelvin     = 373.15
	BoilingC      Celsius    = 100
	BoilingF      Fahrenheit = 212
)

// Format a Celsius temperature
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// Format a Fahrenheit temperature
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// Format a Kelvin temperature
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

// Package weightconv performs weight (kilo, pounds) conversions
package weightconv

import (
	"fmt"
)

// Kilograms is a dedicated type for a number of kilograms
type Kilograms float64

// Pounds is a dedicted type for a number of pounds
type Pounds float64

// Some well known constants
const (
	OnePoundK    Kilograms = 0.453592370   // Number of kilos per pound
	OneKilogramL Pounds    = 2.20462262185 // Number of pounds per kilo
)

// Format a weight in kilos
func (k Kilograms) String() string { return fmt.Sprintf("%g kg", k) }

// Format a weight in pounds
func (p Pounds) String() string { return fmt.Sprintf("%g lb", p) }

// Package weightconv performs weight (kilo, pounds) conversions 
package weightconv

import (
	"fmt"
)

type Kilograms float64
type Pounds float64

// Some well known constants
const (
	OnePoundK Kilograms = 0.453592370 // Number of kilo per pound
	OneKilogramL Pounds = 2.20462262185 // Number of pounds per kilo
)

// Format a weight in kilos
func (k Kilograms) String() string { return fmt.Sprintf("%g kg", k)}

// Format a weight in pounds
func (p Pounds) String() string { return fmt.Sprintf("%g lb", p)}

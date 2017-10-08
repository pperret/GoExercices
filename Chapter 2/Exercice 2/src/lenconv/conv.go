package lenconv

// MToF converts a distance from meters to feet
func MToF(m Meters) Feet { return Feet(m/OneFootM)}

// FToM converts a distance from feet to meters
func FToM(f Feet) Meters { return Meters(f/OneMeterF)}

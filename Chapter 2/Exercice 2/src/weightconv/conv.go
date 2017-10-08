package weightconv

// KToL converts kilos to pounds
func KToL(k Kilograms) Pounds { return Pounds(k/OnePoundK)}

// LToK converts pounds to kilos
func LToK(l Pounds) Kilograms { return Kilograms(l/OneKilogramL)}


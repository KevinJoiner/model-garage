package ruptela

const (
	// StatusEventDS is the data version for status events.
	StatusEventDS = "r/v0/s"
	// DevStatusDS is the data version for device status events.
	DevStatusDS = "r/v0/dev"
	// LocationEventDS is the data version for location events.
	LocationEventDS = "r/v0/loc"
)

// fuelTypeConversion Encodings taken from https://en.wikipedia.org/wiki/OBD-II_PIDs#Fuel_Type_Coding
func fuelTypeConversion(val float64) (string, error) {
	switch val {
	case 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32:
		return "COMBUSTION", nil
	case 8, 15:
		return "ELECTRIC", nil
	case 16, 17, 18, 19, 20, 21, 22:
		return "HYBRID", nil
	default:
		return "", errNotFound
	}
}

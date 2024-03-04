// Package convert holds common functionality for generating conversion functions.
package convert

import (
	"math"
)

// Float64Tofloat32 converts float64 to float32.
func Float64Tofloat32(val float64) float32 {
	if val > math.MaxFloat32 {
		return math.MaxFloat32
	}
	if val < math.SmallestNonzeroFloat32 {
		return math.SmallestNonzeroFloat32
	}
	return float32(val)
}

// Float64toUint32 converts float64 to uint32.
func Float64toUint32(val float64) uint32 {
	if val > math.MaxUint32 {
		return math.MaxUint32
	}
	if val < 0 {
		return 0
	}
	return uint32(val)
}

// Float64toUint16 converts float64 to uint16.
func Float64toUint16(val float64) uint16 {
	if val > math.MaxUint16 {
		return math.MaxUint16
	}
	if val < 0 {
		return 0
	}
	return uint16(val)
}

// Float64toUint16 converts float64 to int16.
func Float64toInt16(val float64) int16 {
	if val > math.MaxInt16 {
		return math.MaxInt16
	}
	if val < math.MinInt16 {
		return math.MinInt16
	}
	return int16(val)
}

// Float64toUint8 converts float64 to uint8
func Float64toUint8(val float64) uint8 {
	if val > math.MaxUint8 {
		return math.MaxUint8
	}
	if val < 0 {
		return 0
	}
	return uint8(val)
}

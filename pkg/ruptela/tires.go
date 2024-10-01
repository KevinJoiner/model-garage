package ruptela

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

const (
	// TirePressureScaling is the scaling factor for tire pressure (0.5 kPa/bit)
	TirePressureScaling = 0.5

	// TireTemperatureScaling is the scaling factor for tire temperature (0.03125 °C/bit)
	TireTemperatureScaling = 0.03125

	// TireTemperatureOffset is the temperature offset (-273°C) applied to tire temperature values
	TireTemperatureOffset = -273

	// TireStatusBits is the number of bits used to represent the tire status
	TireStatusBits = 2

	// TireSensorEnableBits is the number of bits used to represent the tire sensor enable status
	TireSensorEnableBits = 2

	// TireSensorElectricalBits is the number of bits used to represent the tire sensor electrical fault status
	TireSensorElectricalBits = 2

	// ExtendedTirePressureBits is the number of bits used to represent extended tire pressure support
	ExtendedTirePressureBits = 2

	// TirePressureThresholdBits is the number of bits used to represent the tire pressure threshold detection
	TirePressureThresholdBits = 3

	// TireAirLeakageScaling is the scaling factor for tire air leakage rate (0.03125 kPa/bit)
	TireAirLeakageScaling = 0.03125

	// TirePressureOffsetCelsius is the temperature offset applied to tire pressure threshold (-40°C)
	TirePressureOffsetCelsius = -40

	// LocationBytePos is the byte position for tire location.
	LocationBytePos = 0
	// PressureBytePos is the byte position for tire pressure.
	PressureBytePos = 1
	// TemperatureBytePos is the starting byte position for tire temperature.
	TemperatureBytePos = 2
	// StatusBytePos is the byte position for status flags.
	StatusBytePos = 4
	// AirLeakageRateBytePos is the starting byte position for tire air leakage rate.
	AirLeakageRateBytePos = 5
	// PressureThresholdBytePos is the byte position for tire pressure threshold detection.
	PressureThresholdBytePos = 7

	// FrontLeftTire value for front left tire location.
	FrontLeftTire = 0b0000
	// FrontRightTire value for front right tire location.
	FrontRightTire = 0b0001
	// RearLeftTire value for rear left tire location.
	RearLeftTire = 0b0100
	// RearRightTire value for rear right tire location.
	RearRightTire = 0b0101
)

// TireInfo represents the tire information extracted from the PGN 65268 message.
type TireInfo struct {
	// TireLocation is the location of the tire.
	//  The low order 4 bits represent a position number, counting left to right when facing in the direction of normal vehicle travel (forward).
	//  The high order 4 bits represent a position number, counting front to back on the vehicle.
	// The value 0xFF indicates not available.
	TireLocation uint8

	// TirePressure is the tire pressure in kPa.
	TirePressure float64

	// TireTemperature is the temperature at the surface of the tire sidewall in °C.
	TireTemperature float64

	// TireSensorEnableStatus indicates whether the tire sensor is being monitored by the controller or is Enabled/Disabled.
	// 00 Off / Isolated / Disabled
	// 01 On (tire is polled) / Enabled
	// 10 Not Defined
	// 11 Not Supported
	TireSensorEnableStatus uint8

	// TireStatus is the tire sensor enable status.
	// 00 Ok (no fault)
	// 01 Tire leak detected
	// 10 Error
	// 11 Not Supported
	TireStatus uint8

	// TireSensorElectricalFault indicates the status of electrical fault on the tire sensor.
	// 00 Ok (No Fault)
	// 01 Not Defined (Fault)
	// 10 Error
	// 11 Not Supported
	TireSensorElectricalFault uint8

	// ExtendedTirePressureSupport defines the choice of using the tire pressure (PGN 65268) or Extended Tire Pressure (PGN 64578). The state value '01' indicates the extended tire pressure SPN is used for the pressure data. Any other value indicates the extended tire pressure SPN is not being used.
	// 00 - Not Using Extended Tire Pressure
	// 01 - Using Extended Tire Pressure
	// 10 - Error
	// 11 - Not Available/Not Supported
	ExtendedTirePressureSupport uint8

	// TireAirLeakageRate is the pressure loss rate from the tire in kPa/h.
	TireAirLeakageRate float64

	// TirePressureThresholdDetection signal indicataes the pressure level of the tire.  The levels defined represent different pressure conditions of the tire:
	// 000 Extreme over pressure - The tire pressure is at a level where the safety of the vehicle may be jeopardised.
	// 001 Over pressure - The tire pressure is higher than the pressure defined by the vehicle or tire manufacturer.
	// 010 No warning pressure - The tire pressure is within the thresholds defined by the vehicle or tire manufacturer.
	// 011 Under pressure - The tire pressure is lower than the pressure defined by the vehicle or tire manufacturer.
	// 100 Extreme under pressure - The tire pressure is at a level where the safety of the vehicle may be jeopardised.
	// 101 Not defined
	// 110 Error indicator
	// 111 Not available
	TirePressureThresholdDetection uint8
}

// ParseTireInfoFromHex parses a hex string into the TireInfo struct.
func ParseTireInfoFromHex(hexStr string) (TireInfo, error) {
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return TireInfo{}, fmt.Errorf("invalid hex string: %v", err)
	}
	return ParseTireInfo(data), nil
}

// ParseTireInfo parses an 8-byte PGN message into the TireInfo struct.
// The input is expected to be an 8-byte slice that represents the PGN 65268 message.
func ParseTireInfo(data []byte) TireInfo {
	if len(data) < 8 {
		panic("Invalid data length, expected exactly 8 bytes")
	}

	// SPN 929: Tire Location (1 byte)
	tireLocation := data[LocationBytePos]

	// SPN 241: Tire Pressure (1 byte, scaled by 0.5 kPa/bit)
	tirePressure := float64(data[PressureBytePos]) * TirePressureScaling

	// SPN 242: Tire Temperature (2 bytes, scaled by 0.03125 °C/bit with -273°C offset)
	tireTemperature := float64(int16(binary.BigEndian.Uint16(data[TemperatureBytePos:TemperatureBytePos+2])))*TireTemperatureScaling + TireTemperatureOffset

	// SPN 1699, 1698, 1697, 6990: Status flags (2 bits each in 5th byte)
	byte5 := data[StatusBytePos]
	tireSensorEnableStatus := (byte5 >> 6) & 0x03    // First 2 bits of byte 5
	tireStatus := (byte5 >> 4) & 0x03                // Next 2 bits of byte 5
	tireSensorElectricalFault := (byte5 >> 2) & 0x03 // Next 2 bits of byte 5
	extendedTirePressureSupport := byte5 & 0x03      // Last 2 bits of byte 5

	// SPN 2586: Tire Air Leakage Rate (2 bytes, scaled by 0.03125 kPa/bit)
	tireAirLeakageRate := float64(int16(binary.BigEndian.Uint16(data[AirLeakageRateBytePos:AirLeakageRateBytePos+2]))) * TireAirLeakageScaling

	// SPN 2587: Tire Pressure Threshold Detection (3 bits in the 8th byte)
	byte8 := data[PressureThresholdBytePos]
	tirePressureThresholdDetection := (byte8 >> 5) & 0x07 // First 3 bits of byte 8

	return TireInfo{
		TireLocation:                   tireLocation,
		TirePressure:                   tirePressure*TirePressureScaling + TirePressureOffsetCelsius,
		TireTemperature:                tireTemperature*TireTemperatureScaling + TireTemperatureOffset,
		TireSensorEnableStatus:         tireSensorEnableStatus,
		TireStatus:                     tireStatus,
		TireSensorElectricalFault:      tireSensorElectricalFault,
		ExtendedTirePressureSupport:    extendedTirePressureSupport,
		TireAirLeakageRate:             tireAirLeakageRate * TireAirLeakageScaling,
		TirePressureThresholdDetection: tirePressureThresholdDetection,
	}
}

func ToChassisAxleRow1WheelLeftTirePressure(val string) (float64, error) {
	tireInfo, err := ParseTireInfoFromHex(val)
	if err != nil {
		return 0, err
	}
	if tireInfo.TireLocation != FrontLeftTire {
		return 0, errNotFound
	}
	return tireInfo.TirePressure, nil
}

func ToChassisAxleRow1WheelRightTirePressure(val string) (float64, error) {
	tireInfo, err := ParseTireInfoFromHex(val)
	if err != nil {
		return 0, err
	}
	if tireInfo.TireLocation != FrontRightTire {
		return 0, errNotFound
	}
	return tireInfo.TirePressure, nil
}

func ToChassisAxleRow2WheelLeftTirePressure(val string) (float64, error) {
	tireInfo, err := ParseTireInfoFromHex(val)
	if err != nil {
		return 0, err
	}
	if tireInfo.TireLocation != RearLeftTire {
		return 0, errNotFound
	}
	return tireInfo.TirePressure, nil
}

func ToChassisAxleRow2WheelRightTirePressure(val string) (float64, error) {
	tireInfo, err := ParseTireInfoFromHex(val)
	if err != nil {
		return 0, err
	}
	if tireInfo.TireLocation != RearRightTire {
		return 0, errNotFound
	}

	return tireInfo.TirePressure, nil
}

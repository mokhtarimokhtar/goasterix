package bdscode

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestBDSCode60Decode_NoFields(t *testing.T) {
	// Arrange
	input := "7F F7 FE FF DF FB FF"
	input = strings.ReplaceAll(input, " ", "")
	outputMagneticHeading := int16(0)
	outputIndicatedAirspeed := uint16(0)
	outputMach := 0.0
	outputBarometricAltitudeRate := int16(0)
	outputInertialVerticalVelocity := int16(0)
	tmp, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var data [7]byte
	copy(data[:], tmp) // convert slice to array for parameter function decode
	code60 := new(Code60)

	// Act
	err = code60.Decode(data)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if code60.MagneticHeadingStatus != false {
		t.Errorf("FAIL: MagneticHeadingStatus: %v; Expected: %v", code60.MagneticHeadingStatus, false)
	} else {
		t.Logf("SUCCESS: MagneticHeadingStatus: %v; Expected: %v", code60.MagneticHeadingStatus, false)
	}
	if code60.MagneticHeading != outputMagneticHeading {
		t.Errorf("FAIL: MagneticHeading: %v; Expected: %v", code60.MagneticHeading, outputMagneticHeading)
	} else {
		t.Logf("SUCCESS: MagneticHeading: %v; Expected: %v", code60.MagneticHeading, outputMagneticHeading)
	}

	if code60.IndicatedAirspeedStatus != false {
		t.Errorf("FAIL: IndicatedAirspeedStatus: %v; Expected: %v", code60.IndicatedAirspeedStatus, false)
	} else {
		t.Logf("SUCCESS: IndicatedAirspeedStatus: %v; Expected: %v", code60.IndicatedAirspeedStatus, false)
	}
	if code60.IndicatedAirspeed != outputIndicatedAirspeed {
		t.Errorf("FAIL: IndicatedAirspeed: %v; Expected: %v", code60.IndicatedAirspeed, outputIndicatedAirspeed)
	} else {
		t.Logf("SUCCESS: IndicatedAirspeed: %v; Expected: %v", code60.IndicatedAirspeed, outputIndicatedAirspeed)
	}

	if code60.MachStatus != false {
		t.Errorf("FAIL: MachStatus: %v; Expected: %v", code60.MachStatus, false)
	} else {
		t.Logf("SUCCESS: MachStatus: %v; Expected: %v", code60.MachStatus, false)
	}
	if code60.Mach != outputMach {
		t.Errorf("FAIL: Mach: %v; Expected: %v", code60.Mach, outputMach)
	} else {
		t.Logf("SUCCESS: Mach: %v; Expected: %v", code60.Mach, outputMach)
	}

	if code60.BarometricAltitudeRateStatus != false {
		t.Errorf("FAIL: BarometricAltitudeRateStatus: %v; Expected: %v", code60.BarometricAltitudeRateStatus, false)
	} else {
		t.Logf("SUCCESS: BarometricAltitudeRateStatus: %v; Expected: %v", code60.BarometricAltitudeRateStatus, false)
	}
	if code60.BarometricAltitudeRate != outputBarometricAltitudeRate {
		t.Errorf("FAIL: BarometricAltitudeRate: %v; Expected: %v", code60.BarometricAltitudeRate, outputBarometricAltitudeRate)
	} else {
		t.Logf("SUCCESS: BarometricAltitudeRate: %v; Expected: %v", code60.BarometricAltitudeRate, outputBarometricAltitudeRate)
	}

	if code60.InertialVerticalVelocityStatus != false {
		t.Errorf("FAIL: InertialVerticalVelocityStatus: %v; Expected: %v", code60.InertialVerticalVelocityStatus, false)
	} else {
		t.Logf("SUCCESS: InertialVerticalVelocityStatus: %v; Expected: %v", code60.InertialVerticalVelocityStatus, false)
	}
	if code60.InertialVerticalVelocity != outputInertialVerticalVelocity {
		t.Errorf("FAIL: InertialVerticalVelocity: %v; Expected: %v", code60.InertialVerticalVelocity, outputInertialVerticalVelocity)
	} else {
		t.Logf("SUCCESS: InertialVerticalVelocity: %v; Expected: %v", code60.InertialVerticalVelocity, outputInertialVerticalVelocity)
	}
}

func TestBDSCode60Decode_AllFields(t *testing.T) {
	// Arrange
	input := "C0 FC 0F 8F 30 F6 0F"
	input = strings.ReplaceAll(input, " ", "")
	outputMagneticHeading := int16(-177)
	outputIndicatedAirspeed := uint16(519)
	outputMach := 2.288
	outputBarometricAltitudeRate := int16(-482 * 32)
	outputInertialVerticalVelocity := int16(-497 * 32)

	tmp, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var data [7]byte
	copy(data[:], tmp) // convert slice to array for parameter function decode
	code60 := new(Code60)

	// Act
	err = code60.Decode(data)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if code60.MagneticHeadingStatus != true {
		t.Errorf("FAIL: MagneticHeadingStatus: %v; Expected: %v", code60.MagneticHeadingStatus, true)
	} else {
		t.Logf("SUCCESS: MagneticHeadingStatus: %v; Expected: %v", code60.MagneticHeadingStatus, true)
	}

	if code60.MagneticHeading != outputMagneticHeading {
		t.Errorf("FAIL: MagneticHeading: %v; Expected: %v", &code60.MagneticHeading, outputMagneticHeading)
	} else {
		t.Logf("SUCCESS: MagneticHeading: %v; Expected: %v", code60.MagneticHeading, outputMagneticHeading)
	}

	if code60.IndicatedAirspeedStatus != true {
		t.Errorf("FAIL: IndicatedAirspeedStatus: %v; Expected: %v", code60.IndicatedAirspeedStatus, true)
	} else {
		t.Logf("SUCCESS: IndicatedAirspeedStatus: %v; Expected: %v", code60.IndicatedAirspeedStatus, true)
	}
	if code60.IndicatedAirspeed != outputIndicatedAirspeed {
		t.Errorf("FAIL: IndicatedAirspeed: %v; Expected: %v", code60.IndicatedAirspeed, outputIndicatedAirspeed)
	} else {
		t.Logf("SUCCESS: IndicatedAirspeed: %v; Expected: %v", code60.IndicatedAirspeed, outputIndicatedAirspeed)
	}

	if code60.MachStatus != true {
		t.Errorf("FAIL: MachStatus: %v; Expected: %v", code60.MachStatus, true)
	} else {
		t.Logf("SUCCESS: MachStatus: %v; Expected: %v", code60.MachStatus, true)
	}
	if code60.Mach != outputMach {
		t.Errorf("FAIL: Mach: %v; Expected: %v", code60.Mach, outputMach)
	} else {
		t.Logf("SUCCESS: Mach: %v; Expected: %v", code60.Mach, outputMach)
	}

	if code60.BarometricAltitudeRateStatus != true {
		t.Errorf("FAIL: BarometricAltitudeRateStatus: %v; Expected: %v", code60.BarometricAltitudeRateStatus, true)
	} else {
		t.Logf("SUCCESS: BarometricAltitudeRateStatus: %v; Expected: %v", code60.BarometricAltitudeRateStatus, true)
	}
	if code60.BarometricAltitudeRate != outputBarometricAltitudeRate {
		t.Errorf("FAIL: BarometricAltitudeRate: %v; Expected: %v", code60.BarometricAltitudeRate, outputBarometricAltitudeRate)
	} else {
		t.Logf("SUCCESS: BarometricAltitudeRate: %v; Expected: %v", code60.BarometricAltitudeRate, outputBarometricAltitudeRate)
	}

	if code60.InertialVerticalVelocityStatus != true {
		t.Errorf("FAIL: InertialVerticalVelocityStatus: %v; Expected: %v", code60.InertialVerticalVelocityStatus, true)
	} else {
		t.Logf("SUCCESS: InertialVerticalVelocityStatus: %v; Expected: %v", code60.InertialVerticalVelocityStatus, true)
	}
	if code60.InertialVerticalVelocity != outputInertialVerticalVelocity {
		t.Errorf("FAIL: InertialVerticalVelocity: %v; Expected: %v", code60.InertialVerticalVelocity, outputInertialVerticalVelocity)
	} else {
		t.Logf("SUCCESS: InertialVerticalVelocity: %v; Expected: %v", code60.InertialVerticalVelocity, outputInertialVerticalVelocity)
	}
}

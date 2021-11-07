package model

import (
	"encoding/json"
	"errors"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"reflect"
	"testing"
)

func Test_Model_Cat048_to_JSON_Record(t *testing.T) {
	// Arrange
	// bds 02 e79a5d27a00c00 60 a3280030a40000 40
	input := "fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5"
	output := []byte(`{"sourceIdentifier":{"sac":8,"sic":54},"aircraftAddress":"490D01","aircraftIdentification":"NJE834H","timeOfDay":34102.640625,"rhoTheta":{"rho":148.77734375,"theta":2.1174999999999997},"flightLevel":{"v":"code_validated","g":"default","level":180},"radarPlotCharacteristics":{"srr":2,"sam":-73},"mode3ACode":{"squawk":"4423","v":"code_validated","g":"default","l":"code_derived_from_transponder"},"trackNumber":1594,"trackVelocity":{"groundSpeed":0.113464065,"heading":290.5485},"trackStatus":{"cnf":"confirmed_track","rad":"ssr_modes_track","dou":"normal_confidence","mah":"no_horizontal_man_sensed","cdm":"maintaining"},"bdsRegisterData":[{"transponderRegisterNumber":"60","code60":{"magneticHeading":-68,"indicatedAirspeed":302,"mach":0.632,"barometricAltitudeRate":32}},{"transponderRegisterNumber":"40","code40":{"mcpSelectAltitude":18000,"barometricPressureSetting":1013}}],"comAcasCapabilityFlightStatus":{"com":"comm_a_and_comm_b_capability","stat":"no_alert_no_spi_aircraft_airborne","si":"si_code_capable","mssc":"yes","arc":"25_ft_resolution","aic":"yes","b1a":"1","b1b":"5"}}`)

	uap048 := uap.Cat048V127
	data := goasterix.HexStringToByte(input)
	rec := new(goasterix.Record)
	_, err := rec.Decode(data, uap048)

	cat048Model := new(Cat048Model)
	cat048Model.write(rec.Items)

	// Act
	recJson, _ := json.Marshal(cat048Model)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(recJson, output) == false {
		t.Errorf("FAIL: %s; \nExpected: %s", recJson, output)
	} else {
		t.Logf("SUCCESS: %s; Expected: %s", recJson, output)
	}
}

func Test_Model_Cat048_RhoTheta(t *testing.T) {
	// Arrange
	input := [4]byte{0xFF, 0xFF, 0xFF, 0xFF}
	output := PolarPosition{Rho: float64(0xFFFF) / 256, Theta: float64(0xFFFF) * 0.0055}

	// Act
	res, err := rhoTheta(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}

}

func Test_Model_Cat048_Mode3ACodeVGL(t *testing.T) {
	// Arrange
	input := [2]byte{0x1F, 0xFF}
	output := Mode3A{
		Squawk: "7777", // FFF = 111 111 111 111 = 7777
		V:      "code_validated",
		G:      "default",
		L:      "code_derived_from_transponder",
	}
	// Act
	res, err := mode3ACodeVGL(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}

}

func Test_Model_Cat048_FlightLevel(t *testing.T) {
	// Arrange
	input := [2]byte{0x3F, 0xFF}
	output := FL{
		V:     "code_validated",
		G:     "default",
		Level: uint16(0x3FFF) / 4,
	}

	// Act
	res, err := flightLevel(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}

}

func Test_Model_Cat048_ModeSIdentification_valid_char(t *testing.T) {
	// Arrange
	// AFR1234
	// 1 		6 		18 		49 		50 		51 		52		_ 		= IA5
	// 000001 	000110 	010010	110001	110010	110011	110100	100000
	// 0000-0100 0110-0100 1011-0001 1100-1011 0011-1101 0010-0000
	// 0x04 	 0x64      0xB1      0xCB      0x3D      0x20
	input := [6]byte{0x04, 0x64, 0xB1, 0xCB, 0x3D, 0x20}
	output := "AFR1234"

	// Act
	s, err := modeSIdentification(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if s != output {
		t.Errorf("FAIL: s = %s; Expected: %s", s, output)
	} else {
		t.Logf("SUCCESS: s = %s; Expected: %s", s, output)
	}
}

func Test_Model_Cat048_ModeSIdentification_invalid_char(t *testing.T) {
	// Arrange
	input := [6]byte{0x04, 0x64, 0xB1, 0xCB, 0x3D, 0x3A}
	output := "AFR1234"

	// Act
	s, err := modeSIdentification(input)

	// Assert
	if errors.Is(err, ErrCharUnknown) == true {
		t.Errorf("FAIL: error: %v; Expected: %v: %X", err, ErrCharUnknown, input)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v: %X", err, ErrCharUnknown, input)
	}

	if s != output {
		t.Errorf("FAIL: s = %s; Expected: %s", s, output)
	} else {
		t.Logf("SUCCESS: s = %s; Expected: %s", s, output)
	}
}

func Test_Model_Cat048_TrackNumber(t *testing.T) {
	// Arrange
	input := [2]byte{0x0F, 0xFF}
	output := uint16(4095)

	// Act
	res, err := trackNumber(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}

}

func Test_Model_Cat048_CartesianXY(t *testing.T) {
	// Arrange
	input := [4]byte{0x01, 0x00, 0xFF, 0x00}
	output := CartesianXYPosition{
		X: float64(int16(256)) / 128,
		Y: float64(int16(-256)) / 128,
	}

	// Act
	res, err := cartesianXY(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_Model_Cat048_TrackVelocity(t *testing.T) {
	// Arrange
	input := [4]byte{0x07, 0xc3, 0xdf, 0xc6}
	output := Velocity{
		GroundSpeed: 0.121276545,
		Heading:     315.073,
	}

	// Act
	res, err := trackVelocity(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func Test_Model_Cat048_TrackStatus(t *testing.T) {
	// Arrange
	input := []byte{0x40}
	output := Status{
		CNF: "confirmed_track",
		RAD: "ssr_modes_track",
		DOU: "normal_confidence",
		MAH: "no_horizontal_man_sensed",
		CDM: "maintaining",
	}

	// Act
	res, err := trackStatus(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

package transform

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestModelSacSic(t *testing.T) {
	// Arrange
	input := [2]byte{0xFF, 0xFF}
	output := SourceIdentifier{Sac: 0xFF, Sic: 0xFF}

	// Act
	res, err := sacSic(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: % X; Expected: % X", res, output)
	} else {
		t.Logf("SUCCESS: % X; Expected: % X", res, output)
	}

}

func TestModelSacSic_JSON(t *testing.T) {
	// Arrange
	input := [2]byte{0xFF, 0xFF}
	output := []byte(`{"sac":255,"sic":255}`)
	res, err := sacSic(input)

	// Act
	sacSicJson, _ := json.Marshal(res)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if bytes.Equal(sacSicJson, output) == false {
		t.Errorf("FAIL: %s; Expected: %s", sacSicJson, output)
	} else {
		t.Logf("SUCCESS: %s; Expected: %s", sacSicJson, output)
	}
}

func TestModelTimeOfDay(t *testing.T) {
	// Arrange
	input := [3]byte{0xa8, 0xbf, 0x38}
	output := float64(0x00a8bf38) / 128

	// Act
	res, err := timeOfDay(input)

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

func TestModelTimeOfDayHighPrecision(t *testing.T) {
	// Arrange
	input := [4]byte{0x3a, 0xda, 0xb9, 0xf5}
	output := TimeOfDayHighPrecision{
		FSI:             "+0",               // FSI = 00
		TimeOfReception: 0.9195999996736646, // float64(0x3adab9f5) / 2^30
	}
	// Act
	res, err := timeOfDayHighPrecision(input)

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

func TestModeSIdentification(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        [6]byte
		output       string
		err          error
	}
	dataset := []dataTest{
		{
			TestCaseName: "valid chars",
			input:        [6]byte{0x04, 0x64, 0xB1, 0xCB, 0x3D, 0x20},
			output:       "AFR1234 ",
			err:          nil,
		},
		{
			TestCaseName: "chars 6 unknown",
			input:        [6]byte{0x04, 0x64, 0xB1, 0xCB, 0x3D, 0x3A},
			output:       "AFR1234",
			err:          ErrCharUnknown,
		},
		{
			TestCaseName: "chars unknown",
			input:        [6]byte{},
			output:       "",
			err:          ErrCharUnknown,
		},
	}

	for _, row := range dataset {
		// Arrange
		// Act
		s, err := modeSIdentification(row.input)

		// Assert
		if err != row.err {
			//if errors.Is(err, row.err) {
			t.Errorf("FAIL: error: %v; Expected: %v", err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}

		if s != row.output {
			t.Errorf("FAIL: s = %s; Expected: %s", s, row.output)
		} else {
			t.Logf("SUCCESS: s = %s; Expected: %s", s, row.output)
		}
	}
}

func TestTrackNumber(t *testing.T) {
	// Arrange
	input := [2]byte{0x0F, 0xFF}
	output := uint16(4095)

	// Act
	res := trackNumber(input)

	// Assert
	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}

}

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

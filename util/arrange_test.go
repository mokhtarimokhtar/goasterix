package util

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestHexStringToByte_Valid(t *testing.T) {
	// Arrange
	input := "01 0203 04"
	output := []byte{0x01, 0x02, 0x03, 0x04}

	// Act
	data, _ := HexStringToByte(input)

	// Assert
	if bytes.Equal(data, output) == false {
		t.Errorf("MsgFailInValue: data = % X; Expected: % X", data, output)
	} else {
		t.Logf("MsgSuccessInValue: data = % X; Expected: % X", data, output)
	}
}

func TestHexStringToByte_Empty(t *testing.T) {
	// Arrange
	input := ""
	var output []byte

	// Act
	data, err := HexStringToByte(input)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if bytes.Equal(data, output) == false {
		t.Errorf("MsgFailInValue: data = % X; Expected: % X", data, output)
	} else {
		t.Logf("MsgSuccessInValue: data = % X; Expected: % X", data, output)
	}
}

func TestHexStringToByte_Error(t *testing.T) {
	// Arrange
	input := "01 0203 0"
	var output []byte

	// Act
	data, err := HexStringToByte(input)

	// Assert
	if err != hex.ErrLength {
		t.Errorf("MsgFailInValue: error: %s; Expected: %v", err, hex.ErrLength)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, hex.ErrLength)
	}
	if bytes.Equal(data, output) == false {
		t.Errorf("MsgFailInValue: % X; Expected: % X", data, output)
	} else {
		t.Logf("MsgSuccessInValue: % X; Expected: % X", data, output)
	}
}

func TestCleanStringMultiline(t *testing.T) {
	// Arrange
	input := `
			string1
			string2  
				string3
			`
	output := "string1string2string3"

	// Act
	data := CleanStringMultiline(input)

	// Assert
	if data != output {
		t.Errorf("MsgFailInValue: result: %s - Expected: %s", data, output)
	} else {
		t.Logf("MsgSuccessInValue: result: %s - Expected: %s", data, output)
	}
}

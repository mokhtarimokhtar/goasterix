package goasterix

import (
	"bytes"
	"testing"
)

func Test_TwoComplement16_positive_number(t *testing.T)  {
	// Arrange
	input := uint16(0x010F) // 0000 0001 0000 1111
	size := uint8(10)		// ---- --01 0000 1111  -> tenth bit
	output := int16(271)   // 01 0000 1111 = 271

	// Act
	result := TwoComplement16(size, input)

	// Assert
	if result != output {
		t.Errorf("FAIL: result = %v; Expected: %v", result, output)
	} else {
		t.Logf("SUCCESS: result = %v; Expected: %v", result, output)
	}
}

func Test_TwoComplement16_negative_number(t *testing.T)  {
	// Arrange
	input := uint16(0x040F) // 0000 0100 0000 1111
	size := uint8(11)		// ---- -100 0000 1111  -> tenth bit
	output := int16(-1009)  // ---- -011 1111 0001 = -1009

	// Act
	result := TwoComplement16(size, input)

	// Assert
	if result != output {
		t.Errorf("FAIL: result = %v; Expected: %v", result, output)
	} else {
		t.Logf("SUCCESS: result = %v; Expected: %v", result, output)
	}
}

func Test_StringToHex_valid(t *testing.T) {
	// Arrange
	input := "01 0203 04"
	output := []byte{0x01, 0x02, 0x03, 0x04}

	// Act
	data := StringToHex(input)

	// Assert
	if bytes.Equal(data, output) == false {
		t.Errorf("FAIL: data = % X; Expected: % X", data, output)
	} else {
		t.Logf("SUCCESS: data = % X; Expected: % X", data, output)
	}
}

func Test_StringToHex_empty(t *testing.T) {
	// Arrange
	input := ""
	var output []byte

	// Act
	data := StringToHex(input)

	// Assert
	if bytes.Equal(data, output) == false {
		t.Errorf("FAIL: data = % X; Expected: % X", data, output)
	} else {
		t.Logf("SUCCESS: data = % X; Expected: % X", data, output)
	}
}

package goasterix

import "testing"

func TestTwoComplement16_PositiveNumber(t *testing.T) {
	// Arrange
	input := uint16(0x010F) // 0000 0001 0000 1111
	size := uint8(10)       // ---- --01 0000 1111  -> tenth bit
	output := int16(271)    // 01 0000 1111 = 271

	// Act
	result := TwoComplement16(size, input)

	// Assert
	if result != output {
		t.Errorf("FAIL: result = %v; Expected: %v", result, output)
	} else {
		t.Logf("SUCCESS: result = %v; Expected: %v", result, output)
	}
}

func TestTwoComplement16_NegativeNumber(t *testing.T) {
	// Arrange
	input := uint16(0x040F) // 0000 0100 0000 1111
	size := uint8(11)       // ---- -100 0000 1111  -> tenth bit
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

func TestTwoComplement32_PositiveNumber(t *testing.T) {
	// Arrange
	input := uint32(0x0007EE0F) // 0000 0000 0000 0111 1110 1110 0000 1111
	size := uint8(20)           // 	   ---- ---- ---- 0111 1110 1110 0000 1111  -> twentieth bit
	output := int32(519695)

	// Act
	result := TwoComplement32(size, input)

	// Assert
	if result != output {
		t.Errorf("FAIL: result = %v; Expected: %v", result, output)
	} else {
		t.Logf("SUCCESS: result = %v; Expected: %v", result, output)
	}
}

func TestTwoComplement32_NegativeNumber(t *testing.T) {
	// Arrange
	input := uint32(0x000FEE0F) // 0000 0000 0000 0111 1110 1110 0000 1111
	size := uint8(20)           // 	   ---- ---- ---- 0111 1110 1110 0000 1111  -> twentieth bit
	output := int32(-4593)

	// Act
	result := TwoComplement32(size, input)

	// Assert
	if result != output {
		t.Errorf("FAIL: result = %v; Expected: %v", result, output)
	} else {
		t.Logf("SUCCESS: result = %v; Expected: %v", result, output)
	}
}

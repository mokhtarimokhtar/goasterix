package model

import (
	"testing"
)

func Test_TableAI5(t *testing.T) {
	// Arrange
	//A - Z = 1 - 26, 0 - 9 = 48 - 57, _ :  32
	tableSixBit := map[uint8]string{
		0x01: "A",
		0x02: "B",
		0x03: "C",
		0x04: "D",
		0x05: "E",
		0x06: "F",
		0x07: "G",
		0x08: "H",
		0x09: "I",
		0x0A: "J",
		0x0B: "K",
		0x0C: "L",
		0x0D: "M",
		0x0E: "N",
		0x0F: "O",
		0x10: "P",
		0x11: "Q",
		0x12: "R",
		0x13: "S",
		0x14: "T",
		0x15: "U",
		0x16: "V",
		0x17: "W",
		0x18: "X",
		0x19: "Y",
		0x1A: "Z",
		0x20: " ",
		0x30: "0",
		0x31: "1",
		0x32: "2",
		0x33: "3",
		0x34: "4",
		0x35: "5",
		0x36: "6",
		0x37: "7",
		0x38: "8",
		0x39: "9",
	}
	for i, ch := range tableSixBit {
		// Act
		s, found := TableIA5[i]

		// Assert
		if found != true {
			t.Errorf("FAIL: found: %v; Expected: %v", found, false)
		} else {
			t.Logf("SUCCESS: found: %v; Expected: %v", found, true)
		}
		if s != ch {
			t.Errorf("FAIL: Char = %s; Expected: %s", s, ch)
		} else {
			t.Logf("SUCCESS: Char = %s; Expected: %s", s, ch)
		}
	}

}

func Test_TableAI5_invalid(t *testing.T) {
	// Arrange
	// A - Z = 1 - 26, 0 - 9 = 48 - 57, _ :  32
	tableInvalidSixBit := []uint8{0, 27, 45, 58, 31, 33}
	output := "" // empty string

	for _, ch := range tableInvalidSixBit {

		// Act
		s, found := TableIA5[ch]

		// Assert
		if found != false {
			t.Errorf("FAIL: found: %v; Expected: %v", found, false)
		} else {
			t.Logf("SUCCESS: found: %v; Expected: %v", found, false)
		}

		if s != output {
			t.Errorf("FAIL: s = %s; Expected: %s", s, output)
		} else {
			t.Logf("SUCCESS: s = %s; Expected: %s", s, output)
		}
	}

}

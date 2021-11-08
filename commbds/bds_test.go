package commbds

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestBDSDecode_MsgCommBCode60(t *testing.T) {
	// Arrange
	input := "FF FF FF FF FF FF FF 60"
	input = strings.ReplaceAll(input, " ", "")
	output := "60"
	tmp, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var data [8]byte
	copy(data[:], tmp) // convert slice to array
	ds := new(Bds)

	// Act
	err = ds.Decode(data)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if ds.TransponderRegisterNumber != output {
		t.Errorf("FAIL: TransponderRegisterNumber: %s; Expected: %s", ds.TransponderRegisterNumber, output)
	} else {
		t.Logf("SUCCESS: TransponderRegisterNumber: %s; Expected: %s", ds.TransponderRegisterNumber, output)
	}
}

func TestBDSDecode_MsgCommBCode40(t *testing.T) {
	// Arrange
	input := "FF FF FF FF FF FF FF 40"
	input = strings.ReplaceAll(input, " ", "")
	output := "40"
	tmp, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var data [8]byte
	copy(data[:], tmp) // convert slice to array
	ds := new(Bds)

	// Act
	err = ds.Decode(data)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if ds.TransponderRegisterNumber != output {
		t.Errorf("FAIL: TransponderRegisterNumber: %s; Expected: %s", ds.TransponderRegisterNumber, output)
	} else {
		t.Logf("SUCCESS: TransponderRegisterNumber: %s; Expected: %s", ds.TransponderRegisterNumber, output)
	}
}

func TestBDSDecode_MsgCommBCode50(t *testing.T) {
	// Arrange
	input := "FF FF FF FF FF FF FF 50"
	input = strings.ReplaceAll(input, " ", "")
	output := "50"
	tmp, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var data [8]byte
	copy(data[:], tmp) // convert slice to array
	ds := new(Bds)

	// Act
	err = ds.Decode(data)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if ds.TransponderRegisterNumber != output {
		t.Errorf("FAIL: TransponderRegisterNumber: %s; Expected: %s", ds.TransponderRegisterNumber, output)
	} else {
		t.Logf("SUCCESS: TransponderRegisterNumber: %s; Expected: %s", ds.TransponderRegisterNumber, output)
	}
}

func TestBDSDecode_MsgCommBCode00(t *testing.T) {
	// Arrange
	input := "FF FF FF FF FF FF FF 00"
	input = strings.ReplaceAll(input, " ", "")
	output := "0"
	outputCode00 := "Not valid"
	tmp, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var data [8]byte
	copy(data[:], tmp) // convert slice to array
	ds := new(Bds)

	// Act
	err = ds.Decode(data)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if ds.TransponderRegisterNumber != output {
		t.Errorf("FAIL: TransponderRegisterNumber: %s; Expected: %s", ds.TransponderRegisterNumber, output)
	} else {
		t.Logf("SUCCESS: TransponderRegisterNumber: %s; Expected: %s", ds.TransponderRegisterNumber, output)
	}
	if *ds.Code00 != outputCode00 {
		t.Errorf("FAIL: Code00: %s; Expected: %s", *ds.Code00, outputCode00)
	} else {
		t.Logf("SUCCESS: Code00: %s; Expected: %s", *ds.Code00, outputCode00)
	}
}

func TestBDSDecode_MsgCommBCode_Other(t *testing.T) {
	// Arrange
	input := "FF FF FF FF FF FF FF FF"
	input = strings.ReplaceAll(input, " ", "")
	output := "ff"
	outputUndefined := "FFFFFFFFFFFFFF"
	tmp, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var data [8]byte
	copy(data[:], tmp) // convert slice to array
	ds := new(Bds)

	// Act
	err = ds.Decode(data)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if ds.TransponderRegisterNumber != output {
		t.Errorf("FAIL: TransponderRegisterNumber: %s; Expected: %s", ds.TransponderRegisterNumber, output)
	} else {
		t.Logf("SUCCESS: TransponderRegisterNumber: %s; Expected: %s", ds.TransponderRegisterNumber, output)
	}

	if *ds.CodeNotProcessed != outputUndefined {
		t.Errorf("FAIL: CodeNotProcessed: %s; Expected: %s", *ds.CodeNotProcessed, outputUndefined)
	} else {
		t.Logf("SUCCESS: CodeNotProcessed: %s; Expected: %s", *ds.CodeNotProcessed, outputUndefined)
	}
}

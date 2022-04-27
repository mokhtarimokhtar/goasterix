package item

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"testing"
)

func TestFspecReader_Valid(t *testing.T) {
	// Arrange
	input := []byte{0xFF, 0x01, 0xF2, 0xFF}
	output := []byte{0xFF, 0x01, 0xF2}
	rb := bytes.NewReader(input)

	// Act
	fspec, err := FspecReader(rb)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}

	if bytes.Equal(fspec, output) == false {
		t.Errorf("MsgFailInValue: sp = % X; Expected: % X", fspec, output)
	} else {
		t.Logf("MsgSuccessInValue: sp = % X; Expected: % X", fspec, output)
	}
}

func TestFspecReader_Invalid(t *testing.T) {
	// Arrange
	input := []byte{0xFF, 0x01}
	var output []byte
	rb := bytes.NewReader(input)

	// Act
	fspec, err := FspecReader(rb)

	// Assert
	if err != io.EOF {
		t.Errorf("MsgFailInValue: error: %s; Expected: %v", err, io.EOF)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, io.EOF)
	}

	if bytes.Equal(fspec, output) == false {
		t.Errorf("MsgFailInValue: sp = % X; Expected: % X", fspec, output)
	} else {
		t.Logf("MsgSuccessInValue: sp = % X; Expected: % X", fspec, output)
	}
}

func TestFspecIndex(t *testing.T) {
	type testCase struct {
		input  []byte
		output []uint8
	}
	// Arrange
	dataSet := []testCase{
		{input: []byte{0x80}, output: []uint8{1}},
		{input: []byte{0x40}, output: []uint8{2}},
		{input: []byte{0x20}, output: []uint8{3}},
		{input: []byte{0x10}, output: []uint8{4}},
		{input: []byte{0x08}, output: []uint8{5}},
		{input: []byte{0x04}, output: []uint8{6}},
		{input: []byte{0x02}, output: []uint8{7}},
		{input: []byte{0x01}, output: []uint8{}},
		{input: []byte{0x01, 0x80}, output: []uint8{8}},
		{input: []byte{0xfe}, output: []uint8{1, 2, 3, 4, 5, 6, 7}},
		{input: []byte{0xff}, output: []uint8{1, 2, 3, 4, 5, 6, 7}},
		{input: []byte{0xaa}, output: []uint8{1, 3, 5, 7}},
		{input: []byte{0x55}, output: []uint8{2, 4, 6}},
		{input: []byte{}, output: []uint8{}},
		{input: []byte{0xef, 0x98}, output: []uint8{1, 2, 3, 5, 6, 7, 8, 11, 12}},
	}

	for _, row := range dataSet {
		// Act
		frnIndex := FspecIndex(row.input)

		// Assert
		if bytes.Equal(frnIndex, row.output) == false {
			t.Errorf(util.MsgFailInHex, "", frnIndex, row.output)
		} else {
			t.Logf(util.MsgSuccessInHex, "", frnIndex, row.output)
		}
	}

}

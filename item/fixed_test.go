package item

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestFixedReader(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  string
		item   DataItem
		output DataItem
		err    error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:  "testCase 1",
			input: "01 02 03 04 05 06 07 08",
			item: &Fixed{
				Size: 8,
			},
			output: &Fixed{
				Size: 8,
				Data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			},
			err: nil,
		},
		{
			Name:  "testCase 2",
			input: "01 02 03 04 05 06 07 08",
			item: &Fixed{
				Size: 8,
				SubItems: []SubItem{
					&SubItemBit{
						Pos: BitPosition{Bit: 57},
					},
					&SubItemFromTo{
						Pos: BitPosition{From: 56, To: 25},
					},
					&SubItemFromTo{
						Pos: BitPosition{From: 24, To: 1},
					},
				},
			},
			output: &Fixed{
				Size: 8,
				SubItems: []SubItem{
					&SubItemBit{
						Pos:  BitPosition{Bit: 57},
						Data: []byte{0x01},
					},
					&SubItemFromTo{
						Pos:  BitPosition{From: 56, To: 25},
						Data: []byte{0x02, 0x03, 0x04, 0x05},
					},
					&SubItemFromTo{
						Pos:  BitPosition{From: 24, To: 1},
						Data: []byte{0x06, 0x07, 0x08},
					},
				},
			},
			err: nil,
		},
		{
			Name:  "testCase 3",
			input: "01 02 03 04 05 06 07",
			item: &Fixed{
				Size: 8,
			},
			output: &Fixed{
				Size: 8,
				Data: nil,
			},
			err: io.ErrUnexpectedEOF,
		},
	}

	for _, tc := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(tc.input)
		rb := bytes.NewReader(input)
		f := newFixed(tc.item)

		// Act
		err := f.Reader(rb)

		// Assert
		if err != tc.err {
			t.Errorf(util.MsgFailInValue, tc.Name, err, tc.err)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, err, tc.err)
		}

		if reflect.DeepEqual(f, tc.output) == false {
			t.Errorf(util.MsgFailInValue, tc.Name, f, tc.output)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, f, tc.output)
		}
	}
}

func TestFixedString(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Fixed
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Fixed{
				Base: Base{
					FRN:          1,
					DataItemName: "I000/010",
					Description:  "Test item",
				},
				Data: []byte{0xab, 0xcd},
			},
			output: "I000/010:[abcd]",
		},
		{
			Name: "testCase 2",
			input: Fixed{
				Base: Base{
					FRN:          1,
					DataItemName: "I000/010",
					Description:  "Test item",
				},
				SubItems: []SubItem{
					&SubItemFromTo{
						Name: "010-1",
						Pos:  BitPosition{From: 16, To: 9},
						Data: []byte{0xab},
					},
					&SubItemFromTo{
						Name: "010-2",
						Pos:  BitPosition{From: 8, To: 1},
						Data: []byte{0xcd},
					},
					&SubItemBit{
						Name: "010-3",
						Pos:  BitPosition{Bit: 8},
						Data: []byte{0x01},
					},
				},
			},
			output: "I000/010:[010-1:ab][010-2:cd][010-3:01]",
		},
		{
			Name: "testCase 3",
			input: Fixed{
				Base: Base{},
				Data: nil,
			},
			output: ":[]",
		},
	}

	for _, tc := range dataSet {
		// Act
		s := tc.input.String()

		// Assert
		if s != tc.output {
			t.Errorf(util.MsgFailInValue, tc.Name, s, tc.output)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, s, tc.output)
		}
	}
}

func TestFixedPayload(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Fixed
		output []byte
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Fixed{
				Data: []byte{0xab, 0xcd},
			},
			output: []byte{0xab, 0xcd},
		},
		{
			Name: "testCase 2",
			input: Fixed{
				Base: Base{},
				Data: nil,
			},
			output: nil,
		},
	}

	for _, tc := range dataSet {
		// Act
		res := tc.input.Payload()

		// Assert
		if bytes.Equal(res, tc.output) == false {
			t.Errorf(util.MsgFailInHex, tc.Name, res, tc.output)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, res, tc.output)
		}
	}
}

func TestFixedGetCompound(t *testing.T) {
	// Arrange
	input := new(Fixed)
	// Act
	res := input.GetCompound()
	// Assert
	if res != nil {
		t.Errorf(util.MsgFailInValue, "", res, nil)
	} else {
		t.Logf(util.MsgSuccessInValue, "", res, nil)
	}
}

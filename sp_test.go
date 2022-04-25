package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestSpecialPurposeReader(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  string
		item   Item
		output Item
		err    error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:  "testCase 1",
			input: "08 01 02 03 04 05 06 07",
			item:  &SpecialPurpose{},
			err:   nil,
			output: &SpecialPurpose{
				Len:  0x08,
				Data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
			},
		},
		{
			Name:  "testCase 2",
			input: "08 01 02 03 04 05 06",
			item:  &SpecialPurpose{},
			err:   io.ErrUnexpectedEOF,
			output: &SpecialPurpose{
				Len:  0x08,
				Data: nil,
			},
		},
		{
			Name:  "testCase 3",
			input: "",
			item:  &SpecialPurpose{},
			err:   io.EOF,
			output: &SpecialPurpose{
				Len:  0x00,
				Data: nil,
			},
		},
	}

	for _, tc := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(tc.input)
		rb := bytes.NewReader(input)
		f := NewSpecialPurpose(tc.item)

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

func TestSpecialPurposeString(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  SpecialPurpose
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: SpecialPurpose{
				Base: Base{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
				},
				Len:  0x04,
				Data: []byte{0xab, 0xcd, 0xef},
			},
			output: "I000/010:04abcdef",
		},
		{
			Name: "testCase 2",
			input: SpecialPurpose{
				Base: Base{},
				Len:  0,
				Data: nil,
			},
			output: ":00",
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

func TestSpecialPurposePayload(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  SpecialPurpose
		output []byte
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: SpecialPurpose{
				Len:  0x04,
				Data: []byte{0xab, 0xcd, 0xef},
			},
			output: []byte{0x04, 0xab, 0xcd, 0xef},
		},
		{
			Name: "testCase 2",
			input: SpecialPurpose{
				Len:  0,
				Data: nil,
			},
			output: []byte{0x00},
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

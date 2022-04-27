package item

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestExtendedReader(t *testing.T) {
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
			input: "01 03 07 09 0b 0d 0f 0e",
			item: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
			},
			err: nil,
			output: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				Primary:           []byte{0x01},
				Secondary:         []byte{0x03, 0x07, 0x09, 0x0b, 0x0d, 0x0f, 0x0e},
			},
		},
		{
			Name:  "testcase 2",
			input: "fe",
			item: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
			},
			err: nil,
			output: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				Primary:           []byte{0xfe},
				Secondary:         nil,
			},
		},
		{
			Name:  "testcase 3",
			input: "",
			item: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
			},
			err: io.EOF,
			output: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				Primary:           nil,
				Secondary:         nil,
			},
		},
		{
			Name:  "testcase 4",
			input: "ff",
			item: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
			},
			err: io.EOF,
			output: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				Primary:           []byte{0xff},
				Secondary:         nil,
			},
		},
		{
			Name:  "testcase 5",
			input: "ff",
			item: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 1,
			},
			err: io.ErrUnexpectedEOF,
			output: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 1,
				Primary:           nil,
				Secondary:         nil,
			},
		},
		{
			Name:  "testcase 6",
			input: "0001 000001 fffffe",
			item: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 3,
			},
			err: nil,
			output: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 3,
				Primary:           []byte{0x00, 0x01},
				Secondary:         []byte{0x00, 0x00, 0x01, 0xff, 0xff, 0xfe},
			},
		},
		{
			Name:  "testcase 7",
			input: "0001 000001 ffff",
			item: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 3,
			},
			err: io.ErrUnexpectedEOF,
			output: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 3,
				Primary:           []byte{0x00, 0x01},
				Secondary:         []byte{0x00, 0x00, 0x01},
			},
		},
	}

	for _, tc := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(tc.input)
		rb := bytes.NewReader(input)
		f := NewExtended(tc.item)

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

func TestExtendedString(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Extended
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Extended{
				Base: Base{
					FRN:          1,
					DataItemName: "I000/010",
					Description:  "Test item",
				},
				Primary:   []byte{0xc1},
				Secondary: []byte{0xab, 0xcd},
			},
			output: "I000/010:c1abcd",
		},
		{
			Name: "testCase 2",
			input: Extended{
				Base:      Base{},
				Primary:   nil,
				Secondary: nil,
			},
			output: ":",
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

func TestExtendedPayload(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Extended
		output []byte
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Extended{
				Primary:   []byte{0xc1},
				Secondary: []byte{0xab, 0xcd},
			},
			output: []byte{0xc1, 0xab, 0xcd},
		},
		{
			Name: "testCase 2",
			input: Extended{
				Primary:   nil,
				Secondary: nil,
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

package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestExtendedReader(t *testing.T) {
	// setup
	type testCase struct {
		Name      string
		input     string
		dataField uap.DataField
		output    Item
		err       error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:  "testCase 1",
			input: "01 03 07 09 0b 0d 0f 0e",
			dataField: uap.DataField{
				Type: uap.Extended,
				SizeItem: uap.SizeField{
					ForExtendedPrimary:   1,
					ForExtendedSecondary: 1,
				},
			},
			err: nil,
			output: &Extended{
				Base: Base{
					Type: uap.Extended,
				},
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				Primary:           []byte{0x01},
				Secondary:         []byte{0x03, 0x07, 0x09, 0x0b, 0x0d, 0x0f, 0x0e},
			},
		},
		{
			Name:  "testCase 2",
			input: "fe",
			dataField: uap.DataField{
				Type: uap.Extended,
				SizeItem: uap.SizeField{
					ForExtendedPrimary:   1,
					ForExtendedSecondary: 1,
				},
			},
			err: nil,
			output: &Extended{
				Base: Base{
					Type: uap.Extended,
				},
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				Primary:           []byte{0xfe},
			},
		},
		{
			Name:  "testCase 3",
			input: "",
			dataField: uap.DataField{
				Type: uap.Extended,
				SizeItem: uap.SizeField{
					ForExtendedPrimary:   1,
					ForExtendedSecondary: 1,
				},
			},
			err: io.EOF,
			output: &Extended{
				Base: Base{
					Type: uap.Extended,
				},
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				Primary:           nil,
			},
		},
		{
			Name:  "testCase 4",
			input: "ff",
			dataField: uap.DataField{
				Type: uap.Extended,
				SizeItem: uap.SizeField{
					ForExtendedPrimary:   1,
					ForExtendedSecondary: 1,
				},
			},
			err: io.EOF,
			output: &Extended{
				Base: Base{
					Type: uap.Extended,
				},
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				Primary:           []byte{0xff},
			},
		},
		{
			Name:  "testCase 5",
			input: "ff",
			dataField: uap.DataField{
				Type: uap.Extended,
				SizeItem: uap.SizeField{
					ForExtendedPrimary:   2,
					ForExtendedSecondary: 1,
				},
			},
			err: io.ErrUnexpectedEOF,
			output: &Extended{
				Base: Base{
					Type: uap.Extended,
				},
				PrimaryItemSize:   2,
				SecondaryItemSize: 1,
				Primary:           nil,
			},
		},
		{
			Name:  "testCase 6",
			input: "0001 000001 fffffe",
			dataField: uap.DataField{
				Type: uap.Extended,
				SizeItem: uap.SizeField{
					ForExtendedPrimary:   2,
					ForExtendedSecondary: 3,
				},
			},
			err: nil,
			output: &Extended{
				Base: Base{
					Type: uap.Extended,
				},
				PrimaryItemSize:   2,
				SecondaryItemSize: 3,
				Primary:           []byte{0x00, 0x01},
				Secondary:         []byte{0x00, 0x00, 0x01, 0xff, 0xff, 0xfe},
			},
		},
		{
			Name:  "testCase 7",
			input: "0001 000001 ffff",
			dataField: uap.DataField{
				Type: uap.Extended,
				SizeItem: uap.SizeField{
					ForExtendedPrimary:   2,
					ForExtendedSecondary: 3,
				},
			},
			err: io.ErrUnexpectedEOF,
			output: &Extended{
				Base: Base{
					Type: uap.Extended,
				},
				PrimaryItemSize:   2,
				SecondaryItemSize: 3,
				Primary:           []byte{0x00, 0x01},
				Secondary:         []byte{0x00, 0x00, 0x01},
			},
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(row.input)
		rb := bytes.NewReader(input)
		f := NewExtended(row.dataField)

		// Act
		err := f.Reader(rb)

		// Assert
		if err != row.err {
			t.Errorf(util.MsgFailInValue, row.Name, err, row.err)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, err, row.err)
		}

		if reflect.DeepEqual(f, row.output) == false {
			t.Errorf(util.MsgFailInValue, row.Name, f, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, f, row.output)
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
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Extended,
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

	for _, row := range dataSet {
		// Act
		s := row.input.String()

		// Assert
		if s != row.output {
			t.Errorf(util.MsgFailInValue, row.Name, s, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, s, row.output)
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
				Base: Base{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Extended,
				},
				Primary:   []byte{0xc1},
				Secondary: []byte{0xab, 0xcd},
			},
			output: []byte{0xc1, 0xab, 0xcd},
		},
		{
			Name: "testCase 2",
			input: Extended{
				Base:      Base{},
				Primary:   nil,
				Secondary: nil,
			},
			output: nil,
		},
	}

	for _, row := range dataSet {
		// Act
		res := row.input.Payload()

		// Assert
		if bytes.Equal(res, row.output) == false {
			t.Errorf(util.MsgFailInHex, row.Name, res, row.output)
		} else {
			t.Logf(util.MsgSuccessInHex, row.Name, res, row.output)
		}
	}
}

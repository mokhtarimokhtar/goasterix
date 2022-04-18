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
		Name   string
		input  string
		uap    uap.DataField
		output Extended
		err    error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:  "testcase 1",
			input: "01 03 07 09 0b 0d 0f 0e",
			uap: uap.DataField{
				Type: uap.Extended,
				Extended: uap.ExtendedField{
					PrimarySize:   1,
					SecondarySize: 1,
				},
			},
			err: nil,
			output: Extended{
				MetaItem: MetaItem{
					Type: uap.Extended,
				},
				Primary:   []byte{0x01},
				Secondary: []byte{0x03, 0x07, 0x09, 0x0b, 0x0d, 0x0f, 0x0e},
			},
		},
		{
			Name:  "testcase 2",
			input: "fe",
			uap: uap.DataField{
				Type: uap.Extended,
				Extended: uap.ExtendedField{
					PrimarySize:   1,
					SecondarySize: 1,
				},
			},
			err: nil,
			output: Extended{
				MetaItem: MetaItem{
					Type: uap.Extended,
				},
				Primary: []byte{0xfe},
			},
		},
		{
			Name:  "testcase 3",
			input: "",
			uap: uap.DataField{
				Type: uap.Extended,
				Extended: uap.ExtendedField{
					PrimarySize:   1,
					SecondarySize: 1,
				},
			},
			err: io.EOF,
			output: Extended{
				MetaItem: MetaItem{
					Type: uap.Extended,
				},
				Primary: nil,
			},
		},
		{
			Name:  "testcase 4",
			input: "ff",
			uap: uap.DataField{
				Type: uap.Extended,
				Extended: uap.ExtendedField{
					PrimarySize:   1,
					SecondarySize: 1,
				},
			},
			err: io.EOF,
			output: Extended{
				MetaItem: MetaItem{
					Type: uap.Extended,
				},
				Primary: []byte{0xff},
			},
		},
		{
			Name:  "testcase 5",
			input: "ff",
			uap: uap.DataField{
				Type: uap.Extended,
				Extended: uap.ExtendedField{
					PrimarySize:   2,
					SecondarySize: 1,
				},
			},
			err: io.ErrUnexpectedEOF,
			output: Extended{
				MetaItem: MetaItem{
					Type: uap.Extended,
				},
				Primary: nil,
			},
		},
		{
			Name:  "testcase 6",
			input: "0001 000001 fffffe",
			uap: uap.DataField{
				Type: uap.Extended,
				Extended: uap.ExtendedField{
					PrimarySize:   2,
					SecondarySize: 3,
				},
			},
			err: nil,
			output: Extended{
				MetaItem: MetaItem{
					Type: uap.Extended,
				},
				Primary:   []byte{0x00, 0x01},
				Secondary: []byte{0x00, 0x00, 0x01, 0xff, 0xff, 0xfe},
			},
		},
		{
			Name:  "testcase 7",
			input: "0001 000001 ffff",
			uap: uap.DataField{
				Type: uap.Extended,
				Extended: uap.ExtendedField{
					PrimarySize:   2,
					SecondarySize: 3,
				},
			},
			err: io.ErrUnexpectedEOF,
			output: Extended{
				MetaItem: MetaItem{
					Type: uap.Extended,
				},
				Primary:   []byte{0x00, 0x01},
				Secondary: []byte{0x00, 0x00, 0x01},
			},
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(row.input)
		rb := bytes.NewReader(input)
		f := Extended{}

		// Act
		err := f.Reader(rb, row.uap)

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
			Name: "testcase 1",
			input: Extended{
				MetaItem: MetaItem{
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
			Name: "testcase 2",
			input: Extended{
				MetaItem:  MetaItem{},
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
			Name: "testcase 1",
			input: Extended{
				MetaItem: MetaItem{
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
			Name: "testcase 2",
			input: Extended{
				MetaItem:  MetaItem{},
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

func TestExtendedFrn(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Extended
		output uint8
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: Extended{
				MetaItem: MetaItem{
					FRN:         7,
					DataItem:    "I000/070",
					Description: "Test item",
					Type:        uap.Extended,
				},
				Primary:   nil,
				Secondary: nil,
			},
			output: 7,
		},
		{
			Name: "testcase 2",
			input: Extended{
				MetaItem:  MetaItem{},
				Primary:   nil,
				Secondary: nil,
			},
			output: 0,
		},
	}

	for _, row := range dataSet {
		// Act
		res := row.input.Frn()

		// Assert
		if res != row.output {
			t.Errorf(util.MsgFailInValue, row.Name, res, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, res, row.output)
		}
	}
}

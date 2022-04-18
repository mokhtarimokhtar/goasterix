package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

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
			Name: "testcase 1",
			input: Fixed{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Fixed,
				},
				Data: []byte{0xab, 0xcd},
			},
			output: "I000/010:abcd",
		},
		{
			Name: "testcase 2",
			input: Fixed{
				MetaItem: MetaItem{},
				Data:     []byte{},
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
			Name: "testcase 1",
			input: Fixed{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Fixed,
				},
				Data: []byte{0xab, 0xcd},
			},
			output: []byte{0xab, 0xcd},
		},
		{
			Name: "testcase 2",
			input: Fixed{
				MetaItem: MetaItem{},
				Data:     nil,
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

func TestFixedReader(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  string
		uap    uap.DataField
		output Fixed
		err    error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:  "testcase 1",
			input: "01 02 03 04 05 06 07 08",
			uap: uap.DataField{
				FRN:         8,
				DataItem:    "I000/080",
				Description: "Test item",
				Type:        uap.Fixed,
				Fixed:       uap.FixedField{Size: 8},
			},
			err: nil,
			output: Fixed{
				MetaItem: MetaItem{
					FRN:         8,
					DataItem:    "I000/080",
					Description: "Test item",
					Type:        uap.Fixed,
				},
				Data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			},
		},
		{
			Name:  "testcase 2",
			input: "01 02 03 04 05 06 07",
			uap: uap.DataField{
				FRN:         8,
				DataItem:    "I000/080",
				Description: "Test item",
				Type:        uap.Fixed,
				Fixed:       uap.FixedField{Size: 8},
			},
			output: Fixed{
				MetaItem: MetaItem{
					FRN:         8,
					DataItem:    "I000/080",
					Description: "Test item",
					Type:        uap.Fixed,
				},
			},
			err: io.ErrUnexpectedEOF,
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(row.input)
		rb := bytes.NewReader(input)
		f := Fixed{}

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

func TestFixedFrn(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Fixed
		output uint8
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: Fixed{
				MetaItem: MetaItem{
					FRN:         7,
					DataItem:    "I000/070",
					Description: "Test item",
					Type:        uap.Fixed,
				},
				Data: []byte{0xab, 0xcd},
			},
			output: 7,
		},
		{
			Name: "testcase 2",
			input: Fixed{
				MetaItem: MetaItem{},
				Data:     []byte{},
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

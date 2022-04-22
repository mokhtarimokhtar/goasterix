package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestFixedReader(t *testing.T) {
	// setup
	type testCase struct {
		Name      string
		input     string
		FRN         uint8
		DataItem    string
		Description string
		Type        uap.TypeField
		Size        uap.SizeField
		output Item
		err    error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:  "testCase 1",
			input: "01 02 03 04 05 06 07 08",
			FRN:         8,
			DataItem:    "I000/080",
			Description: "Test item",
			Type:        uap.Fixed,
			Size: uap.SizeField{
				ForFixed:             8,
			},
			err: nil,
			output: &Fixed{
				Base: Base{
					FRN:         8,
					DataItem:    "I000/080",
					Description: "Test item",
					Type:        uap.Fixed,
				},
				Size: 8,
				Data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			},
		},
		{
			Name:  "testCase 2",
			input: "01 02 03 04 05 06 07",
			FRN:         8,
			DataItem:    "I000/080",
			Description: "Test item",
			Type:        uap.Fixed,
			Size: uap.SizeField{
				ForFixed:             8,
			},
			output: &Fixed{
				Base: Base{
					FRN:         8,
					DataItem:    "I000/080",
					Description: "Test item",
					Type:        uap.Fixed,
				},
				Size: 8,
			},
			err: io.ErrUnexpectedEOF,
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(row.input)
		rb := bytes.NewReader(input)
		d := uap.DataFieldFactory(row.FRN, row.DataItem, row.Description, row.Type, row.Size)

		//f := newFixed(row.dataField)
		f := newFixed(d)

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
			Name: "testCase 2",
			input: Fixed{
				Base: Base{},
				Data: []byte{},
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
			Name: "testCase 1",
			input: Fixed{
				Base: Base{
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
			Name: "testCase 2",
			input: Fixed{
				Base: Base{},
				Data: nil,
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

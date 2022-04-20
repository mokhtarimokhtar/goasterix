package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestSpecialPurposeReader(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input     string
		dataField uap.DataField
		output    Item
		err    error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:  "testCase 1",
			input: "08 01 02 03 04 05 06 07",
			dataField: uap.DataField{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.SP,
			},
			err: nil,
			output: &SpecialPurpose{
				Base: Base{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.SP,
				},
				Len: 0x08,
				Data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
			},
		},
		{
			Name:  "testCase 2",
			input: "08 01 02 03 04 05 06",
			dataField: uap.DataField{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.SP,
			},
			err: io.ErrUnexpectedEOF,
			output: &SpecialPurpose{
				Base: Base{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.SP,
				},
				Len: 0x08,
				Data: nil,
			},
		},
		{
			Name:  "testCase 3",
			input: "",
			dataField: uap.DataField{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.SP,
			},
			err: io.EOF,
			output: &SpecialPurpose{
				Base: Base{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.SP,
				},
				Len: 0x00,
				Data: nil,
			},
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(row.input)
		rb := bytes.NewReader(input)
		f := NewSpecialPurpose(row.dataField)

		// Act
		//err := f.Reader(rb, row.dataField)
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
					Type:        uap.SP,
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
				Base: Base{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.SP,
				},
				Len:  0x04,
				Data: []byte{0xab, 0xcd, 0xef},
			},
			output: []byte{0x04, 0xab, 0xcd, 0xef},
		},
		{
			Name:   "testCase 2",
			input:  SpecialPurpose{
				Base: Base{},
				Len:  0,
				Data: nil,
			},
			output: []byte{0x00},
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
package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestRepetitiveReader(t *testing.T) {
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
			input: "03 01 02 03 01 02 03 01 02 03",
			dataField: uap.DataField{
				Type:       uap.Repetitive,
				Repetitive: uap.RepetitiveField{SubItemSize: 3},
			},
			err: nil,
			output: &Repetitive{
				Base: Base{
					Type: uap.Repetitive,
				},
				SubItemSize: 3,
				Rep:  0x03,
				Data: []byte{0x01, 0x02, 0x03, 0x01, 0x02, 0x03, 0x01, 0x02, 0x03},
			},
		},
		{
			Name:  "testCase 2",
			input: "04 01 02 03 01 02 03 01 02 03",
			dataField: uap.DataField{
				Type:       uap.Repetitive,
				Repetitive: uap.RepetitiveField{SubItemSize: 3},
			},
			err: io.ErrUnexpectedEOF,
			output: &Repetitive{
				Base: Base{
					Type: uap.Repetitive,
				},
				SubItemSize: 3,
				Rep:  0x04,
				Data: nil,
			},
		},
		{
			Name:  "testCase 3",
			input: "",
			dataField: uap.DataField{
				Type:       uap.Repetitive,
				Repetitive: uap.RepetitiveField{SubItemSize: 3},
			},
			err: io.EOF,
			output: &Repetitive{
				Base: Base{
					Type: uap.Repetitive,
				},
				SubItemSize: 3,
				Rep:  0x00,
				Data: nil,
			},
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(row.input)
		rb := bytes.NewReader(input)
		f := NewRepetitive(row.dataField)

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

func TestRepetitiveString(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Repetitive
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Repetitive{
				Base: Base{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Repetitive,
				},
				Rep:  0x02,
				Data: []byte{0xab, 0xcd},
			},
			output: "I000/010:02abcd",
		},
		{
			Name: "testCase 2",
			input: Repetitive{
				Base: Base{},
				Rep:  0,
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

func TestRepetitivePayload(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Repetitive
		output []byte
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Repetitive{
				Base: Base{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Repetitive,
				},
				Rep:  0x02,
				Data: []byte{0xab, 0xcd},
			},
			output: []byte{0x02, 0xab, 0xcd},
		},
		{
			Name: "testCase 2",
			input: Repetitive{
				Base: Base{},
				Rep:  0,
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


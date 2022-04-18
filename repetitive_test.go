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
		input  string
		uap    uap.DataField
		output Repetitive
		err    error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:  "testcase 1",
			input: "03 01 02 03 01 02 03 01 02 03",
			uap: uap.DataField{
				Type:       uap.Repetitive,
				Repetitive: uap.RepetitiveField{SubItemSize: 3},
			},
			err: nil,
			output: Repetitive{
				MetaItem: MetaItem{
					Type: uap.Repetitive,
				},
				Rep:  0x03,
				Data: []byte{0x01, 0x02, 0x03, 0x01, 0x02, 0x03, 0x01, 0x02, 0x03},
			},
		},
		{
			Name:  "testcase 2",
			input: "04 01 02 03 01 02 03 01 02 03",
			uap: uap.DataField{
				Type:       uap.Repetitive,
				Repetitive: uap.RepetitiveField{SubItemSize: 3},
			},
			err: io.ErrUnexpectedEOF,
			output: Repetitive{
				MetaItem: MetaItem{
					Type: uap.Repetitive,
				},
				Rep:  0x04,
				Data: nil,
			},
		},
		{
			Name:  "testcase 3",
			input: "",
			uap: uap.DataField{
				Type:       uap.Repetitive,
				Repetitive: uap.RepetitiveField{SubItemSize: 3},
			},
			err: io.EOF,
			output: Repetitive{
				MetaItem: MetaItem{
					Type: uap.Repetitive,
				},
				Rep:  0x00,
				Data: nil,
			},
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(row.input)
		rb := bytes.NewReader(input)
		f := Repetitive{}

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
			Name: "testcase 1",
			input: Repetitive{
				MetaItem: MetaItem{
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
			Name: "testcase 2",
			input: Repetitive{
				MetaItem: MetaItem{},
				Rep:      0,
				Data:     nil,
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
			Name: "testcase 1",
			input: Repetitive{
				MetaItem: MetaItem{
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
			Name: "testcase 2",
			input: Repetitive{
				MetaItem: MetaItem{},
				Rep:      0,
				Data:     nil,
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

func TestRepetitiveFrn(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Repetitive
		output uint8
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: Repetitive{
				MetaItem: MetaItem{
					FRN:         7,
					DataItem:    "I000/070",
					Description: "Test item",
					Type:        uap.Repetitive,
				},
				Rep:  0,
				Data: nil,
			},
			output: 7,
		},
		{
			Name: "testcase 2",
			input: Repetitive{
				MetaItem: MetaItem{},
				Rep:      0,
				Data:     nil,
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

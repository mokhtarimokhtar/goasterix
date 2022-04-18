package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestExplicitReader(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  string
		uap    uap.DataField
		output Explicit
		err    error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:  "testCase 1",
			input: "08 01 02 03 04 05 06 07",
			uap: uap.DataField{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.Explicit,
				Explicit:       uap.ExplicitField{},
			},
			err: nil,
			output: Explicit{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Explicit,
				},
				Len: 0x08,
				Data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
			},
		},
		{
			Name:  "testCase 2",
			input: "08 01 02 03 04 05 06",
			uap: uap.DataField{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.Explicit,
				Explicit:       uap.ExplicitField{},
			},
			err: io.ErrUnexpectedEOF,
			output: Explicit{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Explicit,
				},
				Len: 0x08,
				Data: nil,
			},
		},
		{
			Name:  "testCase 3",
			input: "",
			uap: uap.DataField{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.Explicit,
				Explicit:       uap.ExplicitField{},
			},
			err: io.EOF,
			output: Explicit{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Explicit,
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
		f := Explicit{}

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

func TestExplicitString(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Explicit
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Explicit{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Explicit,
				},
				Len:  0x04,
				Data: []byte{0xab, 0xcd, 0xef},
			},
			output: "I000/010:04abcdef",
		},
		{
			Name: "testCase 2",
			input: Explicit{
				MetaItem: MetaItem{},
				Len:      0,
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

func TestExplicitPayload(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Explicit
		output []byte
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Explicit{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Explicit,
				},
				Len:  0x04,
				Data: []byte{0xab, 0xcd, 0xef},
			},
			output: []byte{0x04, 0xab, 0xcd, 0xef},
		},
		{
			Name: "testCase 2",
			input: Explicit{
				MetaItem: MetaItem{},
				Len:      0,
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

func TestExplicitFrn(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Explicit
		output uint8
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Explicit{
				MetaItem: MetaItem{
					FRN:         7,
					DataItem:    "I000/070",
					Description: "Test item",
					Type:        uap.SP,
				},
				Len:  0,
				Data: nil,
			},
			output: 7,
		},
		{
			Name: "testCase 2",
			input: Explicit{
				MetaItem: MetaItem{},
				Len:      0,
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

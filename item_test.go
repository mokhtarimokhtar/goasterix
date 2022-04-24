package goasterix

import (
	"bytes"
	"encoding/binary"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"math"
	"reflect"
	"testing"
)

func TestNewBase(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  uap.DataField
		output Base
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: uap.DataField{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.Fixed,
				SizeItem:    uap.SizeField{ForFixed: 1},
			},
			output: Base{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.Fixed,
			},
		},
		{
			Name: "testCase 2",
			input: uap.DataField{
				FRN:         0,
				DataItem:    "",
				Description: "",
				Type:        0,
				SizeItem:    uap.SizeField{},
			},
			output: Base{},
		},
		{
			Name: "testCase 3",
			input: uap.DataField{
				FRN:         3,
				DataItem:    "I000/030",
				Description: "Test item",
				Type:        uap.Extended,
				SizeItem: uap.SizeField{
					ForExtendedPrimary:   1,
					ForExtendedSecondary: 2,
				},
			},
			output: Base{
				FRN:         3,
				DataItem:    "I000/030",
				Description: "Test item",
				Type:        uap.Extended,
			},
		},
		{
			Name: "testCase 4",
			input: uap.DataField{
				FRN:         4,
				DataItem:    "I000/040",
				Description: "Test item",
				Type:        uap.Explicit,
			},
			output: Base{
				FRN:         4,
				DataItem:    "I000/040",
				Description: "Test item",
				Type:        uap.Explicit,
			},
		},
		{
			Name: "testCase 5",
			input: uap.DataField{
				FRN:         5,
				DataItem:    "I000/050",
				Description: "Test item",
				Type:        uap.Repetitive,
				SizeItem:    uap.SizeField{ForRepetitive: 2},
			},
			output: Base{
				FRN:         5,
				DataItem:    "I000/050",
				Description: "Test item",
				Type:        uap.Repetitive,
			},
		},
		{
			Name: "testCase 6",
			input: uap.DataField{
				FRN:         6,
				DataItem:    "I000/060",
				Description: "Test item",
				Type:        uap.Compound,
				Compound:    []uap.DataField{},
			},
			output: Base{
				FRN:         6,
				DataItem:    "I000/060",
				Description: "Test item",
				Type:        uap.Compound,
			},
		},
	}

	for _, row := range dataSet {
		// Arrange
		m := Base{}
		// Act
		m.NewBase(row.input)

		// Assert
		if reflect.DeepEqual(m, row.output) == false {
			t.Errorf(util.MsgFailInValue, row.Name, m, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, m, row.output)
		}
	}

}

func TestBaseFrn(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Base
		output uint8
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Base{
				FRN:         7,
				DataItem:    "I000/070",
				Description: "Test item",
				Type:        uap.Fixed,
			},
			output: 7,
		},
		{
			Name:   "testCase 2",
			input:  Base{},
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

func TestFromToBitReader8(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  byte
		from   uint8
		to     uint8
		output byte
	}
	// Arrange
	dataSet := []testCase{
		{input: 0xff, from: 8, to: 8, output: 0x01},
		{input: 0xff, from: 8, to: 7, output: 0x03},
		{input: 0xff, from: 8, to: 6, output: 0x07},
		{input: 0xff, from: 8, to: 5, output: 0x0f},
		{input: 0xff, from: 8, to: 4, output: 0x1f},
		{input: 0xff, from: 8, to: 3, output: 0x3f},
		{input: 0xff, from: 8, to: 2, output: 0x7f},
		{input: 0xff, from: 8, to: 1, output: 0xff},

		{input: 0xff, from: 5, to: 4, output: 0x03},
		{input: 0xff, from: 6, to: 3, output: 0x0f},
		{input: 0xff, from: 7, to: 2, output: 0x3f},

		{input: 0xff, from: 1, to: 1, output: 0x01},
		{input: 0xff, from: 2, to: 1, output: 0x03},
		{input: 0xff, from: 3, to: 1, output: 0x07},
		{input: 0xff, from: 4, to: 1, output: 0x0f},
		{input: 0xff, from: 5, to: 1, output: 0x1f},
		{input: 0xff, from: 6, to: 1, output: 0x3f},
		{input: 0xff, from: 7, to: 1, output: 0x7f},
		{input: 0xff, from: 8, to: 1, output: 0xff},
	}
	for _, tc := range dataSet {
		// Act
		res := FromToBitReader8(tc.input, tc.from, tc.to)

		// Assert
		if res != tc.output {
			t.Errorf(util.MsgFailInHex, tc.Name, res, tc.output)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, res, tc.output)
		}
	}
}

func TestFromToBitReader16(t *testing.T) {
	// setup
	type testCase struct {
		Name string
		//input  uint16
		input []byte
		from  uint8
		to    uint8
		//output uint16
		output []byte
	}
	var dataSet []testCase
	// Arrange
	for i := 1; i <= 16; i++ {
		for j := i; j <= 16; j++ {
			tmo := uint16(math.Pow(2, float64(j-i+1))) - 1
			out := make([]byte, 2)
			binary.BigEndian.PutUint16(out, tmo)

			tmp := testCase{
				//input:  0xffff,
				input: []byte{0xff, 0xff},
				from:  uint8(j),
				to:    uint8(i),
				//output: uint16(math.Pow(2, float64(j-i+1))) - 1,
				output: out,
			}
			dataSet = append(dataSet, tmp)
		}
	}

	for _, tc := range dataSet {
		// Act
		res := FromToBitReader16(tc.input, tc.from, tc.to)

		// Assert
		//if res != tc.output {
		if bytes.Equal(res, tc.output) == false {
			t.Errorf(util.MsgFailInHex, tc.Name, res, tc.output)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, res, tc.output)
		}
	}
}

func TestFromToBitReader32(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  uint32
		from   uint8
		to     uint8
		output uint32
	}
	var dataSet []testCase

	for i := 1; i <= 32; i++ {
		for j := i; j <= 32; j++ {
			tmp := testCase{
				input:  0xffffffff,
				from:   uint8(j),
				to:     uint8(i),
				output: uint32(math.Pow(2, float64(j-i+1))) - 1,
			}
			dataSet = append(dataSet, tmp)
		}
	}
	// Arrange
	for _, tc := range dataSet {
		// Act
		res := FromToBitReader32(tc.input, tc.from, tc.to)

		// Assert
		if res != tc.output {
			t.Errorf(util.MsgFailInHex, tc.Name, res, tc.output)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, res, tc.output)
		}
	}
}

func TestGetBitsFromTo(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  []byte
		from   uint8
		to     uint8
		output []byte
	}
	dataSet := []testCase{
		{
			input:  []byte{0xff, 0xff},
			from:   15,
			to:     1,
			output: []byte{0x07, 0xff},
		},
	}
	for _, tc := range dataSet {
		// Act
		res := GetBitsFromTo(tc.input, tc.from, tc.to)

		// Assert
		if bytes.Equal(res, tc.output) != false {
			t.Errorf(util.MsgFailInHex, tc.Name, res, tc.output)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, res, tc.output)
		}
	}
}

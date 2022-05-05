package item

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestExtendedClone(t *testing.T) {
	// Arrange
	input := Extended{
		Base: Base{
			FRN:          1,
			DataItemName: "I000/010",
			Description:  "Test item",
			Type:         ExtendedField,
		},
		PrimaryItemSize:   1,
		SecondaryItemSize: 1,
	}
	output := &Extended{
		Base: Base{
			FRN:          1,
			DataItemName: "I000/010",
			Description:  "Test item",
			Type:         ExtendedField,
		},
		PrimaryItemSize:   1,
		SecondaryItemSize: 1,
	}
	// Act
	res := input.Clone()

	// Assert
	if reflect.DeepEqual(res, output) == false {
		t.Errorf(util.MsgFailInValue, "", res, output)
	} else {
		t.Logf(util.MsgSuccessInValue, "", res, output)
	}

}

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
			input: "e1 06",
			item: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6},
					{Name: "SIM", Type: BitField, Bit: 5},
					{Name: "RDP", Type: BitField, Bit: 4},
					{Name: "SPI", Type: BitField, Bit: 3},
					{Name: "RAB", Type: BitField, Bit: 2},
					{Name: "FX", Type: BitField, Bit: 1},

					{Name: "TST", Type: BitField, Bit: 8},
					{Name: "ERR", Type: BitField, Bit: 7},
					{Name: "XPP", Type: BitField, Bit: 6},
					{Name: "ME", Type: BitField, Bit: 5},
					{Name: "MI", Type: BitField, Bit: 4},
					{Name: "FOE/FRI", Type: FromToField, From: 3, To: 2},
					{Name: "FX", Type: BitField, Bit: 1},
				},
			},
			err: nil,
			output: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6, Data: []byte{0x07}},
					{Name: "SIM", Type: BitField, Bit: 5, Data: []byte{0x00}},
					{Name: "RDP", Type: BitField, Bit: 4, Data: []byte{0x00}},
					{Name: "SPI", Type: BitField, Bit: 3, Data: []byte{0x00}},
					{Name: "RAB", Type: BitField, Bit: 2, Data: []byte{0x00}},

					{Name: "TST", Type: BitField, Bit: 8, Data: []byte{0x00}},
					{Name: "ERR", Type: BitField, Bit: 7, Data: []byte{0x00}},
					{Name: "XPP", Type: BitField, Bit: 6, Data: []byte{0x00}},
					{Name: "ME", Type: BitField, Bit: 5, Data: []byte{0x00}},
					{Name: "MI", Type: BitField, Bit: 4, Data: []byte{0x00}},
					{Name: "FOE/FRI", Type: FromToField, From: 3, To: 2, Data: []byte{0x03}},
				},
			},
		},
		{
			Name:  "testCase 2",
			input: "e1 8606",
			item: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 2,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6},
					{Name: "SIM", Type: BitField, Bit: 5},
					{Name: "RDP", Type: BitField, Bit: 4},
					{Name: "SPI", Type: BitField, Bit: 3},
					{Name: "RAB", Type: BitField, Bit: 2},
					{Name: "FX", Type: BitField, Bit: 1},

					{Name: "TST1", Type: BitField, Bit: 16},
					{Name: "ERR1", Type: BitField, Bit: 15},
					{Name: "XPP1", Type: BitField, Bit: 14},
					{Name: "ME1", Type: BitField, Bit: 13},
					{Name: "MI1", Type: BitField, Bit: 12},
					{Name: "FOE/FRI1", Type: FromToField, From: 11, To: 10},
					{Name: "TST2", Type: BitField, Bit: 8},
					{Name: "ERR2", Type: BitField, Bit: 7},
					{Name: "XPP2", Type: BitField, Bit: 6},
					{Name: "ME2", Type: BitField, Bit: 5},
					{Name: "MI2", Type: BitField, Bit: 4},
					{Name: "FOE/FRI2", Type: FromToField, From: 3, To: 2},
					{Name: "FX", Type: BitField, Bit: 1},
				},
			},
			err: nil,
			output: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 2,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6, Data: []byte{0x07}},
					{Name: "SIM", Type: BitField, Bit: 5, Data: []byte{0x00}},
					{Name: "RDP", Type: BitField, Bit: 4, Data: []byte{0x00}},
					{Name: "SPI", Type: BitField, Bit: 3, Data: []byte{0x00}},
					{Name: "RAB", Type: BitField, Bit: 2, Data: []byte{0x00}},

					{Name: "TST1", Type: BitField, Bit: 16, Data: []byte{0x01}},
					{Name: "ERR1", Type: BitField, Bit: 15, Data: []byte{0x00}},
					{Name: "XPP1", Type: BitField, Bit: 14, Data: []byte{0x00}},
					{Name: "ME1", Type: BitField, Bit: 13, Data: []byte{0x00}},
					{Name: "MI1", Type: BitField, Bit: 12, Data: []byte{0x00}},
					{Name: "FOE/FRI1", Type: FromToField, From: 11, To: 10, Data: []byte{0x03}},
					{Name: "TST2", Type: BitField, Bit: 8, Data: []byte{0x00}},
					{Name: "ERR2", Type: BitField, Bit: 7, Data: []byte{0x00}},
					{Name: "XPP2", Type: BitField, Bit: 6, Data: []byte{0x00}},
					{Name: "ME2", Type: BitField, Bit: 5, Data: []byte{0x00}},
					{Name: "MI2", Type: BitField, Bit: 4, Data: []byte{0x00}},
					{Name: "FOE/FRI2", Type: FromToField, From: 3, To: 2, Data: []byte{0x03}},
				},
			},
		},
		{
			Name:  "testCase 3",
			input: "e1 07 06",
			item: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6},
					{Name: "SIM", Type: BitField, Bit: 5},
					{Name: "RDP", Type: BitField, Bit: 4},
					{Name: "SPI", Type: BitField, Bit: 3},
					{Name: "RAB", Type: BitField, Bit: 2},
					{Name: "FX", Type: BitField, Bit: 1},

					{Name: "TST", Type: BitField, Bit: 8},
					{Name: "ERR", Type: BitField, Bit: 7},
					{Name: "XPP", Type: BitField, Bit: 6},
					{Name: "ME", Type: BitField, Bit: 5},
					{Name: "MI", Type: BitField, Bit: 4},
					{Name: "FOE/FRI", Type: FromToField, From: 3, To: 2},
					{Name: "FX", Type: BitField, Bit: 1},

					{Name: "TST", Type: BitField, Bit: 8},
					{Name: "ERR", Type: BitField, Bit: 7},
					{Name: "XPP", Type: BitField, Bit: 6},
					{Name: "ME", Type: BitField, Bit: 5},
					{Name: "MI", Type: BitField, Bit: 4},
					{Name: "FOE/FRI", Type: FromToField, From: 3, To: 2},
					{Name: "FX", Type: BitField, Bit: 1},
				},
			},
			err: nil,
			output: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6, Data: []byte{0x07}},
					{Name: "SIM", Type: BitField, Bit: 5, Data: []byte{0x00}},
					{Name: "RDP", Type: BitField, Bit: 4, Data: []byte{0x00}},
					{Name: "SPI", Type: BitField, Bit: 3, Data: []byte{0x00}},
					{Name: "RAB", Type: BitField, Bit: 2, Data: []byte{0x00}},

					{Name: "TST", Type: BitField, Bit: 8, Data: []byte{0x00}},
					{Name: "ERR", Type: BitField, Bit: 7, Data: []byte{0x00}},
					{Name: "XPP", Type: BitField, Bit: 6, Data: []byte{0x00}},
					{Name: "ME", Type: BitField, Bit: 5, Data: []byte{0x00}},
					{Name: "MI", Type: BitField, Bit: 4, Data: []byte{0x00}},
					{Name: "FOE/FRI", Type: FromToField, From: 3, To: 2, Data: []byte{0x03}},

					{Name: "TST", Type: BitField, Bit: 8, Data: []byte{0x00}},
					{Name: "ERR", Type: BitField, Bit: 7, Data: []byte{0x00}},
					{Name: "XPP", Type: BitField, Bit: 6, Data: []byte{0x00}},
					{Name: "ME", Type: BitField, Bit: 5, Data: []byte{0x00}},
					{Name: "MI", Type: BitField, Bit: 4, Data: []byte{0x00}},
					{Name: "FOE/FRI", Type: FromToField, From: 3, To: 2, Data: []byte{0x03}},
				},
			},
		},
		{
			Name:  "testCase 4",
			input: "e0",
			item: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6},
					{Name: "SIM", Type: BitField, Bit: 5},
					{Name: "RDP", Type: BitField, Bit: 4},
					{Name: "SPI", Type: BitField, Bit: 3},
					{Name: "RAB", Type: BitField, Bit: 2},
					{Name: "FX", Type: BitField, Bit: 1},

					{Name: "TST", Type: BitField, Bit: 8},
					{Name: "ERR", Type: BitField, Bit: 7},
					{Name: "XPP", Type: BitField, Bit: 6},
					{Name: "ME", Type: BitField, Bit: 5},
					{Name: "MI", Type: BitField, Bit: 4},
					{Name: "FOE/FRI", Type: FromToField, From: 3, To: 2},
					{Name: "FX", Type: BitField, Bit: 1},
				},
			},
			err: nil,
			output: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6, Data: []byte{0x07}},
					{Name: "SIM", Type: BitField, Bit: 5, Data: []byte{0x00}},
					{Name: "RDP", Type: BitField, Bit: 4, Data: []byte{0x00}},
					{Name: "SPI", Type: BitField, Bit: 3, Data: []byte{0x00}},
					{Name: "RAB", Type: BitField, Bit: 2, Data: []byte{0x00}},
				},
			},
		},
		{
			Name:  "testCase 5",
			input: "",
			item: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6},
					{Name: "SIM", Type: BitField, Bit: 5},
					{Name: "RDP", Type: BitField, Bit: 4},
					{Name: "SPI", Type: BitField, Bit: 3},
					{Name: "RAB", Type: BitField, Bit: 2},
					{Name: "FX", Type: BitField, Bit: 1},
				},
			},
			err: io.EOF,
			output: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				SubItems:          nil,
			},
		},
		{
			Name:  "testCase 6",
			input: "e1",
			item: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6},
					{Name: "SIM", Type: BitField, Bit: 5},
					{Name: "RDP", Type: BitField, Bit: 4},
					{Name: "SPI", Type: BitField, Bit: 3},
					{Name: "RAB", Type: BitField, Bit: 2},
					{Name: "FX", Type: BitField, Bit: 1},
				},
			},
			err: io.EOF,
			output: &Extended{
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6, Data: []byte{0x07}},
					{Name: "SIM", Type: BitField, Bit: 5, Data: []byte{0x00}},
					{Name: "RDP", Type: BitField, Bit: 4, Data: []byte{0x00}},
					{Name: "SPI", Type: BitField, Bit: 3, Data: []byte{0x00}},
					{Name: "RAB", Type: BitField, Bit: 2, Data: []byte{0x00}},
				},
			},
		},
		{
			Name:  "testCase 7",
			input: "e1",
			item: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 1,
				SubItems: []SubItem{
					{Name: "TYP", Type: FromToField, From: 8, To: 6},
					{Name: "SIM", Type: BitField, Bit: 5},
					{Name: "RDP", Type: BitField, Bit: 4},
					{Name: "SPI", Type: BitField, Bit: 3},
					{Name: "RAB", Type: BitField, Bit: 2},
					{Name: "FX", Type: BitField, Bit: 1},
				},
			},
			err: io.ErrUnexpectedEOF,
			output: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 1,
				SubItems:          nil,
			},
		},
		{
			Name:  "testCase 8",
			input: "ffff fffffe",
			item: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 3,
				SubItems: []SubItem{
					{Name: "sub-item-1", Type: FromToField, From: 16, To: 2},
					{Name: "FX", Type: BitField, Bit: 1},
					{Name: "sub-item-2", Type: FromToField, From: 24, To: 2},
					{Name: "FX", Type: BitField, Bit: 1},
				},
			},
			err: nil,
			output: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 3,
				SubItems: []SubItem{
					{Name: "sub-item-1", Type: FromToField, From: 16, To: 2, Data: []byte{0x7f, 0xff}},
					{Name: "sub-item-2", Type: FromToField, From: 24, To: 2, Data: []byte{0x7f, 0xff, 0xff}},
				},
			},
		},
		{
			Name:  "testCase 9",
			input: "ffff ffffff",
			item: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 3,
				SubItems: []SubItem{
					{Name: "sub-item-1", Type: FromToField, From: 16, To: 2},
					{Name: "FX", Type: BitField, Bit: 1},
					{Name: "sub-item-2", Type: FromToField, From: 24, To: 2},
					{Name: "FX", Type: BitField, Bit: 1},
				},
			},
			err: io.EOF,
			output: &Extended{
				PrimaryItemSize:   2,
				SecondaryItemSize: 3,
				SubItems: []SubItem{
					{Name: "sub-item-1", Type: FromToField, From: 16, To: 2, Data: []byte{0x7f, 0xff}},
					{Name: "sub-item-2", Type: FromToField, From: 24, To: 2, Data: []byte{0x7f, 0xff, 0xff}},
				},
			},
		},
	}

	for _, tc := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(tc.input)
		rb := bytes.NewReader(input)
		f := tc.item.Clone()

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
				SubItems: []SubItem{
					{Name: "SUB-A", Type: FromToField, From: 8, To: 3, Data: []byte{0x3f}},
					{Name: "SUB-B", Type: BitField, Bit: 1, Data: []byte{0x01}},
				},
			},
			output: "I000/010:[SUB-A:3f][SUB-B:01]",
		},
		{
			Name: "testCase 2",
			input: Extended{
				Base: Base{},
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

/*
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
*/

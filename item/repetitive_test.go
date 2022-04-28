package item

import (
	"bytes"
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
		item   DataItem
		output DataItem
		err    error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:  "testCase 1",
			input: "03 aaaaaa bbbbbb cccccc",
			item: &Repetitive{
				SubItemSize: 3,
			},
			err: nil,
			output: &Repetitive{
				SubItemSize: 3,
				Rep:         0x03,
				Data:        []byte{0xaa, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xcc, 0xcc, 0xcc},
			},
		},
		{
			Name:  "testCase 2",
			input: "03 8aaaaa 0bbbbb 8ccccc", // 8a 1000-1010 1010-1010 1010-1010
			item: &Repetitive{
				SubItemSize: 3,
				SubItems: []SubItem{
					&SubItemBit{
						Type: BitField,
						Pos:  BitPosition{Bit: 24},
					},
					&SubItemFromTo{
						Type: FromToField,
						Pos:  BitPosition{From: 20, To: 1},
					},
				},
			},
			err: nil,
			output: &Repetitive{
				SubItemSize: 3,
				Rep:         0x03,
				SubItems: []SubItem{
					&SubItemBit{
						Type: BitField,
						Pos:  BitPosition{Bit: 24},
						Data: []byte{0x01},
					},
					&SubItemFromTo{
						Type: FromToField,
						Pos:  BitPosition{From: 20, To: 1},
						Data: []byte{0x0a, 0xaa, 0xaa},
					},
					&SubItemBit{
						Type: BitField,
						Pos:  BitPosition{Bit: 24},
						Data: []byte{0x00},
					},
					&SubItemFromTo{
						Type: FromToField,
						Pos:  BitPosition{From: 20, To: 1},
						Data: []byte{0x0b, 0xbb, 0xbb},
					},
					&SubItemBit{
						Type: BitField,
						Pos:  BitPosition{Bit: 24},
						Data: []byte{0x01},
					},
					&SubItemFromTo{
						Type: FromToField,
						Pos:  BitPosition{From: 20, To: 1},
						Data: []byte{0x0c, 0xcc, 0xcc},
					},
				},
			},
		},
		{
			Name:  "testCase 3",
			input: "04 aaaaaa bbbbbb cccccc",
			item: &Repetitive{
				SubItemSize: 3,
			},
			err: io.ErrUnexpectedEOF,
			output: &Repetitive{
				SubItemSize: 3,
				Rep:         0x04,
				Data:        nil,
			},
		},
		{
			Name:  "testCase 4",
			input: "",
			item: &Repetitive{
				SubItemSize: 3,
			},
			err: io.EOF,
			output: &Repetitive{
				SubItemSize: 3,
				Rep:         0x00,
				Data:        nil,
			},
		},
		{
			Name:  "testCase 5",
			input: "02",
			item: &Repetitive{
				SubItemSize: 3,
			},
			err: io.EOF,
			output: &Repetitive{
				SubItemSize: 3,
				Rep:         0x02,
				Data:        nil,
			},
		},
	}

	for _, tc := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(tc.input)
		rb := bytes.NewReader(input)
		f := NewRepetitive(tc.item)

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
					FRN:          1,
					DataItemName: "I000/010",
					Description:  "Test item",
				},
				Rep:  0x02,
				Data: []byte{0xab, 0xcd},
			},
			output: "I000/010:[02abcd]",
		},
		{
			Name: "testCase 2",
			input: Repetitive{
				Base: Base{
					FRN:          1,
					DataItemName: "I000/010",
					Description:  "Test item",
				},
				Rep: 0x02,
				SubItems: []SubItem{
					&SubItemFromTo{
						Name: "010-1",
						Pos:  BitPosition{From: 16, To: 9},
						Data: []byte{0xab},
					},
					&SubItemFromTo{
						Name: "010-2",
						Pos:  BitPosition{From: 8, To: 1},
						Data: []byte{0xcd},
					},
					&SubItemBit{
						Name: "010-3",
						Pos:  BitPosition{Bit: 8},
						Data: []byte{0x01},
					},

					&SubItemFromTo{
						Name: "010-1",
						Pos:  BitPosition{From: 16, To: 9},
						Data: []byte{0x12},
					},
					&SubItemFromTo{
						Name: "010-2",
						Pos:  BitPosition{From: 8, To: 1},
						Data: []byte{0x34},
					},
					&SubItemBit{
						Name: "010-3",
						Pos:  BitPosition{Bit: 8},
						Data: []byte{0x00},
					},
				},
			},
			output: "I000/010:[rep:02][010-1:ab][010-2:cd][010-3:01][010-1:12][010-2:34][010-3:00]",
		},
		{
			Name: "testCase 3",
			input: Repetitive{
				Base: Base{},
				Rep:  0,
				Data: nil,
			},
			output: ":[00]",
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
				Rep:  0x02,
				Data: []byte{0xab, 0xcd},
			},
			output: []byte{0x02, 0xab, 0xcd},
		},
		{
			Name: "testCase 2",
			input: Repetitive{
				Rep:  0,
				Data: nil,
			},
			output: []byte{0x00},
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

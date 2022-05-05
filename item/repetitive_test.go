package item

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestRepetitiveClone(t *testing.T) {
	// Arrange
	input := Repetitive{
		Base: Base{
			FRN:          1,
			DataItemName: "I000/010",
			Description:  "Test item",
			Type:         RepetitiveField,
		},
		SubItemSize: 3,
	}
	output := &Repetitive{
		Base: Base{
			FRN:          1,
			DataItemName: "I000/010",
			Description:  "Test item",
			Type:         RepetitiveField,
		},
		SubItemSize: 3,
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
			input: "03 8aaaaa 0bbbbb 8ccccc", // 8a 1000-1010 1010-1010 1010-1010
			item: &Repetitive{
				SubItemSize: 3,
				SubItems: []SubItem{
					{
						Type: BitField,
						Bit:  24,
					},
					{
						Type: FromToField,
						From: 20, To: 1,
					},
				},
			},
			err: nil,
			output: &Repetitive{
				SubItemSize: 3,
				Rep:         0x03,
				SubItems: []SubItem{
					{
						Type: BitField,
						Bit:  24,
						Data: []byte{0x01},
					},
					{
						Type: FromToField,
						From: 20, To: 1,
						Data: []byte{0x0a, 0xaa, 0xaa},
					},
					{
						Type: BitField,
						Bit:  24,
						Data: []byte{0x00},
					},
					{
						Type: FromToField,
						From: 20, To: 1,
						Data: []byte{0x0b, 0xbb, 0xbb},
					},
					{
						Type: BitField,
						Bit:  24,
						Data: []byte{0x01},
					},
					{
						Type: FromToField,
						From: 20, To: 1,
						Data: []byte{0x0c, 0xcc, 0xcc},
					},
				},
			},
		},
		{
			Name:  "testCase 2",
			input: "04 aaaaaa bbbbbb cccccc",
			item: &Repetitive{
				SubItemSize: 3,
			},
			err: io.EOF,
			output: &Repetitive{
				SubItemSize: 3,
				Rep:         0x04,
			},
		},
		{
			Name:  "testCase 3",
			input: "",
			item: &Repetitive{
				SubItemSize: 3,
			},
			err: io.EOF,
			output: &Repetitive{
				SubItemSize: 3,
				Rep:         0x00,
			},
		},
		{
			Name:  "testCase 4",
			input: "02",
			item: &Repetitive{
				SubItemSize: 3,
			},
			err: io.EOF,
			output: &Repetitive{
				SubItemSize: 3,
				Rep:         0x02,
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
				Rep: 0x02,
				SubItems: []SubItem{
					{
						Name: "010-1",
						From: 16, To: 9,
						Data: []byte{0xab},
					},
					{
						Name: "010-2",
						From: 8, To: 1,
						Data: []byte{0xcd},
					},
					{
						Name: "010-3",
						Bit:  8,
						Data: []byte{0x01},
					},

					{
						Name: "010-1",
						From: 16, To: 9,
						Data: []byte{0x12},
					},
					{
						Name: "010-2",
						From: 8, To: 1,
						Data: []byte{0x34},
					},
					{
						Name: "010-3",
						Bit:  8,
						Data: []byte{0x00},
					},
				},
			},
			output: "I000/010:[rep:02][010-1:ab][010-2:cd][010-3:01][010-1:12][010-2:34][010-3:00]",
		},
		{
			Name: "testCase 2",
			input: Repetitive{
				Base: Base{},
				Rep:  0,
			},
			output: ":[rep:00]",
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
*/

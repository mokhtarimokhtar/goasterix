package item

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestFixedReader(t *testing.T) {
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
			input: "01 02 03 04 05 06 07 08",
			item: &Fixed{
				Size: 8,
				SubItems: []SubItem{
					{
						Type: BitField,
						Bit:  57,
					},
					{
						Type: FromToField,
						From: 56,
						To:   25,
					},
					{
						Type: FromToField,
						From: 24,
						To:   1,
					},
				},
			},
			output: &Fixed{
				Size: 8,
				SubItems: []SubItem{
					{
						Type: BitField,
						Bit:  57,
						Data: []byte{0x01},
					},
					{
						Type: FromToField,
						From: 56,
						To:   25,
						Data: []byte{0x02, 0x03, 0x04, 0x05},
					},
					{
						Type: FromToField,
						From: 24,
						To:   1,
						Data: []byte{0x06, 0x07, 0x08},
					},
				},
			},
			err: nil,
		},
		{
			Name:  "testCase 2",
			input: "01 02 03 04 05 06 07",
			item: &Fixed{
				Size: 8,
			},
			output: &Fixed{
				Size: 8,
			},
			err: io.ErrUnexpectedEOF,
		},
		{
			Name:  "testCase 3",
			input: "ff",
			item: &Fixed{
				Size: 1,
				SubItems: []SubItem{
					{
						Type: FromToField,
						From: 1,
						To:   8,
					},
				},
			},
			output: &Fixed{
				Size:     1,
				SubItems: nil,
			},
			err: ErrSubDataFieldFormat,
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

func TestFixedClone(t *testing.T) {
	// Arrange
	input := Fixed{
		Base: Base{
			FRN:          1,
			DataItemName: "I000/010",
			Description:  "Test item",
			Type:         FixedField,
		},
		Size:     2,
		SubItems: nil,
	}
	output := &Fixed{
		Base: Base{
			FRN:          1,
			DataItemName: "I000/010",
			Description:  "Test item",
			Type:         FixedField,
		},
		Size:     2,
		SubItems: nil,
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
					FRN:          1,
					DataItemName: "I000/010",
					Description:  "Test item",
				},
				SubItems: []SubItem{
					{
						Name: "SUB-A",
						From: 16, To: 9,
						Data: []byte{0xab},
					},
					{
						Name: "SUB-B",
						From: 8, To: 1,
						Data: []byte{0xcd},
					},
					{
						Name: "SUB-C",
						Bit:  8,
						Data: []byte{0x01},
					},
				},
			},
			output: "I000/010:[SUB-A:ab][SUB-B:cd][SUB-C:01]",
		},
		{
			Name: "testCase 2",
			input: Fixed{
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

func TestFixedGetSubItems(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Fixed
		output []SubItem
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Fixed{
				Base: Base{
					FRN:          1,
					DataItemName: "I000/010",
					Description:  "Test item",
				},
				SubItems: []SubItem{
					{
						Name: "SUB-A",
						From: 16, To: 9,
						Data: []byte{0xab},
					},
					{
						Name: "SUB-B",
						From: 8, To: 1,
						Data: []byte{0xcd},
					},
					{
						Name: "SUB-C",
						Bit:  8,
						Data: []byte{0x01},
					},
				},
			},
			output: []SubItem{
				{
					Name: "SUB-A",
					From: 16, To: 9,
					Data: []byte{0xab},
				},
				{
					Name: "SUB-B",
					From: 8, To: 1,
					Data: []byte{0xcd},
				},
				{
					Name: "SUB-C",
					Bit:  8,
					Data: []byte{0x01},
				},
			},
		},
		{
			Name: "testCase 2",
			input: Fixed{
				Base: Base{},
			},
			output: nil,
		},
	}

	for _, tc := range dataSet {
		// Act
		s := tc.input.GetSubItems()

		// Assert
		if reflect.DeepEqual(s, tc.output) == false {
			t.Errorf(util.MsgFailInValue, tc.Name, s, tc.output)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, s, tc.output)
		}
	}
}

/*
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

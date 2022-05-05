package item

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestExplicitClone(t *testing.T) {
	// Arrange
	input := Explicit{
		Base: Base{
			FRN:          1,
			DataItemName: "I000/010",
			Description:  "Test item",
			Type:         ExplicitField,
		},
	}
	output := &Explicit{
		Base: Base{
			FRN:          1,
			DataItemName: "I000/010",
			Description:  "Test item",
			Type:         ExplicitField,
		},
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

func TestExplicitReader(t *testing.T) {
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
			input: "08 01 02 03 04 05 06 07",
			item:  &Explicit{},
			err:   nil,
			output: &Explicit{
				Len: 0x08,
				SubItems: []SubItem{
					{
						Type: FromToField,
						From: 56,
						To:   1,
						Data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
					},
				},
			},
		},
		{
			Name:  "testCase 2",
			input: "08 01 02 03 04 05 06",
			item:  &Explicit{},
			err:   io.ErrUnexpectedEOF,
			output: &Explicit{
				Len:      0x08,
				SubItems: nil,
			},
		},
		{
			Name:  "testCase 3",
			input: "",
			item:  &Explicit{},
			err:   io.EOF,
			output: &Explicit{
				Len:      0x00,
				SubItems: nil,
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
				Base: Base{
					FRN:          1,
					DataItemName: "I000/010",
					Description:  "Test item",
				},
				Len: 0x04,
				SubItems: []SubItem{
					{
						Type: FromToField,
						From: 24,
						To:   0,
						Data: []byte{0xab, 0xcd, 0xef},
					},
				},
			},
			output: "I000/010:[len:04][:abcdef]",
		},
		{
			Name: "testCase 2",
			input: Explicit{
				Base:     Base{},
				Len:      0,
				SubItems: nil,
			},
			output: ":[len:00]",
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

func TestExplicitGetSubItems(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Explicit
		output []SubItem
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Explicit{
				Base: Base{
					FRN:          1,
					DataItemName: "I000/010",
					Description:  "Test item",
				},
				Len: 0x03,
				SubItems: []SubItem{
					{
						Name: "SUB-A",
						From: 16, To: 1,
						Data: []byte{0xab, 0xcd},
					},
				},
			},
			output: []SubItem{
				{
					Name: "SUB-A",
					From: 16, To: 1,
					Data: []byte{0xab, 0xcd},
				},
			},
		},
		{
			Name: "testCase 2",
			input: Explicit{
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
				Len:  0x04,
				Data: []byte{0xab, 0xcd, 0xef},
			},
			output: []byte{0x04, 0xab, 0xcd, 0xef},
		},
		{
			Name: "testCase 2",
			input: Explicit{
				Len:  0,
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

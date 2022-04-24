package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"testing"
)

func TestSubItemBitReader(t *testing.T) {
	// setup
	type testCase struct {
		Name  string
		input []byte
		item  SubItemBit
		//output SubItemBit
		output []byte
		err    error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:   "testcase 1",
			input:  []byte{0x40},
			item:   SubItemBit{Pos: 7},
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 2",
			input:  []byte{0x00, 0x40},
			item:   SubItemBit{Pos: 7},
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 3",
			input:  []byte{0x10, 0x00},
			item:   SubItemBit{Pos: 13},
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 4",
			input:  []byte{0x10, 0x00, 0x00},
			item:   SubItemBit{Pos: 21},
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 5",
			input:  []byte{0x10, 0x00, 0x00, 0x00},
			item:   SubItemBit{Pos: 29},
			output: []byte{0x01},
			err:    nil,
		},
		/*{input: 0x80, pos: 8, output: []byte{0x01}},
		{input: 0x40, pos: 7, output: []byte{0x01}},
		{input: 0x20, pos: 6, output: []byte{0x01}},
		{input: 0x10, pos: 5, output: []byte{0x01}},
		{input: 0x08, pos: 4, output: []byte{0x01}},
		{input: 0x04, pos: 3, output: []byte{0x01}},
		{input: 0x02, pos: 2, output: []byte{0x01}},
		{input: 0x01, pos: 1, output: []byte{0x01}},
		{input: 0x7f, pos: 8, output: []byte{0x00}},
		{input: 0xbf, pos: 7, output: []byte{0x00}},
		{input: 0xdf, pos: 6, output: []byte{0x00}},
		{input: 0xef, pos: 5, output: []byte{0x00}},
		{input: 0xf7, pos: 4, output: []byte{0x00}},
		{input: 0xfb, pos: 3, output: []byte{0x00}},
		{input: 0xfd, pos: 2, output: []byte{0x00}},
		{input: 0xfe, pos: 1, output: []byte{0x00}},*/
	}

	for _, tc := range dataSet {
		// Act
		sub := tc.item
		err := sub.Reader(tc.input)

		// Assert
		if err != tc.err {
			t.Errorf(util.MsgFailInValue, tc.Name, err, tc.err)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, err, tc.err)
		}
		if bytes.Equal(sub.Data, tc.output) == false {
			t.Errorf(util.MsgFailInHex, tc.Name, sub.Data, tc.output)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, sub.Data, tc.output)
		}
	}
}

func TestOneBitReader(t *testing.T) {
	// setup
	type testCase struct {
		Name  string
		input byte
		//input  []byte
		pos uint8 // position of bit
		//output []byte
		output byte
	}
	// Arrange
	dataSet := []testCase{
		{input: 0x80, pos: 8, output: 0x01},
		{input: 0x40, pos: 7, output: 0x01},
		{input: 0x20, pos: 6, output: 0x01},
		{input: 0x10, pos: 5, output: 0x01},
		{input: 0x08, pos: 4, output: 0x01},
		{input: 0x04, pos: 3, output: 0x01},
		{input: 0x02, pos: 2, output: 0x01},
		{input: 0x01, pos: 1, output: 0x01},
		{input: 0x7f, pos: 8, output: 0x00},
		{input: 0xbf, pos: 7, output: 0x00},
		{input: 0xdf, pos: 6, output: 0x00},
		{input: 0xef, pos: 5, output: 0x00},
		{input: 0xf7, pos: 4, output: 0x00},
		{input: 0xfb, pos: 3, output: 0x00},
		{input: 0xfd, pos: 2, output: 0x00},
		{input: 0xfe, pos: 1, output: 0x00},
	}

	for _, tc := range dataSet {
		// Act
		res := OneBitReader(tc.input, tc.pos)

		// Assert
		if res != tc.output {
			//if bytes.Equal(res != tc.output {
			t.Errorf(util.MsgFailInHex, tc.Name, res, tc.output)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, res, tc.output)
		}
	}
}

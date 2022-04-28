package item

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"reflect"
	"testing"
)

//func BitShiftGenerator(max int, size int) [][]byte {
//	var table = make([][]byte, 16)
//
//	for pos := 1; pos <= max; pos++ {
//		//var value = make([]byte, size)
//		table[pos-1] = make([]byte, size)
//		val := math.Pow(2, float64(pos-1))
//		//tmp := data[j] | 0x01<<pos
//		fmt.Printf("pos=%v %v\n", pos, uint16(val))
//		binary.BigEndian.PutUint16(table[pos-1], uint16(val))
//		//fmt.Printf("pos=%v %08b-%08b\n", pos, value[0], value[1])
//
//		//table[pos-1] = append(table[pos-1], value[0], value[1])
//		//data[j-1] = uint8(val)
//		//fmt.Printf("pos=%v %08b\n", pos, data)
//	}
//	return table
//}

func TestSubItemBitReader(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  []byte
		Pos    uint8
		output []byte
		err    error
	}

	// Arrange
	dataSet := []testCase{
		{
			Name:   "testcase 1",
			input:  []byte{0x40},
			Pos:    7,
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 2",
			input:  []byte{0x00, 0x40},
			Pos:    7,
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 3",
			input:  []byte{0x10, 0x00},
			Pos:    13,
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 4",
			input:  []byte{0x10, 0x00, 0x00},
			Pos:    21,
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 5",
			input:  []byte{0x10, 0x00, 0x00, 0x00},
			Pos:    29,
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 6",
			input:  []byte{0x80, 0x00, 0x00},
			Pos:    24,
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 7",
			input:  []byte{0x80},
			Pos:    8,
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 8",
			input:  []byte{0x01},
			Pos:    1,
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 9",
			input:  []byte{0x00, 0x01, 0x00},
			Pos:    9,
			output: []byte{0x01},
			err:    nil,
		},
		{
			Name:   "testcase 10",
			input:  []byte{0x00, 0x00, 0x01},
			Pos:    1,
			output: []byte{0x01},
			err:    nil,
		},
	}

	for _, tc := range dataSet {
		// Arrange
		sub := new(SubItemBit)
		sub.Pos.Bit = tc.Pos

		// Act
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

func TestSubItemFromToReader(t *testing.T) {
	type testCase struct {
		Name   string
		input  []byte
		From   uint8
		To     uint8
		err    error
		output []byte
	}

	dataSet := []testCase{
		{
			Name:  "testcase 1",
			input: []byte{0xff},
			From:  6, To: 2,
			err:    nil,
			output: []byte{0x1f},
		},
		{
			Name:  "testcase 2",
			input: []byte{0xff},
			From:  8, To: 1,
			err:    nil,
			output: []byte{0xff},
		},
		{
			Name:  "testcase 3",
			input: []byte{0xff},
			From:  3, To: 2,
			err:    nil,
			output: []byte{0x03},
		},
		{
			Name:  "testcase 4",
			input: []byte{0xff, 0xff},
			From:  6, To: 2,
			err:    nil,
			output: []byte{0x1f},
		},
		{
			Name:  "testcase 5",
			input: []byte{0xff, 0xff},
			From:  9, To: 2,
			err:    nil,
			output: []byte{0xff},
		},
		{
			Name:  "testcase 6",
			input: []byte{0xff, 0xff},
			From:  11, To: 2,
			err:    nil,
			output: []byte{0x03, 0xff},
		},
		{
			Name:  "testcase 7",
			input: []byte{0xff, 0xff},
			From:  11, To: 9,
			err:    nil,
			output: []byte{0x07},
		},
		{
			Name:  "testcase 8",
			input: []byte{0xff, 0xff},
			From:  16, To: 1,
			err:    nil,
			output: []byte{0xff, 0xff},
		},
		{
			Name:  "testcase 9",
			input: []byte{0xff, 0xff, 0xff},
			From:  6, To: 2,
			err:    nil,
			output: []byte{0x1f},
		},
		{
			Name:  "testcase 10",
			input: []byte{0xff, 0xff, 0xff},
			From:  9, To: 2,
			err:    nil,
			output: []byte{0xff},
		},
		{
			Name:  "testcase 11",
			input: []byte{0xff, 0xff, 0xff},
			From:  20, To: 2,
			err:    nil,
			output: []byte{0x07, 0xff, 0xff},
		},
		{
			Name:  "testcase 12",
			input: []byte{0xff, 0xff, 0xff, 0xff},
			From:  6, To: 2,
			err:    nil,
			output: []byte{0x1f},
		},
		{
			Name:  "testcase 13",
			input: []byte{0xff, 0xff, 0xff, 0xff},
			From:  27, To: 8,
			err:    nil,
			output: []byte{0x0f, 0xff, 0xff},
		},
		{
			Name:  "testcase 14",
			input: []byte{0xff, 0xff, 0xff, 0xff},
			From:  8, To: 27,
			err:    ErrSubDataFieldFormat,
			output: nil,
		},
	}
	for _, tc := range dataSet {
		// Arrange
		sub := new(SubItemFromTo)
		sub.Pos.From = tc.From
		sub.Pos.To = tc.To

		// Act
		err := sub.Reader(tc.input)

		// Assert
		if err != tc.err {
			t.Errorf(util.MsgFailInValue, tc.Name, err, tc.err)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, err, tc.err)
		}

		if len(sub.Data) != len(tc.output) {
			t.Errorf(util.MsgFailInHex, tc.Name, len(sub.Data), len(tc.output))
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, len(sub.Data), len(tc.output))
		}

		if bytes.Equal(sub.Data, tc.output) == false {
			t.Errorf(util.MsgFailInHex, tc.Name, sub.Data, tc.output)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, sub.Data, tc.output)
		}
	}
}

func TestFromToBitReader(t *testing.T) {
	type testCase struct {
		Name   string
		input  []byte
		from   uint8
		to     uint8
		err    error
		output []byte
	}

	dataSet := []testCase{
		{
			Name:   "testcase 1",
			input:  []byte{0x7e}, // 0[111 111]0
			from:   7,
			to:     2,
			output: []byte{0x3f},
		},
		{
			Name:   "testcase 2",
			input:  []byte{0x06},
			from:   3,
			to:     2,
			output: []byte{0x03},
		},
		{
			Name:   "testcase 3",
			input:  []byte{0xff},
			from:   8,
			to:     1,
			output: []byte{0xff},
		},
		{
			Name:   "testcase 4",
			input:  []byte{0x0f, 0xf0},
			from:   12,
			to:     5,
			output: []byte{0xff},
		},
		{
			Name:   "testcase 5",
			input:  []byte{0x00, 0x3e},
			from:   6,
			to:     2,
			output: []byte{0x1f},
		},
		{
			Name:   "testcase 6",
			input:  []byte{0x01, 0xfe},
			from:   9,
			to:     2,
			output: []byte{0xff},
		},
		{
			Name:   "testcase 7",
			input:  []byte{0x07, 0xfe},
			from:   11,
			to:     2,
			output: []byte{0x03, 0xff},
		},
		{
			Name:   "testcase 8",
			input:  []byte{0x07, 0x00},
			from:   11,
			to:     9,
			output: []byte{0x07},
		},
		{
			Name:   "testcase 9",
			input:  []byte{0xff, 0xff},
			from:   16,
			to:     1,
			output: []byte{0xff, 0xff},
		},
		{
			Name:   "testcase 10",
			input:  []byte{0x00, 0x00, 0x3e},
			from:   6,
			to:     2,
			output: []byte{0x1f},
		},
		{
			Name:   "testcase 11",
			input:  []byte{0x00, 0x01, 0xfe},
			from:   9,
			to:     2,
			output: []byte{0xff},
		},
		{
			Name:   "testcase 12",
			input:  []byte{0x0f, 0xff, 0xfe},
			from:   20,
			to:     2,
			output: []byte{0x07, 0xff, 0xff},
		},
		{
			Name:   "testcase 13",
			input:  []byte{0x00, 0x00, 0x00, 0x3e},
			from:   6,
			to:     2,
			output: []byte{0x1f},
		},
		{
			Name:   "testcase 14",
			input:  []byte{0x07, 0xff, 0xff, 0x80},
			from:   27,
			to:     8,
			output: []byte{0x0f, 0xff, 0xff},
		},
		{
			Name:   "testcase 15",
			input:  []byte{0xff, 0xff, 0xff, 0xff},
			from:   32,
			to:     1,
			output: []byte{0xff, 0xff, 0xff, 0xff},
		},
		{
			Name:   "testcase 16",
			input:  []byte{0xff, 0xff, 0xff},
			from:   24,
			to:     1,
			output: []byte{0xff, 0xff, 0xff},
		},
		{
			Name:   "testcase 17",
			input:  []byte{0xff, 0xff, 0xff, 0xff, 0xff},
			from:   40,
			to:     1,
			output: []byte{0xff, 0xff, 0xff, 0xff, 0xff},
		},
		{
			Name:   "testcase 18",
			input:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			from:   48,
			to:     1,
			output: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
		{
			Name:   "testcase 19",
			input:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			from:   56,
			to:     1,
			output: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
		{
			Name:   "testcase 20",
			input:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			from:   64,
			to:     1,
			output: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
		{
			Name:   "testcase 21",
			input:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			from:   2,
			to:     10,
			output: nil,
			err:    ErrSubDataFieldFormat,
		},
	}

	for _, tc := range dataSet {
		// Act
		res, err := FromToBitReader(tc.input, tc.from, tc.to)

		// Assert
		if err != tc.err {
			t.Errorf(util.MsgFailInValue, tc.Name, err, tc.err)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, err, tc.err)
		}

		if bytes.Equal(res, tc.output) == false {
			t.Errorf(util.MsgFailInHex, tc.Name, res, tc.output)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, res, tc.output)
		}
	}
}

func TestUint64ToByteLess(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  uint64
		nb     uint8
		output []byte
	}
	dataSet := []testCase{
		{
			input:  0xaabbccddaabbccdd,
			nb:     8,
			output: []byte{0xaa, 0xbb, 0xcc, 0xdd, 0xaa, 0xbb, 0xcc, 0xdd},
		},
		{
			input:  0x00bbccddaabbccdd,
			nb:     7,
			output: []byte{0xbb, 0xcc, 0xdd, 0xaa, 0xbb, 0xcc, 0xdd},
		},
		{
			input:  0x0000ccddaabbccdd,
			nb:     6,
			output: []byte{0xcc, 0xdd, 0xaa, 0xbb, 0xcc, 0xdd},
		},
		{
			input:  0x000000ddaabbccdd,
			nb:     5,
			output: []byte{0xdd, 0xaa, 0xbb, 0xcc, 0xdd},
		},
		{
			input:  0x00000000aabbccdd,
			nb:     4,
			output: []byte{0xaa, 0xbb, 0xcc, 0xdd},
		},
		{
			input:  0x0000000000bbccdd,
			nb:     3,
			output: []byte{0xbb, 0xcc, 0xdd},
		},
		{
			input:  0x00000000000ccdd,
			nb:     2,
			output: []byte{0xcc, 0xdd},
		},
		{
			input:  0x0000000000000dd,
			nb:     1,
			output: []byte{0xdd},
		},
	}
	for _, tc := range dataSet {
		// Act
		res := Uint64ToByteLess(tc.input, tc.nb)

		// Assert
		if bytes.Equal(res, tc.output) == false {
			t.Errorf(util.MsgFailInHex, tc.Name, res, tc.output)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, res, tc.output)
		}
	}
}

func TestGetSubItem(t *testing.T) {
	type testCase struct {
		Name   string
		input  SubItem
		err    error
		output SubItem
	}
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: &SubItemBit{
				Type: BitField,
			},
			err: nil,
			output: &SubItemBit{
				Type: BitField,
			},
		},
		{
			Name: "testcase 2",
			input: &SubItemFromTo{
				Type: FromToField,
			},
			err: nil,
			output: &SubItemFromTo{
				Type: FromToField,
			},
		},
		{
			Name: "testcase 3",
			input: &SubItemFromTo{
				Type: ExtendedField,
			},
			err:    ErrSubDataFieldUnknown,
			output: nil,
		},
	}
	for _, tc := range dataSet {
		// Act
		sub, err := GetSubItem(tc.input)

		// Assert
		if err != tc.err {
			t.Errorf(util.MsgFailInValue, tc.Name, err, tc.err)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, err, tc.err)
		}
		if reflect.DeepEqual(sub, tc.output) == false {
			t.Errorf(util.MsgFailInValue, tc.Name, sub, tc.output)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, sub, tc.output)
		}
	}

}

/*
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
*/

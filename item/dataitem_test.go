package item

import (
	"github.com/mokhtarimokhtar/goasterix/util"
	"testing"
)

func TestTypeFieldString(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  TypeField
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:   "testcase 1",
			input:  FixedField,
			output: "Fixed",
		},
		{
			Name:   "testcase 2",
			input:  ExtendedField,
			output: "Extended",
		},
		{
			Name:   "testcase 3",
			input:  CompoundField,
			output: "Compound",
		},
		{
			Name:   "testcase 4",
			input:  RepetitiveField,
			output: "Repetitive",
		},
		{
			Name:   "testcase 5",
			input:  ExplicitField,
			output: "Explicit",
		},
		{
			Name:   "testcase 6",
			input:  SPField,
			output: "SP",
		},
		{
			Name:   "testcase 7",
			input:  REField,
			output: "RE",
		},
		{
			Name:   "testcase 8",
			input:  RFSField,
			output: "RFS",
		},
		{
			Name:   "testcase 9",
			input:  SpareField,
			output: "Spare",
		},
		{
			Name:   "testcase 10",
			input:  BitField,
			output: "Bit",
		},
		{
			Name:   "testcase 11",
			input:  FromToField,
			output: "FromToBit",
		},
		{
			Name:   "testcase 12",
			input:  0,
			output: "",
		},
	}
	for _, tc := range dataSet {
		// Act
		res := tc.input.String()

		// Assert
		if res != tc.output {
			t.Errorf(util.MsgFailInString, tc.Name, res, tc.output)
		} else {
			t.Logf(util.MsgSuccessInString, tc.Name, res, tc.output)
		}
	}
}

/*
func TestPayloadDataItemPayload(t *testing.T) {
	// setup
	type testCase struct {
		Name     string
		dataItem DataItem
		output   []byte
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			dataItem: &Fixed{
				Base: Base{
					Type: FixedField,
				},
				Data: []byte{0x01, 0x02, 0x03},
				Size: 3,
			},
			output: []byte{0x01, 0x02, 0x03},
		},
		{
			Name: "testcase 2",
			dataItem: &Extended{
				Base: Base{
					Type: ExtendedField,
				},
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
				Primary:           []byte{0x01},
				Secondary:         []byte{0x01, 0x03, 0x04},
			},
			output: []byte{0x01, 0x01, 0x03, 0x04},
		},
		{
			Name: "testcase 3",
			dataItem: &Repetitive{
				Base: Base{
					Type: RepetitiveField,
				},
				SubItemSize: 2,
				Rep:         3,
				Data:        []byte{0x01, 0x03, 0x02, 0x01, 0x03, 0x02},
			},
			output: []byte{0x03, 0x01, 0x03, 0x02, 0x01, 0x03, 0x02},
		},
	}
	for _, tc := range dataSet {
		p := new(PayloadDataItem)

		// Act
		p.Payload(tc.dataItem)

		// Assert
		if bytes.Equal(p.Data, tc.output) == false {
			t.Errorf(util.MsgFailInHex, tc.Name, p.Data, tc.output)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, p.Data, tc.output)
		}
	}
}
*/

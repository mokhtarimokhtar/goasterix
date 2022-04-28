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

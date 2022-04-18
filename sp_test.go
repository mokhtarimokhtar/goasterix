package goasterix

import (
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"testing"
)

func TestSpecialPurposeString(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  SpecialPurpose
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: SpecialPurpose{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.SP,
				},
				Len:  0x04,
				Data: []byte{0xab, 0xcd, 0xef},
			},
			output: "I000/010:04abcdef",
		},
		{
			Name: "testcase 2",
			input: SpecialPurpose{
				MetaItem: MetaItem{},
				Len:      0,
				Data:     nil,
			},
			output: ":00",
		},
	}

	for _, row := range dataSet {
		// Act
		s := row.input.String()

		// Assert
		if s != row.output {
			t.Errorf(util.MsgFailInValue, row.Name, s, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, s, row.output)
		}
	}
}

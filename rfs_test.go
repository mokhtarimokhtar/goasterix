package goasterix

import (
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"testing"
)

func TestRandomFieldString(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  RandomField
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: RandomField{
				FRN: 1,
				Field: &Fixed{
					MetaItem: MetaItem{
						FRN:         1,
						DataItem:    "I000/010",
						Description: "Test item",
						Type:        uap.Fixed,
					},
					Data: []byte{0xab, 0xcd},
				},
			},
			output: "FRN:01 I000/010:abcd",
		},
		{
			Name: "testcase 2",
			input: RandomField{
				FRN: 0,
				Field: &Fixed{
					MetaItem: MetaItem{},
					Data:     nil,
				},
			},
			output: "FRN:00 :",
		},
		{
			Name: "testcase 3",
			input: RandomField{
				FRN: 3,
				Field: &Extended{
					MetaItem: MetaItem{
						FRN:         3,
						DataItem:    "I000/030",
						Description: "Test item",
						Type:        uap.Extended,
					},
					Primary:   []byte{0xc1},
					Secondary: []byte{0xab, 0xcd},
				},
			},
			output: "FRN:03 I000/030:c1abcd",
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

func TestRandomFieldSequencingString(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  RandomFieldSequencing
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: RandomFieldSequencing{
				MetaItem: MetaItem{
					FRN:         0,
					DataItem:    "I000/000",
					Description: "Test item",
					Type:        uap.RFS,
				},
				N: 2,
				Sequence: []RandomField{
					{
						FRN: 1,
						Field: &Fixed{
							MetaItem: MetaItem{
								FRN:         1,
								DataItem:    "I000/010",
								Description: "Test item",
								Type:        uap.Fixed,
							},
							Data: []byte{0xab, 0xcd},
						},
					},
					{
						FRN: 3,
						Field: &Extended{
							MetaItem: MetaItem{
								FRN:         3,
								DataItem:    "I000/030",
								Description: "Test item",
								Type:        uap.Extended,
							},
							Primary:   []byte{0xc1},
							Secondary: []byte{0xab, 0xcd},
						},
					},
				},
			},
			output: "I000/000:[N:02][FRN:01 I000/010:abcd][FRN:03 I000/030:c1abcd]",
		},
		{
			Name: "testcase 2",
			input: RandomFieldSequencing{
				MetaItem: MetaItem{},
				N:        0,
				Sequence: nil,
			},
			output: ":[N:00]",
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

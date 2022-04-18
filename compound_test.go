package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"testing"
)

func TestCompoundString(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Compound
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: Compound{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/000",
					Description: "Test item",
					Type:        uap.Compound,
				},
				Primary: []byte{0xaa},
				Secondary: []Item{
					&Fixed{
						MetaItem: MetaItem{
							FRN:         1,
							DataItem:    "I000/010",
							Description: "Test item",
							Type:        uap.Fixed,
						},
						Data: []byte{0xab, 0xcd},
					},
					&Extended{
						MetaItem: MetaItem{
							FRN:         3,
							DataItem:    "I000/030",
							Description: "Test item",
							Type:        uap.Extended,
						},
						Primary:   []byte{0xc1},
						Secondary: []byte{0xab, 0xcd},
					},
					&Explicit{
						MetaItem: MetaItem{
							FRN:         5,
							DataItem:    "I000/050",
							Description: "Test item",
							Type:        uap.Explicit,
						},
						Len:  0x04,
						Data: []byte{0xab, 0xcd, 0xef},
					},
					&Repetitive{
						MetaItem: MetaItem{
							FRN:         7,
							DataItem:    "I000/070",
							Description: "Test item",
							Type:        uap.Repetitive,
						},
						Rep:  0x02,
						Data: []byte{0xab, 0xcd},
					},
				},
			},
			output: "I000/000:[primary:aa][I000/010:abcd][I000/030:c1abcd][I000/050:04abcdef][I000/070:02abcd]",
		},
		{
			Name: "testcase 2",
			input: Compound{
				MetaItem:  MetaItem{},
				Primary:   nil,
				Secondary: nil,
			},
			output: ":[primary:]",
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

func TestCompoundPayload(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Compound
		output []byte
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: Compound{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/000",
					Description: "Test item",
					Type:        uap.Compound,
				},
				Primary: []byte{0xaa},
				Secondary: []Item{
					&Fixed{
						MetaItem: MetaItem{
							FRN:         1,
							DataItem:    "I000/010",
							Description: "Test item",
							Type:        uap.Fixed,
						},
						Data: []byte{0xab, 0xcd},
					},
					&Extended{
						MetaItem: MetaItem{
							FRN:         3,
							DataItem:    "I000/030",
							Description: "Test item",
							Type:        uap.Extended,
						},
						Primary:   []byte{0xc1},
						Secondary: []byte{0xab, 0xcd},
					},
					&Explicit{
						MetaItem: MetaItem{
							FRN:         5,
							DataItem:    "I000/050",
							Description: "Test item",
							Type:        uap.Explicit,
						},
						Len:  0x04,
						Data: []byte{0xab, 0xcd, 0xef},
					},
					&Repetitive{
						MetaItem: MetaItem{
							FRN:         7,
							DataItem:    "I000/070",
							Description: "Test item",
							Type:        uap.Repetitive,
						},
						Rep:  0x02,
						Data: []byte{0xab, 0xcd},
					},
				},
			},
			output: []byte{0xaa, 0xab, 0xcd, 0xc1, 0xab, 0xcd, 0x04, 0xab, 0xcd, 0xef, 0x02, 0xab, 0xcd},
		},
		{
			Name: "testcase 2",
			input: Compound{
				MetaItem:  MetaItem{},
				Primary:   nil,
				Secondary: nil,
			},
			output: nil,
		},
	}

	for _, row := range dataSet {
		// Act
		res := row.input.Payload()

		// Assert
		if bytes.Equal(res, row.output) == false {
			t.Errorf(util.MsgFailInHex, row.Name, res, row.output)
		} else {
			t.Logf(util.MsgSuccessInHex, row.Name, res, row.output)
		}
	}
}

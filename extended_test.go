package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"testing"
)

func TestExtendedString(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Extended
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: Extended{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Extended,
				},
				Primary:   []byte{0xc1},
				Secondary: []byte{0xab, 0xcd},
			},
			output: "I000/010:c1abcd",
		},
		{
			Name: "testcase 2",
			input: Extended{
				MetaItem:  MetaItem{},
				Primary:   nil,
				Secondary: nil,
			},
			output: ":",
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

func TestExtendedPayload(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Extended
		output []byte
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: Extended{
				MetaItem: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Extended,
				},
				Primary:   []byte{0xc1},
				Secondary: []byte{0xab, 0xcd},
			},
			output: []byte{0xc1, 0xab, 0xcd},
		},
		{
			Name: "testcase 2",
			input: Extended{
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

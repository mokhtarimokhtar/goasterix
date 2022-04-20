package goasterix

import (
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"reflect"
	"testing"
)

func TestNewBase(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  uap.DataField
		output Base
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: uap.DataField{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.Fixed,
				Fixed:       uap.FixedField{Size: 1},
			},
			output: Base{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.Fixed,
			},
		},
		{
			Name: "testCase 2",
			input: uap.DataField{
				FRN:         0,
				DataItem:    "",
				Description: "",
				Type:        0,
				Fixed:       uap.FixedField{},
			},
			output: Base{},
		},
		{
			Name: "testCase 3",
			input: uap.DataField{
				FRN:         3,
				DataItem:    "I000/030",
				Description: "Test item",
				Type:        uap.Extended,
				Extended: uap.ExtendedField{
					PrimarySize:   1,
					SecondarySize: 2,
				},
			},
			output: Base{
				FRN:         3,
				DataItem:    "I000/030",
				Description: "Test item",
				Type:        uap.Extended,
			},
		},
		{
			Name: "testCase 4",
			input: uap.DataField{
				FRN:         4,
				DataItem:    "I000/040",
				Description: "Test item",
				Type:        uap.Explicit,
			},
			output: Base{
				FRN:         4,
				DataItem:    "I000/040",
				Description: "Test item",
				Type:        uap.Explicit,
			},
		},
		{
			Name: "testCase 5",
			input: uap.DataField{
				FRN:         5,
				DataItem:    "I000/050",
				Description: "Test item",
				Type:        uap.Repetitive,
				Repetitive:  uap.RepetitiveField{SubItemSize: 2},
			},
			output: Base{
				FRN:         5,
				DataItem:    "I000/050",
				Description: "Test item",
				Type:        uap.Repetitive,
			},
		},
		{
			Name: "testCase 6",
			input: uap.DataField{
				FRN:         6,
				DataItem:    "I000/060",
				Description: "Test item",
				Type:        uap.Compound,
				Compound:    []uap.DataField{},
			},
			output: Base{
				FRN:         6,
				DataItem:    "I000/060",
				Description: "Test item",
				Type:        uap.Compound,
			},
		},
	}

	for _, row := range dataSet {
		// Arrange
		m := Base{}
		// Act
		m.NewBase(row.input)

		// Assert
		if reflect.DeepEqual(m, row.output) == false {
			t.Errorf(util.MsgFailInValue, row.Name, m, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, m, row.output)
		}
	}

}

func TestBaseFrn(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Base
		output uint8
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Base{
					FRN:         7,
					DataItem:    "I000/070",
					Description: "Test item",
					Type:        uap.Fixed,
				},
			output: 7,
		},
		{
			Name: "testCase 2",
			input: Base{},
			output: 0,
		},
	}

	for _, row := range dataSet {
		// Act
		res := row.input.Frn()

		// Assert
		if res != row.output {
			t.Errorf(util.MsgFailInValue, row.Name, res, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, res, row.output)
		}
	}
}

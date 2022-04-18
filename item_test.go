package goasterix

import (
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"reflect"
	"testing"
)

func TestNewMetaItem(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  uap.DataField
		output MetaItem
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: uap.DataField{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.Fixed,
				Fixed:       uap.FixedField{Size: 1},
			},
			output: MetaItem{
				FRN:         1,
				DataItem:    "I000/010",
				Description: "Test item",
				Type:        uap.Fixed,
			},
		},
		{
			Name: "testcase 2",
			input: uap.DataField{
				FRN:         0,
				DataItem:    "",
				Description: "",
				Type:        0,
				Fixed:       uap.FixedField{},
			},
			output: MetaItem{},
		},
		{
			Name: "testcase 3",
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
			output: MetaItem{
				FRN:         3,
				DataItem:    "I000/030",
				Description: "Test item",
				Type:        uap.Extended,
			},
		},
		{
			Name: "testcase 4",
			input: uap.DataField{
				FRN:         4,
				DataItem:    "I000/040",
				Description: "Test item",
				Type:        uap.Explicit,
				Explicit:    uap.ExplicitField{},
			},
			output: MetaItem{
				FRN:         4,
				DataItem:    "I000/040",
				Description: "Test item",
				Type:        uap.Explicit,
			},
		},
		{
			Name: "testcase 5",
			input: uap.DataField{
				FRN:         5,
				DataItem:    "I000/050",
				Description: "Test item",
				Type:        uap.Repetitive,
				Repetitive:  uap.RepetitiveField{SubItemSize: 2},
			},
			output: MetaItem{
				FRN:         5,
				DataItem:    "I000/050",
				Description: "Test item",
				Type:        uap.Repetitive,
			},
		},
		{
			Name: "testcase 6",
			input: uap.DataField{
				FRN:         6,
				DataItem:    "I000/060",
				Description: "Test item",
				Type:        uap.Compound,
				Compound:    []uap.DataField{},
			},
			output: MetaItem{
				FRN:         6,
				DataItem:    "I000/060",
				Description: "Test item",
				Type:        uap.Compound,
			},
		},
	}

	for _, row := range dataSet {
		// Arrange
		m := MetaItem{}
		// Act
		m.NewMetaItem(row.input)

		// Assert
		if reflect.DeepEqual(m, row.output) == false {
			t.Errorf(util.MsgFailInValue, row.Name, m, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, m, row.output)
		}
	}

}

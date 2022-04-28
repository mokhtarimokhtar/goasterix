package item

import (
	"github.com/mokhtarimokhtar/goasterix/util"
	"testing"
)

func TestBaseGetFrn(t *testing.T) {
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
				FRN:          7,
				DataItemName: "I000/070",
				Description:  "Test item",
				Type:         FixedField,
			},
			output: 7,
		},
		{
			Name:   "testCase 2",
			input:  Base{},
			output: 0,
		},
	}

	for _, row := range dataSet {
		// Act
		res := row.input.GetFrn()

		// Assert
		if res != row.output {
			t.Errorf(util.MsgFailInValue, row.Name, res, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, res, row.output)
		}
	}
}

func TestBaseGetType(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Base
		output TypeField
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Base{
				FRN:          7,
				DataItemName: "I000/070",
				Description:  "Test item",
				Type:         FixedField,
			},
			output: FixedField,
		},
		{
			Name:   "testCase 2",
			input:  Base{},
			output: 0,
		},
	}

	for _, row := range dataSet {
		// Act
		res := row.input.GetType()

		// Assert
		if res != row.output {
			t.Errorf(util.MsgFailInValue, row.Name, res, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, res, row.output)
		}
	}
}

func TestBaseGetDataItemName(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Base
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Base{
				FRN:          7,
				DataItemName: "I000/070",
				Description:  "Test item",
				Type:         FixedField,
			},
			output: "I000/070",
		},
		{
			Name:   "testCase 2",
			input:  Base{},
			output: "",
		},
	}

	for _, row := range dataSet {
		// Act
		res := row.input.GetDataItemName()

		// Assert
		if res != row.output {
			t.Errorf(util.MsgFailInValue, row.Name, res, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, res, row.output)
		}
	}
}

func TestBaseGetDescription(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Base
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Base{
				FRN:          7,
				DataItemName: "I000/070",
				Description:  "Test item",
				Type:         FixedField,
			},
			output: "Test item",
		},
		{
			Name:   "testCase 2",
			input:  Base{},
			output: "",
		},
	}

	for _, row := range dataSet {
		// Act
		res := row.input.GetDescription()

		// Assert
		if res != row.output {
			t.Errorf(util.MsgFailInValue, row.Name, res, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, res, row.output)
		}
	}
}

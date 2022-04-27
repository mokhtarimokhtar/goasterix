package item

/*
func TestNewBase(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  _uap.DataField
		output Base
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: _uap.DataField{
				FRN:         1,
				DataItemName:    "I000/010",
				Description: "Test item",
				Type:        _uap.Fixed,
				SizeItem:    _uap.SizeField{ForFixed: 1},
			},
			output: Base{
				FRN:         1,
				DataItemName:    "I000/010",
				Description: "Test item",
				Type:        _uap.Fixed,
			},
		},
		{
			Name: "testCase 2",
			input: _uap.DataField{
				FRN:         0,
				DataItemName:    "",
				Description: "",
				Type:        0,
				SizeItem:    _uap.SizeField{},
			},
			output: Base{},
		},
		{
			Name: "testCase 3",
			input: _uap.DataField{
				FRN:         3,
				DataItemName:    "I000/030",
				Description: "Test item",
				Type:        _uap.Extended,
				SizeItem: _uap.SizeField{
					ForExtendedPrimary:   1,
					ForExtendedSecondary: 2,
				},
			},
			output: Base{
				FRN:         3,
				DataItemName:    "I000/030",
				Description: "Test item",
				Type:        _uap.Extended,
			},
		},
		{
			Name: "testCase 4",
			input: _uap.DataField{
				FRN:         4,
				DataItemName:    "I000/040",
				Description: "Test item",
				Type:        _uap.Explicit,
			},
			output: Base{
				FRN:         4,
				DataItemName:    "I000/040",
				Description: "Test item",
				Type:        _uap.Explicit,
			},
		},
		{
			Name: "testCase 5",
			input: _uap.DataField{
				FRN:         5,
				DataItemName:    "I000/050",
				Description: "Test item",
				Type:        _uap.Repetitive,
				SizeItem:    _uap.SizeField{ForRepetitive: 2},
			},
			output: Base{
				FRN:         5,
				DataItemName:    "I000/050",
				Description: "Test item",
				Type:        _uap.Repetitive,
			},
		},
		{
			Name: "testCase 6",
			input: _uap.DataField{
				FRN:         6,
				DataItemName:    "I000/060",
				Description: "Test item",
				Type:        _uap.Compound,
				Compound:    []_uap.DataField{},
			},
			output: Base{
				FRN:         6,
				DataItemName:    "I000/060",
				Description: "Test item",
				Type:        _uap.Compound,
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
				DataItemName:    "I000/070",
				Description: "Test item",
				Type:        _uap.Fixed,
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
		res := row.input.Frn()

		// Assert
		if res != row.output {
			t.Errorf(util.MsgFailInValue, row.Name, res, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, res, row.output)
		}
	}
}
*/

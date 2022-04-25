package goasterix

/*
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
			Name: "testCase 1",
			input: RandomField{
				FRN: 1,
				Field: &Fixed{
					Base: Base{
						FRN:         1,
						DataItem:    "I000/010",
						Description: "Test item",
						Type:        _uap.Fixed,
					},
					Data: []byte{0xab, 0xcd},
				},
			},
			output: "FRN:01 I000/010:abcd",
		},
		{
			Name: "testCase 2",
			input: RandomField{
				FRN: 0,
				Field: &Fixed{
					Base: Base{},
					Data: nil,
				},
			},
			output: "FRN:00 :",
		},
		{
			Name: "testCase 3",
			input: RandomField{
				FRN: 3,
				Field: &Extended{
					Base: Base{
						FRN:         3,
						DataItem:    "I000/030",
						Description: "Test item",
						Type:        _uap.Extended,
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

func TestRandomFieldPayload(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  RandomField
		output []byte
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: RandomField{
				FRN: 1,
				Field: &Fixed{
					Base: Base{
						FRN:         2,
						DataItem:    "I000/020",
						Description: "Test item",
						Type:        _uap.Fixed,
					},
					Data: []byte{0xab, 0xcd},
				},
			},
			output: []byte{0x01, 0xab, 0xcd},
		},
		{
			Name: "testCase 2",
			input: RandomField{
				FRN: 0,
				Field: &Fixed{
					Base: Base{},
					Data: nil,
				},
			},
			output: []byte{0x00},
		},
		{
			Name: "testCase 3",
			input: RandomField{
				FRN: 16,
				Field: &Extended{
					Base: Base{
						FRN:         3,
						DataItem:    "I000/030",
						Description: "Test item",
						Type:        _uap.Extended,
					},
					Primary:   []byte{0xc1},
					Secondary: []byte{0xab, 0xcd},
				},
			},
			output: []byte{0x10, 0xc1, 0xab, 0xcd},
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
			Name: "testCase 1",
			input: RandomFieldSequencing{
				Base: Base{
					FRN:         0,
					DataItem:    "I000/000",
					Description: "Test item",
					Type:        _uap.RFS,
				},
				N: 2,
				Sequence: []RandomField{
					{
						FRN: 1,
						Field: &Fixed{
							Base: Base{
								FRN:         1,
								DataItem:    "I000/010",
								Description: "Test item",
								Type:        _uap.Fixed,
							},
							Data: []byte{0xab, 0xcd},
						},
					},
					{
						FRN: 3,
						Field: &Extended{
							Base: Base{
								FRN:         3,
								DataItem:    "I000/030",
								Description: "Test item",
								Type:        _uap.Extended,
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
			Name: "testCase 2",
			input: RandomFieldSequencing{
				Base:     Base{},
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

func TestRandomFieldSequencingPayload(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  RandomFieldSequencing
		output []byte
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: RandomFieldSequencing{
				Base: Base{
					FRN:         0,
					DataItem:    "I000/000",
					Description: "Test item",
					Type:        _uap.RFS,
				},
				N: 2,
				Sequence: []RandomField{
					{
						FRN: 16,
						Field: &Fixed{
							Base: Base{
								FRN:         1,
								DataItem:    "I000/010",
								Description: "Test item",
								Type:        _uap.Fixed,
							},
							Data: []byte{0xab, 0xcd},
						},
					},
					{
						FRN: 3,
						Field: &Extended{
							Base: Base{
								FRN:         3,
								DataItem:    "I000/030",
								Description: "Test item",
								Type:        _uap.Extended,
							},
							Primary:   []byte{0xc1},
							Secondary: []byte{0xab, 0xcd},
						},
					},
				},
			},
			output: []byte{0x02, 0x10, 0xab, 0xcd, 0x03, 0xc1, 0xab, 0xcd},
		},
		{
			Name: "testCase 2",
			input: RandomFieldSequencing{
				Base:     Base{},
				N:        0,
				Sequence: nil,
			},
			output: []byte{0x00},
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
*/

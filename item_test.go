package goasterix

import (
	"bytes"
	"testing"
)

func TestFixed_Payload(t *testing.T) {
	// Arrange
	fixed := new(Fixed)
	fixed.Data = []byte{0xff, 0xff, 0xff, 0xff}
	output := []byte{0xff, 0xff, 0xff, 0xff}

	// Act
	b := fixed.Payload()

	// Assert
	if len(b) != 4 {
		t.Errorf("FAIL: len(items) = %v; Expected: %v", len(b), 4)
	} else {
		t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(b), 4)
	}
	if bytes.Equal(b, output) == false {
		t.Errorf("FAIL: sp = % X; Expected: % X", b, output)
	} else {
		t.Logf("SUCCESS: sp = % X; Expected: % X", b, output)
	}
}

func TestExtended_Payload(t *testing.T) {
	// Arrange
	ext := new(Extended)
	ext.Primary = []byte{0xff}
	ext.Secondary = []byte{0xff, 0xff, 0xfe}
	output := []byte{0xff, 0xff, 0xff, 0xfe}

	// Act
	b := ext.Payload()

	// Assert
	if len(b) != 4 {
		t.Errorf("FAIL: len(items) = %v; Expected: %v", len(b), 4)
	} else {
		t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(b), 4)
	}
	if bytes.Equal(b, output) == false {
		t.Errorf("FAIL: sp = % X; Expected: % X", b, output)
	} else {
		t.Logf("SUCCESS: sp = % X; Expected: % X", b, output)
	}
}

func TestExplicit_Payload(t *testing.T) {
	// Arrange
	exp := new(Explicit)
	exp.Len = 0x04
	exp.Data = []byte{0xff, 0xff, 0xfe}
	output := []byte{0x04, 0xff, 0xff, 0xfe}

	// Act
	b := exp.Payload()

	// Assert
	if len(b) != 4 {
		t.Errorf("FAIL: len(items) = %v; Expected: %v", len(b), 4)
	} else {
		t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(b), 4)
	}
	if bytes.Equal(b, output) == false {
		t.Errorf("FAIL: sp = % X; Expected: % X", b, output)
	} else {
		t.Logf("SUCCESS: sp = % X; Expected: % X", b, output)
	}
}

func TestRepetitive_Payload(t *testing.T) {
	// Arrange
	rp := new(Repetitive)
	rp.Rep = 0x03
	rp.Data = []byte{0xff, 0xff, 0xfe}

	output := []byte{0x03, 0xff, 0xff, 0xfe}
	// Act
	b := rp.Payload()

	// Assert
	if len(b) != 4 {
		t.Errorf("FAIL: len(items) = %v; Expected: %v", len(b), 4)
	} else {
		t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(b), 4)
	}
	if bytes.Equal(b, output) == false {
		t.Errorf("FAIL: sp = % X; Expected: % X", b, output)
	} else {
		t.Logf("SUCCESS: sp = % X; Expected: % X", b, output)
	}
}

/*func TestCompound_Payload(t *testing.T) {
	// Arrange
	cp := new(Compound)
	cp.Primary = []byte{0xf0}
	cp.Secondary = []Item{
		{
			Meta: MetaItem{
				Type: uap.Fixed,
			},
			Fixed: &Fixed{
				Data: []byte{0xff},
			},
		},
		{
			Meta: MetaItem{
				Type: uap.Extended,
			},
			Extended: &Extended{
				Primary:   []byte{0xff},
				Secondary: []byte{0xff, 0xfe},
			},
		},
		{
			Meta: MetaItem{
				Type: uap.Explicit,
			},
			Explicit: &Explicit{
				Len:  0x02,
				Data: []byte{0xff},
			},
		},
		{
			Meta: MetaItem{
				Type: uap.Repetitive,
			},
			Repetitive: &Repetitive{
				Rep:  0x02,
				Data: []byte{0xff, 0xff},
			},
		},
	}
	output := []byte{0xf0, 0xff, 0xff, 0xff, 0xfe, 0x02, 0xff, 0x02, 0xff, 0xff}

	// Act
	b := cp.Payload()

	// Assert
	if len(b) != 10 {
		t.Errorf("FAIL: len(items) = %v; Expected: %v", len(b), 10)
	} else {
		t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(b), 10)
	}
	if bytes.Equal(b, output) == false {
		t.Errorf("FAIL: cp = % X; Expected: % X", b, output)
	} else {
		t.Logf("SUCCESS: cp = % X; Expected: % X", b, output)
	}
}

func TestItem_Payload(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        Item
		output       []byte
		len          int
	}
	dataSet := []dataTest{
		{
			TestCaseName: "testcase 1",
			input: Item{
				Meta: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Fixed,
				},
				Fixed: &Fixed{
					Data: []byte{0xff, 0xff},
				},
			},
			output: []byte{0xff, 0xff},
			len:    2,
		},
		{
			TestCaseName: "testcase 2",
			input: Item{
				Meta: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Extended,
				},
				Extended: &Extended{
					Primary:   []byte{0xff},
					Secondary: []byte{0xff, 0xfe},
				},
			},
			output: []byte{0xff, 0xff, 0xfe},
			len:    3,
		},
		{
			TestCaseName: "testcase 3",
			input: Item{
				Meta: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Explicit,
				},
				Explicit: &Explicit{
					Len:  0x04,
					Data: []byte{0xff, 0xff, 0xff},
				},
			},
			output: []byte{0x04, 0xff, 0xff, 0xff},
			len:    4,
		},
		{
			TestCaseName: "testcase 4",
			input: Item{
				Meta: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Repetitive,
				},
				Repetitive: &Repetitive{
					Rep:  0x02,
					Data: []byte{0xff, 0xff},
				},
			},
			output: []byte{0x02, 0xff, 0xff},
			len:    3,
		},
		{
			TestCaseName: "testcase 5",
			input: Item{
				Meta: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Compound,
				},
				Compound: &Compound{
					Primary: []byte{0xc0},
					Secondary: []Item{
						{
							Meta: MetaItem{
								FRN:         1,
								DataItem:    "I000/010",
								Description: "Test item",
								Type:        uap.Fixed,
							},
							Fixed: &Fixed{
								Data: []byte{0xff, 0xff},
							},
						},
						{
							Meta: MetaItem{
								FRN:         1,
								DataItem:    "I000/010",
								Description: "Test item",
								Type:        uap.Fixed,
							},
							Fixed: &Fixed{
								Data: []byte{0xff, 0xff},
							},
						},
					},
				},
			},
			output: []byte{0xc0, 0xff, 0xff, 0xff, 0xff},
			len:    5,
		},
	}
	for _, row := range dataSet {
		// Arrange
		// Act
		b := row.input.Payload()

		// Assert
		if len(b) != row.len {
			t.Errorf("FAIL: len(items) = %v; Expected: %v", len(b), row.len)
		} else {
			t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(b), row.len)
		}
		if bytes.Equal(b, row.output) == false {
			t.Errorf("FAIL: item = % X; Expected: % X", b, row.output)
		} else {
			t.Logf("SUCCESS: item = % X; Expected: % X", b, row.output)
		}
	}

}

// Test all item string
func TestItem_String(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Item
		output string
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: Item{
				Meta: MetaItem{
					FRN:         1,
					DataItem:    "I000/010",
					Description: "Test item",
					Type:        uap.Fixed,
				},
				Fixed: &Fixed{
					Data: []byte{0xff, 0xff},
				},
			},
			output: "ffff",
		},
		{
			Name: "testcase 2",
			input: Item{
				Meta: MetaItem{
					FRN:         1,
					DataItem:    "I000/020",
					Description: "Test item",
					Type:        uap.Extended,
				},
				Extended: &Extended{
					Primary:   []byte{0xff},
					Secondary: []byte{0xff, 0xfe},
				},
			},
			output: "fffffe",
		},
		{
			Name: "testcase 3",
			input: Item{
				Meta: MetaItem{
					FRN:         1,
					DataItem:    "I000/030",
					Description: "Test item",
					Type:        uap.Explicit,
				},
				Explicit: &Explicit{
					Len:  0x04,
					Data: []byte{0xff, 0xff, 0xff},
				},
			},
			output: "04ffffff",
		},
		{
			Name: "testcase 4",
			input: Item{
				Meta: MetaItem{
					FRN:         1,
					DataItem:    "I000/040",
					Description: "Test item",
					Type:        uap.Repetitive,
				},
				Repetitive: &Repetitive{
					Rep:  0x02,
					Data: []byte{0xff, 0xff},
				},
			},
			output: "02ffff",
		},
		{
			Name: "testcase 5",
			input: Item{
				Meta: MetaItem{
					FRN:         1,
					DataItem:    "I000/050",
					Description: "Test item",
					Type:        uap.Compound,
				},
				Compound: &Compound{
					Primary: []byte{0xc0},
					Secondary: []Item{
						{
							Meta: MetaItem{
								FRN:         1,
								DataItem:    "I000/010",
								Description: "Test item 010",
								Type:        uap.Fixed,
							},
							Fixed: &Fixed{
								Data: []byte{0xff, 0xff},
							},
						},
						{
							Meta: MetaItem{
								FRN:         1,
								DataItem:    "I000/020",
								Description: "Test item 020",
								Type:        uap.Fixed,
							},
							Fixed: &Fixed{
								Data: []byte{0xff, 0xff},
							},
						},
					},
				},
			},
			output: "[primary: c0][I000/010: ffff][I000/020: ffff]",
		},
	}
	for _, row := range dataSet {
		// Act
		s := row.input.String()

		// Assert
		if s == row.output {
			t.Errorf("FAIL: %s - item = %s; Expected: %s", row.Name, s, row.output)
		} else {
			t.Logf("SUCCESS: %s - item = %s; Expected: %s", row.Name, s, row.output)
		}
	}
}*/

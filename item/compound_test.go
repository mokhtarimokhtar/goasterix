package item

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"io"
	"reflect"
	"testing"
)

func TestCompoundReader(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  string
		item   DataItem
		output DataItem
		err    error
	}

	// Arrange
	dataSet := []testCase{
		{
			Name:  "testCase 1",
			input: "80 01",
			item: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Secondary: []DataItem{
					&Fixed{
						Base: Base{
							FRN:  1,
							Type: FixedField,
						},
						Size: 1,
						SubItems: []SubItem{
							{
								Name: "SUB-A",
								Type: FromToField,
								From: 8,
								To:   1,
							},
						},
					},
				},
			},
			err: nil,
			output: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Primary: []byte{0x80},
				Secondary: []DataItem{
					&Fixed{
						Base: Base{
							FRN:  1,
							Type: FixedField,
						},
						Size: 1,
						SubItems: []SubItem{
							{
								Name: "SUB-A",
								Type: FromToField,
								From: 8,
								To:   1,
								Data: []byte{0x01},
							},
						},
					},
				},
			},
		},
		{
			Name:  "testCase 2",
			input: "40 01fe",
			item: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Secondary: []DataItem{
					&Spare{Base: Base{FRN: 1, Type: SpareField}},
					&Extended{
						Base:              Base{FRN: 2, Type: ExtendedField},
						PrimaryItemSize:   1,
						SecondaryItemSize: 1,
						SubItems: []SubItem{
							{Name: "SUB-A", Type: FromToField, From: 8, To: 2},
							{Name: "FX", Type: BitField, Bit: 1},
							{Name: "SUB-B", Type: BitField, Bit: 8},
							{Name: "SUB-C", Type: FromToField, From: 7, To: 2},
							{Name: "FX", Type: BitField, Bit: 1},
						},
					},
				},
			},
			output: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Primary: []byte{0x40},
				Secondary: []DataItem{
					&Extended{
						Base:              Base{FRN: 2, Type: ExtendedField},
						PrimaryItemSize:   1,
						SecondaryItemSize: 1,
						SubItems: []SubItem{
							{Name: "SUB-A", Type: FromToField, From: 8, To: 2, Data: []byte{0x00}},
							{Name: "SUB-B", Type: BitField, Bit: 8, Data: []byte{0x01}},
							{Name: "SUB-C", Type: FromToField, From: 7, To: 2, Data: []byte{0x3f}},
						},
					},
				},
			},
			err: nil,
		},
		{
			Name:  "testCase 3",
			input: "20 02 8000 0007",
			item: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Secondary: []DataItem{
					&Spare{Base: Base{FRN: 1, Type: SpareField}},
					&Spare{Base: Base{FRN: 2, Type: SpareField}},
					&Repetitive{
						Base:        Base{FRN: 3, Type: RepetitiveField},
						SubItemSize: 2,
						SubItems: []SubItem{
							{Type: BitField, Bit: 16},
							{Type: FromToField, From: 12, To: 1},
						},
					},
				},
			},
			output: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Primary: []byte{0x20},
				Secondary: []DataItem{
					&Repetitive{
						Base:        Base{FRN: 3, Type: RepetitiveField},
						SubItemSize: 2,
						Rep:         0x02,
						SubItems: []SubItem{
							{Type: BitField, Bit: 16, Data: []byte{0x01}},
							{Type: FromToField, From: 12, To: 1, Data: []byte{0x00, 0x00}},

							{Type: BitField, Bit: 16, Data: []byte{0x00}},
							{Type: FromToField, From: 12, To: 1, Data: []byte{0x00, 0x07}},
						},
					},
				},
			},
			err: nil,
		},
		{
			Name:  "testCase 4",
			input: "10 03ffff",
			item: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Secondary: []DataItem{
					&Spare{Base: Base{FRN: 1, Type: SpareField}},
					&Spare{Base: Base{FRN: 2, Type: SpareField}},
					&Spare{Base: Base{FRN: 3, Type: SpareField}},
					&Explicit{
						Base: Base{FRN: 4, Type: ExplicitField},
					},
				},
			},
			output: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Primary: []byte{0x10},
				Secondary: []DataItem{
					&Explicit{
						Base: Base{FRN: 4, Type: ExplicitField},
						Len:  0x03,
						SubItems: []SubItem{
							{Type: FromToField, From: 16, To: 1, Data: []byte{0xff, 0xff}},
						},
					},
				},
			},
			err: nil,
		},
		{
			Name:  "testCase 5",
			input: "",
			item: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Secondary: []DataItem{
					&Fixed{
						Base: Base{
							FRN:  1,
							Type: FixedField,
						},
						Size: 1,
					},
				},
			},
			err: io.EOF,
			output: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Primary:   nil,
				Secondary: nil,
			},
		},
		{
			Name:  "testCase 6",
			input: "80 01",
			item: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Secondary: []DataItem{
					&Extended{
						Base: Base{
							FRN:  1,
							Type: ExtendedField,
						},
						PrimaryItemSize:   1,
						SecondaryItemSize: 1,
					},
				},
			},
			err: io.EOF,
			output: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Primary:   []byte{0x80},
				Secondary: []DataItem{},
			},
		},
		{
			Name:  "testCase 7",
			input: "80 01",
			item: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Secondary: []DataItem{
					&SpecialPurpose{
						Base: Base{
							FRN:  1,
							Type: SPField,
						},
					},
				},
			},
			err: ErrDataFieldUnknown,
			output: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Primary:   []byte{0x80},
				Secondary: []DataItem{},
			},
		},
		{
			Name:  "testCase 8",
			input: "80 02",
			item: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Secondary: []DataItem{
					&Explicit{
						Base: Base{
							FRN:  1,
							Type: ExplicitField,
						},
					},
				},
			},
			err: io.EOF,
			output: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Primary:   []byte{0x80},
				Secondary: []DataItem{},
			},
		},
		{
			Name:  "testCase 9",
			input: "80 02",
			item: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Secondary: []DataItem{
					&Repetitive{
						Base: Base{
							FRN:  1,
							Type: RepetitiveField,
						},
						SubItemSize: 2,
					},
				},
			},
			err: io.EOF,
			output: &Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Primary:   []byte{0x80},
				Secondary: []DataItem{},
			},
		},
	}

	for _, tc := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(tc.input)
		rb := bytes.NewReader(input)
		f := tc.item.Clone()

		// Act
		err := f.Reader(rb)

		// Assert
		if err != tc.err {
			t.Errorf(util.MsgFailInValue, tc.Name, err, tc.err)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, err, tc.err)
		}

		if reflect.DeepEqual(f, tc.output) == false {
			t.Errorf(util.MsgFailInValue, tc.Name, f, tc.output)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, f, tc.output)
		}
	}
}

func TestCompoundGetSubItems(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Compound
		output []SubItem
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testCase 1",
			input: Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Secondary: []DataItem{
					&Fixed{
						Base: Base{
							FRN:  1,
							Type: FixedField,
						},
						Size: 1,
						SubItems: []SubItem{
							{
								Name: "SUB-A",
								Type: FromToField,
								From: 8,
								To:   1,
								Data: []byte{0xab},
							},
						},
					},
				},
			},
			output: []SubItem{
				{
					Name: "SUB-A",
					Type: FromToField,
					From: 8, To: 1,
					Data: []byte{0xab},
				},
			},
		},
		{
			Name: "testCase 2",
			input: Compound{
				Base: Base{
					FRN:  16,
					Type: CompoundField,
				},
				Secondary: []DataItem{
					&Fixed{
						Base: Base{
							FRN:  1,
							Type: FixedField,
						},
						Size: 1,
						SubItems: []SubItem{
							{
								Name: "SUB-A",
								Type: FromToField,
								From: 8,
								To:   1,
								Data: []byte{0xab},
							},
						},
					},
					&Extended{
						Base:              Base{FRN: 2, Type: ExtendedField},
						PrimaryItemSize:   1,
						SecondaryItemSize: 1,
						SubItems: []SubItem{
							{Name: "SUB-B", Type: FromToField, From: 8, To: 2, Data: []byte{0x00}},
							{Name: "SUB-C", Type: BitField, Bit: 8, Data: []byte{0x01}},
							{Name: "SUB-D", Type: FromToField, From: 7, To: 2, Data: []byte{0x3f}},
						},
					},
					&Repetitive{
						Base:        Base{FRN: 3, Type: RepetitiveField},
						SubItemSize: 2,
						Rep:         0x02,
						SubItems: []SubItem{
							{Name: "SUB-E", Type: BitField, Bit: 16, Data: []byte{0x01}},
							{Name: "SUB-F", Type: FromToField, From: 12, To: 1, Data: []byte{0x00, 0x00}},
							{Name: "SUB-G", Type: BitField, Bit: 16, Data: []byte{0x00}},
							{Name: "SUB-H", Type: FromToField, From: 12, To: 1, Data: []byte{0x00, 0x07}},
						},
					},
				},
			},
			output: []SubItem{
				{Name: "SUB-A", Type: FromToField, From: 8, To: 1, Data: []byte{0xab}},
				{Name: "SUB-B", Type: FromToField, From: 8, To: 2, Data: []byte{0x00}},
				{Name: "SUB-C", Type: BitField, Bit: 8, Data: []byte{0x01}},
				{Name: "SUB-D", Type: FromToField, From: 7, To: 2, Data: []byte{0x3f}},
				{Name: "SUB-E", Type: BitField, Bit: 16, Data: []byte{0x01}},
				{Name: "SUB-F", Type: FromToField, From: 12, To: 1, Data: []byte{0x00, 0x00}},
				{Name: "SUB-G", Type: BitField, Bit: 16, Data: []byte{0x00}},
				{Name: "SUB-H", Type: FromToField, From: 12, To: 1, Data: []byte{0x00, 0x07}},
			},
		},
		{
			Name: "testCase 3",
			input: Compound{
				Base: Base{},
			},
			output: nil,
		},
	}

	for _, tc := range dataSet {
		// Act
		s := tc.input.GetSubItems()

		// Assert
		if reflect.DeepEqual(s, tc.output) == false {
			t.Errorf(util.MsgFailInValue, tc.Name, s, tc.output)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, s, tc.output)
		}
	}
}

/*
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
			Name: "testCase 1",
			input: Compound{
				Base: Base{
					FieldReferenceNumber:         1,
					DataItemName:    "I000/000",
					Description: "Test item",
					Type:        _uap.Compound,
				},
				Primary: []byte{0xaa},
				Secondary: []DataItemName{
					&Fixed{
						Base: Base{
							FieldReferenceNumber:         1,
							DataItemName:    "I000/010",
							Description: "Test item",
							Type:        _uap.Fixed,
						},
						Data: []byte{0xab, 0xcd},
					},
					&Extended{
						Base: Base{
							FieldReferenceNumber:         3,
							DataItemName:    "I000/030",
							Description: "Test item",
							Type:        _uap.Extended,
						},
						Primary:   []byte{0xc1},
						Secondary: []byte{0xab, 0xcd},
					},
					&Explicit{
						Base: Base{
							FieldReferenceNumber:         5,
							DataItemName:    "I000/050",
							Description: "Test item",
							Type:        _uap.Explicit,
						},
						Len:  0x04,
						Data: []byte{0xab, 0xcd, 0xef},
					},
					&Repetitive{
						Base: Base{
							FieldReferenceNumber:         7,
							DataItemName:    "I000/070",
							Description: "Test item",
							Type:        _uap.Repetitive,
						},
						Rep:  0x02,
						Data: []byte{0xab, 0xcd},
					},
				},
			},
			output: "I000/000:[primary:aa][I000/010:abcd][I000/030:c1abcd][I000/050:04abcdef][I000/070:02abcd]",
		},
		{
			Name: "testCase 2",
			input: Compound{
				Base:      Base{},
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
			Name: "testCase 1",
			input: Compound{
				Base: Base{
					FieldReferenceNumber:         1,
					DataItemName:    "I000/000",
					Description: "Test item",
					Type:        _uap.Compound,
				},
				Primary: []byte{0xaa},
				Secondary: []DataItemName{
					&Fixed{
						Base: Base{
							FieldReferenceNumber:         1,
							DataItemName:    "I000/010",
							Description: "Test item",
							Type:        _uap.Fixed,
						},
						Data: []byte{0xab, 0xcd},
					},
					&Extended{
						Base: Base{
							FieldReferenceNumber:         3,
							DataItemName:    "I000/030",
							Description: "Test item",
							Type:        _uap.Extended,
						},
						Primary:   []byte{0xc1},
						Secondary: []byte{0xab, 0xcd},
					},
					&Explicit{
						Base: Base{
							FieldReferenceNumber:         5,
							DataItemName:    "I000/050",
							Description: "Test item",
							Type:        _uap.Explicit,
						},
						Len:  0x04,
						Data: []byte{0xab, 0xcd, 0xef},
					},
					&Repetitive{
						Base: Base{
							FieldReferenceNumber:         7,
							DataItemName:    "I000/070",
							Description: "Test item",
							Type:        _uap.Repetitive,
						},
						Rep:  0x02,
						Data: []byte{0xab, 0xcd},
					},
				},
			},
			output: []byte{0xaa, 0xab, 0xcd, 0xc1, 0xab, 0xcd, 0x04, 0xab, 0xcd, 0xef, 0x02, 0xab, 0xcd},
		},
		{
			Name: "testCase 2",
			input: Compound{
				Base:      Base{},
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
*/

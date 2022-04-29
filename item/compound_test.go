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
						Data: []byte{0x01},
					},
				},
			},
		},
		{
			Name:  "testCase 2",
			input: "d4 01 01fe 0201020102 030102",
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
					&Extended{
						Base: Base{
							FRN:  2,
							Type: ExtendedField,
						},
						PrimaryItemSize:   1,
						SecondaryItemSize: 1,
					},
					&Fixed{
						Base: Base{
							FRN:  3,
							Type: FixedField,
						},
						Size: 1,
					},
					&Repetitive{
						Base: Base{
							FRN:  4,
							Type: RepetitiveField,
						},
						SubItemSize: 2,
					},
					&Spare{Base{
						FRN: 5,
					}},
					&Explicit{
						Base: Base{
							FRN:  6,
							Type: ExplicitField,
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
				Primary: []byte{0xd4},
				Secondary: []DataItem{
					&Fixed{
						Base: Base{
							FRN:  1,
							Type: FixedField,
						},
						Size: 1,
						Data: []byte{0x01},
					},
					&Extended{
						Base: Base{
							FRN:  2,
							Type: ExtendedField,
						},
						PrimaryItemSize:   1,
						SecondaryItemSize: 1,
						Primary:           []byte{0x01},
						Secondary:         []byte{0xfe},
					},
					&Repetitive{
						Base: Base{
							FRN:  4,
							Type: RepetitiveField,
						},
						SubItemSize: 2,
						Rep:         0x02,
						Data:        []byte{0x01, 0x02, 0x01, 0x02},
					},
					&Explicit{
						Base: Base{
							FRN:  6,
							Type: ExplicitField,
						},
						Len:  0x03,
						Data: []byte{0x01, 0x02},
					},
				},
			},
		},
		{
			Name:  "testCase 3",
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
			Name:  "testCase 5",
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
			Name:  "testCase 6",
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
			Name:  "testCase 7",
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
			Name:  "testCase 8",
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

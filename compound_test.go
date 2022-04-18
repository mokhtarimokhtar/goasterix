package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
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
		uap    uap.DataField
		output Compound
		err    error
	}
	// Arrange
	dataSet := []testCase{
		{
			Name:  "testcase 1",
			input: "80 01",
			uap: uap.DataField{
				FRN:  16,
				Type: uap.Compound,
				Compound: []uap.DataField{
					{
						FRN:   1,
						Type:  uap.Fixed,
						Fixed: uap.FixedField{Size: 1},
					},
				},
			},
			err: nil,
			output: Compound{
				MetaItem: MetaItem{
					FRN:  16,
					Type: uap.Compound,
				},
				Primary: []byte{0x80},
				Secondary: []Item{
					&Fixed{
						MetaItem: MetaItem{
							FRN:  1,
							Type: uap.Fixed,
						},
						Data: []byte{0x01},
					},
				},
			},
		},
		{
			Name:  "testcase 2",
			input: "d4 01 01fe 0201020102 030102",
			uap: uap.DataField{
				Type: uap.Compound,
				Compound: []uap.DataField{
					{
						FRN:   1,
						Type:  uap.Fixed,
						Fixed: uap.FixedField{Size: 1},
					},
					{
						FRN:  2,
						Type: uap.Extended,
						Extended: uap.ExtendedField{
							PrimarySize:   1,
							SecondarySize: 1,
						},
					},
					{
						FRN:   3,
						Type:  uap.Fixed,
						Fixed: uap.FixedField{Size: 1},
					},
					{
						FRN:        4,
						Type:       uap.Repetitive,
						Repetitive: uap.RepetitiveField{SubItemSize: 2},
					},
					{
						FRN:  5,
						Type: uap.Spare,
					},
					{
						FRN:  6,
						Type: uap.Explicit,
					},
				},
			},
			err: nil,
			output: Compound{
				MetaItem: MetaItem{
					Type: uap.Compound,
				},
				Primary: []byte{0xd4},
				Secondary: []Item{
					&Fixed{
						MetaItem: MetaItem{
							FRN:  1,
							Type: uap.Fixed,
						},
						Data: []byte{0x01},
					},
					&Extended{
						MetaItem: MetaItem{
							FRN:  2,
							Type: uap.Extended,
						},
						Primary:   []byte{0x01},
						Secondary: []byte{0xfe},
					},
					&Repetitive{
						MetaItem: MetaItem{
							FRN:  4,
							Type: uap.Repetitive,
						},
						Rep:  0x02,
						Data: []byte{0x01, 0x02, 0x01, 0x02},
					},
					&Explicit{
						MetaItem: MetaItem{
							FRN:  6,
							Type: uap.Explicit,
						},
						Len:  0x03,
						Data: []byte{0x01, 0x02},
					},
				},
			},
		},
		{
			Name:  "testcase 3",
			input: "",
			uap: uap.DataField{
				Type: uap.Compound,
				Compound: []uap.DataField{
					{
						FRN:   1,
						Type:  uap.Fixed,
						Fixed: uap.FixedField{Size: 1},
					},
				},
			},
			err: io.EOF,
			output: Compound{
				MetaItem: MetaItem{
					Type: uap.Compound,
				},
				Primary:   nil,
				Secondary: nil,
			},
		},
		{
			Name:  "testcase 4",
			input: "80 01",
			uap: uap.DataField{
				FRN:  16,
				Type: uap.Compound,
				Compound: []uap.DataField{
					{
						FRN:  1,
						Type: uap.Extended,
						Extended: uap.ExtendedField{
							PrimarySize:   1,
							SecondarySize: 1,
						},
					},
				},
			},
			err: io.EOF,
			output: Compound{
				MetaItem: MetaItem{
					FRN:  16,
					Type: uap.Compound,
				},
				Primary:   []byte{0x80},
				Secondary: nil,
			},
		},
		{
			Name:  "testcase 5",
			input: "80 01",
			uap: uap.DataField{
				FRN:  16,
				Type: uap.Compound,
				Compound: []uap.DataField{
					{
						FRN:  1,
						Type: uap.SP,
					},
				},
			},
			err: ErrDataFieldUnknown,
			output: Compound{
				MetaItem: MetaItem{
					FRN:  16,
					Type: uap.Compound,
				},
				Primary:   []byte{0x80},
				Secondary: nil,
			},
		},
		{
			Name:  "testcase 6",
			input: "80 02",
			uap: uap.DataField{
				FRN:  16,
				Type: uap.Compound,
				Compound: []uap.DataField{
					{
						FRN:  1,
						Type: uap.Explicit,
					},
				},
			},
			err: io.EOF,
			output: Compound{
				MetaItem: MetaItem{
					FRN:  16,
					Type: uap.Compound,
				},
				Primary:   []byte{0x80},
				Secondary: nil,
			},
		},
		{
			Name:  "testcase 7",
			input: "80 020102",
			uap: uap.DataField{
				FRN:  16,
				Type: uap.Compound,
				Compound: []uap.DataField{
					{
						FRN:        1,
						Type:       uap.Repetitive,
						Repetitive: uap.RepetitiveField{SubItemSize: 2},
					},
				},
			},
			err: io.ErrUnexpectedEOF,
			output: Compound{
				MetaItem: MetaItem{
					FRN:  16,
					Type: uap.Compound,
				},
				Primary:   []byte{0x80},
				Secondary: nil,
			},
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(row.input)
		rb := bytes.NewReader(input)
		f := Compound{}

		// Act
		err := f.Reader(rb, row.uap)
		
		// Assert
		if err != row.err {
			t.Errorf(util.MsgFailInValue, row.Name, err, row.err)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, err, row.err)
		}

		if reflect.DeepEqual(f, row.output) == false {
			t.Errorf(util.MsgFailInValue, row.Name, f, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, f, row.output)
		}
	}
}

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

func TestCompoundFrn(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  Compound
		output uint8
	}
	// Arrange
	dataSet := []testCase{
		{
			Name: "testcase 1",
			input: Compound{
				MetaItem: MetaItem{
					FRN:         7,
					DataItem:    "I000/070",
					Description: "Test item",
					Type:        uap.Compound,
				},
				Primary:   nil,
				Secondary: nil,
			},
			output: 7,
		},
		{
			Name: "testcase 2",
			input: Compound{
				MetaItem:  MetaItem{},
				Primary:   nil,
				Secondary: nil,
			},
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

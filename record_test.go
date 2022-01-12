package goasterix

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/mokhtarimokhtar/goasterix/uap"
)

func TestRecord_Payload(t *testing.T) {
	// Arrange
	data, _ := HexStringToByte("ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020a0")
	nbOfBytes := 42
	rec := NewRecord()
	_, _ = rec.Decode(data, uap.Cat048V127)

	// Act
	items := rec.Payload()

	// Assert
	if len(items) != nbOfBytes {
		t.Errorf("FAIL: len(items) = %v; Expected: %v", len(items), nbOfBytes)
	} else {
		t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(items), nbOfBytes)
	}
}

func TestRecord_String(t *testing.T) {
	// Arrange
	data, _ := HexStringToByte("ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020a0")
	nbOfItems := 15
	rec := NewRecord()
	_, _ = rec.Decode(data, uap.Cat048V127)

	// Act
	items := rec.String()

	// Assert
	if len(items) != nbOfItems {
		t.Errorf("FAIL: len(items) = %v; Expected: %v", len(items), nbOfItems)
	} else {
		t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(items), nbOfItems)
	}
}

func TestFspecReader_Valid(t *testing.T) {
	// Arrange
	input := []byte{0xFF, 0x01, 0xF2, 0xFF}
	output := []byte{0xFF, 0x01, 0xF2}
	rb := bytes.NewReader(input)

	// Act
	fspec, err := FspecReader(rb)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if bytes.Equal(fspec, output) == false {
		t.Errorf("FAIL: sp = % X; Expected: % X", fspec, output)
	} else {
		t.Logf("SUCCESS: sp = % X; Expected: % X", fspec, output)
	}
}

func TestFspecReader_Invalid(t *testing.T) {
	// Arrange
	input := []byte{0xFF, 0x01}
	var output []byte
	rb := bytes.NewReader(input)

	// Act
	fspec, err := FspecReader(rb)

	// Assert
	if err != io.EOF {
		t.Errorf("FAIL: error: %s; Expected: %v", err, io.EOF)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, io.EOF)
	}

	if bytes.Equal(fspec, output) == false {
		t.Errorf("FAIL: sp = % X; Expected: % X", fspec, output)
	} else {
		t.Logf("SUCCESS: sp = % X; Expected: % X", fspec, output)
	}
}

func TestFspecIndex(t *testing.T) {
	type frnIndexTest struct {
		input  []byte
		output []uint8
	}
	// Arrange
	dataSet := []frnIndexTest{
		{input: []byte{0x80}, output: []uint8{1}},
		{input: []byte{0x40}, output: []uint8{2}},
		{input: []byte{0x20}, output: []uint8{3}},
		{input: []byte{0x10}, output: []uint8{4}},
		{input: []byte{0x08}, output: []uint8{5}},
		{input: []byte{0x04}, output: []uint8{6}},
		{input: []byte{0x02}, output: []uint8{7}},
		{input: []byte{0x01}, output: []uint8{}},
		{input: []byte{0x01, 0x80}, output: []uint8{8}},
		{input: []byte{0xfe}, output: []uint8{1, 2, 3, 4, 5, 6, 7}},
		{input: []byte{0xff}, output: []uint8{1, 2, 3, 4, 5, 6, 7}},
		{input: []byte{0xaa}, output: []uint8{1, 3, 5, 7}},
		{input: []byte{0x55}, output: []uint8{2, 4, 6}},
		{input: []byte{}, output: []uint8{}},
	}

	for _, row := range dataSet {
		// Act
		frnIndex := FspecIndex(row.input)

		// Assert
		if bytes.Equal(frnIndex, row.output) == false {
			t.Errorf("FAIL: % X; Expected: % X", frnIndex, row.output)
		} else {
			t.Logf("SUCCESS: % X; Expected: % X", frnIndex, row.output)
		}
	}

}

// FixedDataField
func TestFixedDataFieldReader_Valid(t *testing.T) {
	// Arrange
	input, _ := HexStringToByte("FF FE FD BF 00 01 02 03")
	nb := uint8(8)
	rb := bytes.NewReader(input)
	output := []byte{0xFF, 0xFE, 0xFD, 0xBF, 0x00, 0x01, 0x02, 0x03}

	// Act
	item, err := FixedDataFieldReader(rb, nb)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if bytes.Equal(item.Data, output) == false {
		t.Errorf("FAIL: sp = % X; Expected: % X", item, output)
	} else {
		t.Logf("SUCCESS: sp = % X; Expected: % X", item, output)
	}
}

func TestFixedDataFieldReader_Invalid(t *testing.T) {
	// Arrange
	input, _ := HexStringToByte("FF FE BF 00 01 02")
	nb := uint8(7)
	rb := bytes.NewReader(input)

	// Act
	item, err := FixedDataFieldReader(rb, nb)

	// Assert
	if err != io.ErrUnexpectedEOF {
		t.Errorf("FAIL: error: %v; Expected: %v", err, io.ErrUnexpectedEOF)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, io.ErrUnexpectedEOF)
	}
	if item.Data != nil {
		t.Errorf("FAIL: item = %v; Expected: %v", item, nil)
	} else {
		t.Logf("SUCCESS: item = %v; Expected: %v", item, nil)
	}
}

// ExtendedDataField
func TestExtendedDataFieldReader(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName  string
		input         string
		primarySize   uint8
		secondarySize uint8
		output        Extended
		err           error
	}
	dataSet := []dataTest{
		{
			TestCaseName:  "testcase 1",
			input:         "01 03 07 09 0B 0D 0F 0E",
			primarySize:   1,
			secondarySize: 1,
			output: Extended{
				Primary:   []byte{0x01},
				Secondary: []byte{0x03, 0x07, 0x09, 0x0B, 0x0D, 0x0F, 0x0E},
			},
			err: nil,
		},
		{
			TestCaseName:  "testcase 2",
			input:         "FE",
			primarySize:   1,
			secondarySize: 1,
			output: Extended{
				Primary: []byte{0xFE},
			},
			err: nil,
		},
		{
			TestCaseName:  "testcase 3",
			input:         "",
			primarySize:   1,
			secondarySize: 1,
			output: Extended{
				Primary: nil,
			},
			err: io.EOF,
		},
		{
			TestCaseName:  "testcase 4",
			input:         "FF",
			primarySize:   1,
			secondarySize: 1,
			output: Extended{
				Primary: []byte{0xff},
			},
			err: io.EOF,
		},
		{
			TestCaseName:  "testcase 5",
			input:         "FF",
			primarySize:   2,
			secondarySize: 1,
			output: Extended{
				Primary: nil,
			},
			err: io.ErrUnexpectedEOF,
		},
		{
			TestCaseName:  "testcase 6",
			input:         "0001 000001 FFFFFE",
			primarySize:   2,
			secondarySize: 3,
			output: Extended{
				Primary:   []byte{0x00, 0x01},
				Secondary: []byte{0x00, 0x00, 0x01, 0xFF, 0xFF, 0xFE},
			},
			err: nil,
		},
		{
			TestCaseName:  "testcase 7",
			input:         "0001 000001 FFFF",
			primarySize:   2,
			secondarySize: 3,
			output: Extended{
				Primary:   []byte{0x00, 0x01},
				Secondary: []byte{0x00, 0x00, 0x01},
			},
			err: io.ErrUnexpectedEOF,
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := HexStringToByte(row.input)
		rb := bytes.NewReader(input)

		// Act
		item, err := ExtendedDataFieldReader(rb, row.primarySize, row.secondarySize)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: %s - error: %v; Expected: %v", row.TestCaseName, err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}
		if reflect.DeepEqual(item, row.output) == false {
			t.Errorf("FAIL: %s - item = % X; Expected: % X", row.TestCaseName, item, row.output)
		} else {
			t.Logf("SUCCESS: item = % X; Expected: % X", item, row.output)
		}
	}
}

// ExplicitDataField
func TestExplicitDataFieldReader(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        string
		output       Explicit
		err          error
	}
	dataSet := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        "03 FF FF",
			output: Explicit{
				Len:  0x03,
				Data: []byte{0xFF, 0xFF},
			},
			err: nil,
		},
		{
			TestCaseName: "testcase 2",
			input:        "03 FF",
			output: Explicit{
				Len:  0x03,
				Data: nil,
			},
			err: io.ErrUnexpectedEOF,
		},
		{
			TestCaseName: "testcase 3",
			input:        "",
			output: Explicit{
				Len:  0,
				Data: nil,
			},
			err: io.EOF,
		},
	}
	for _, row := range dataSet {
		// Arrange
		input, _ := HexStringToByte(row.input)
		rb := bytes.NewReader(input)

		// Act
		item, err := ExplicitDataFieldReader(rb)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: %s - error: %v; Expected: %v", row.TestCaseName, err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}
		if reflect.DeepEqual(item, row.output) == false {
			t.Errorf("FAIL: %s - item = % X; Expected: % X", row.TestCaseName, item, row.output)
		} else {
			t.Logf("SUCCESS: item = % X; Expected: % X", item, row.output)
		}
	}

}

// RepetitiveDataField
func TestRepetitiveDataFieldReader(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        string
		SubItemSize  uint8
		output       Repetitive
		err          error
	}
	dataSet := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        "03 01 02 03 04 05 06 07 08 09",
			SubItemSize:  3,
			output: Repetitive{
				Rep:  0x03,
				Data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
			},
			err: nil,
		},
		{
			TestCaseName: "testcase 2",
			input:        "04 01 02 03 04 05 06 07 08 09",
			SubItemSize:  3,
			output: Repetitive{
				Rep:  0x04,
				Data: nil,
			},
			err: io.ErrUnexpectedEOF,
		},
		{
			TestCaseName: "testcase 3",
			input:        "",
			SubItemSize:  3,
			output: Repetitive{
				Rep:  0,
				Data: nil,
			},
			err: io.EOF,
		},
	}
	for _, row := range dataSet {
		// Arrange
		input, _ := HexStringToByte(row.input)
		rb := bytes.NewReader(input)

		// Act
		item, err := RepetitiveDataFieldReader(rb, row.SubItemSize)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: %s - error: %v; Expected: %v", row.TestCaseName, err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}
		if reflect.DeepEqual(item, row.output) == false {
			t.Errorf("FAIL: %s - item = % X; Expected: % X", row.TestCaseName, item, row.output)
		} else {
			t.Logf("SUCCESS: item = % X; Expected: % X", item, row.output)
		}
	}
}

// CompoundDataField
func TestCompoundDataFieldReader(t *testing.T) {
	// Setup
	type dataTest struct {
		TestCaseName string
		input        string
		output       Compound
		item         []uap.DataField
		err          error
	}
	dataSet := []dataTest{
		{
			TestCaseName: "Compound type: two primaries subitems and follow valid subitems",
			input: "FF FE " +
				"FFFFFF  FFFFFFFFFFFF FFFF FFFF FFFF FFFF FFFF" +
				"FF " +
				"02 FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF " +
				"FFFF FFFF FFFFFFFFFFFFFF FFFF FFFF",
			output: Compound{
				Primary: []byte{0xFF, 0xFE},
				Secondary: []Item{
					{
						Meta:  MetaItem{FRN: 1, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF, 0xFF}},
					},
					{
						Meta:  MetaItem{FRN: 2, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
					},
					{
						Meta:  MetaItem{FRN: 3, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF}},
					},
					{
						Meta:  MetaItem{FRN: 4, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF}},
					},
					{
						Meta:  MetaItem{FRN: 5, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF}},
					},
					{
						Meta:  MetaItem{FRN: 6, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF}},
					},
					{
						Meta:  MetaItem{FRN: 7, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF}},
					},
					{
						Meta:  MetaItem{FRN: 8, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF}},
					},
					{
						Meta: MetaItem{FRN: 9, Type: uap.Repetitive},
						Repetitive: &Repetitive{
							Rep:  0x02,
							Data: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
						},
					},
					{
						Meta:  MetaItem{FRN: 10, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF}},
					},
					{
						Meta:  MetaItem{FRN: 11, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF}},
					},
					{
						Meta:  MetaItem{FRN: 12, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
					},
					{
						Meta:  MetaItem{FRN: 13, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF}},
					},
					{
						Meta:  MetaItem{FRN: 14, Type: uap.Fixed},
						Fixed: &Fixed{Data: []byte{0xFF, 0xFF}},
					},
				},
			},
			item: []uap.DataField{
				{FRN: 1, Type: uap.Fixed, Fixed: uap.FixedField{Size: 3}},
				{FRN: 2, Type: uap.Fixed, Fixed: uap.FixedField{Size: 6}},
				{FRN: 3, Type: uap.Fixed, Fixed: uap.FixedField{Size: 2}},
				{FRN: 4, Type: uap.Fixed, Fixed: uap.FixedField{Size: 2}},
				{FRN: 5, Type: uap.Fixed, Fixed: uap.FixedField{Size: 2}},
				{FRN: 6, Type: uap.Fixed, Fixed: uap.FixedField{Size: 2}},
				{FRN: 7, Type: uap.Fixed, Fixed: uap.FixedField{Size: 2}},

				{FRN: 8, Type: uap.Fixed, Fixed: uap.FixedField{Size: 1}},
				{FRN: 9, Type: uap.Repetitive, Repetitive: uap.RepetitiveField{SubItemSize: 15}},
				{FRN: 10, Type: uap.Fixed, Fixed: uap.FixedField{Size: 2}},
				{FRN: 11, Type: uap.Fixed, Fixed: uap.FixedField{Size: 2}},
				{FRN: 12, Type: uap.Fixed, Fixed: uap.FixedField{Size: 7}},
				{FRN: 13, Type: uap.Fixed, Fixed: uap.FixedField{Size: 2}},
				{FRN: 14, Type: uap.Fixed, Fixed: uap.FixedField{Size: 2}},
			},
			err: nil,
		},
		{
			TestCaseName: "Compound type: empty",
			input:        "",
			output:       Compound{},
			item:         []uap.DataField{},
			err:          io.EOF,
		},
		{
			TestCaseName: "Compound type: ErrDataFieldUnknown",
			input:        "40 ff",
			output: Compound{
				Primary: []byte{0x40},
			},
			item: []uap.DataField{
				{FRN: 1, Type: uap.Fixed, Fixed: uap.FixedField{Size: 3}},
				{FRN: 2, Type: uap.Spare, Fixed: uap.FixedField{Size: 2}},
			},
			err: ErrDataFieldUnknown,
		},
		{
			TestCaseName: "Compound type: extended error",
			input:        "80 ff",
			output: Compound{
				Primary: []byte{0x80},
			},
			item: []uap.DataField{
				{FRN: 1, Type: uap.Extended, Extended: uap.ExtendedField{PrimarySize: 1, SecondarySize: 1}},
			},
			err: io.EOF,
		},
		{
			TestCaseName: "Compound type: repetitive error",
			input:        "80 02ffff",
			output: Compound{
				Primary: []byte{0x80},
			},
			item: []uap.DataField{
				{FRN: 1, Type: uap.Repetitive, Repetitive: uap.RepetitiveField{SubItemSize: 2}},
			},
			err: io.ErrUnexpectedEOF,
		},
		{
			TestCaseName: "Compound type: explicit error",
			input:        "80 03ff",
			output: Compound{
				Primary: []byte{0x80},
			},
			item: []uap.DataField{
				{FRN: 1, Type: uap.Explicit, Explicit: uap.ExplicitField{}},
			},
			err: io.ErrUnexpectedEOF,
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := HexStringToByte(row.input)
		rb := bytes.NewReader(input)

		// Act
		cp, err := CompoundDataFieldReader(rb, row.item)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: error: %v; Expected: %v", err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}
		if reflect.DeepEqual(cp, row.output) == false {
			t.Errorf("FAIL: %s - \nCompound = %v;\nExpected: %v", row.TestCaseName, cp, row.output)
		} else {
			t.Logf("SUCCESS: Compound = %v; Expected: %v", cp, row.output)
		}
	}
}

// SPAndREDataField
func TestSPAndREDataFieldReader(t *testing.T) {
	// Setup
	type dataTest struct {
		TestCaseName string
		input        string
		output       SpecialPurpose
		err          error
	}
	dataSet := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        "03 FF FF",
			output: SpecialPurpose{
				Len:  0x03,
				Data: []byte{0xFF, 0xFF},
			},
			err: nil,
		},
		{
			TestCaseName: "testcase 2",
			input:        "03 FF",
			output: SpecialPurpose{
				Len:  0x03,
				Data: nil,
			},
			err: io.ErrUnexpectedEOF,
		},
		{
			TestCaseName: "testcase 3",
			input:        "",
			output: SpecialPurpose{
				Len:  0x00,
				Data: nil,
			},
			err: io.EOF,
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := HexStringToByte(row.input)
		rb := bytes.NewReader(input)

		// Act
		cp, err := SPAndREDataFieldReader(rb)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: error: %v; Expected: %v", err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}
		if reflect.DeepEqual(cp, row.output) == false {
			t.Errorf("FAIL: %s - \nCompound = %v;\nExpected: %v", row.TestCaseName, cp, row.output)
		} else {
			t.Logf("SUCCESS: Compound = %v; Expected: %v", cp, row.output)
		}
	}
}

// RFSDataField
func TestRFSDataFieldReader(t *testing.T) {
	// Setup
	type dataTest struct {
		TestCaseName string
		input        string
		item         []uap.DataField
		output       RandomFieldSequencing
		err          error
	}
	dataSet := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        "02 03 ffffffff 0a ff",
			item:         uap.Cat001PlotV12,
			output: RandomFieldSequencing{
				N: 0x02,
				Sequence: []RandomField{
					{
						FRN: 0x03,
						Field: Item{
							Meta: MetaItem{
								FRN:         3,
								DataItem:    "I001/040",
								Description: "Measured Position in Polar Coordinates",
								Type:        uap.Fixed,
							},
							Fixed: &Fixed{Data: []byte{0xff, 0xff, 0xff, 0xff}},
						},
					},
					{
						FRN: 0x0a,
						Field: Item{
							Meta: MetaItem{
								FRN:         10,
								DataItem:    "I001/131",
								Description: "Received Power",
								Type:        uap.Fixed,
							},
							Fixed: &Fixed{Data: []byte{0xff}},
						},
					},
				},
			},
			err: nil,
		},
		{
			TestCaseName: "testcase 2",
			input:        "02",
			item:         uap.Cat001PlotV12,
			output: RandomFieldSequencing{
				N: 0x02,
			},
			err: io.EOF,
		},
		{
			TestCaseName: "testcase 3",
			input:        "02 03 ffffffff 0a",
			item:         uap.Cat001PlotV12,
			output: RandomFieldSequencing{
				N: 0x02,
				Sequence: []RandomField{
					{
						FRN: 0x03,
						Field: Item{
							Meta: MetaItem{
								FRN:         3,
								DataItem:    "I001/040",
								Description: "Measured Position in Polar Coordinates",
								Type:        uap.Fixed,
							},
							Fixed: &Fixed{Data: []byte{0xff, 0xff, 0xff, 0xff}},
						},
					},
				},
			},
			err: io.EOF,
		},
		{
			TestCaseName: "testcase 4",
			input:        "",
			item:         uap.Cat001PlotV12,
			output:       RandomFieldSequencing{},
			err:          io.EOF,
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := HexStringToByte(row.input)
		rb := bytes.NewReader(input)

		// Act
		cp, err := RFSDataFieldReader(rb, row.item)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: error: %v; Expected: %v", err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}
		if reflect.DeepEqual(cp, row.output) == false {
			t.Errorf("FAIL: %s - \nCompound = %v;\nExpected: %v", row.TestCaseName, cp, row.output)
		} else {
			t.Logf("SUCCESS: Compound = %v; Expected: %v", cp, row.output)
		}
	}
}

//Testing by record
func TestRecordDecode_NbOfItems(t *testing.T) {
	// setup
	type dataTest struct {
		input     string          // data test one record = fspec + items
		uap       uap.StandardUAP // Items of category corresponding to data test input
		nbOfItems int
		err       error // error expected
	}
	dataSet := []dataTest{
		{
			input:     "f6083602429b7110940028200094008000",
			uap:       uap.Cat034V127,
			err:       nil,
			nbOfItems: 6,
		},
		{
			input:     "f6083602429b71109400282000940080",
			uap:       uap.Cat034V127,
			err:       io.EOF,
			nbOfItems: 5,
		},
		{
			input:     "ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020a0",
			uap:       uap.Cat048V127,
			err:       nil,
			nbOfItems: 14,
		},
		{
			// 0xA0 last byte is removed
			input:     "ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020",
			uap:       uap.Cat048V127,
			err:       io.ErrUnexpectedEOF,
			nbOfItems: 13,
		},
		{
			input:     "f0 0831 00 0a8abb2e 3802",
			uap:       uap.Cat001V12,
			err:       nil,
			nbOfItems: 4,
		},
		{
			input:     "f0 0831 00 0a8abb2e 38",
			uap:       uap.Cat001V12,
			err:       io.ErrUnexpectedEOF,
			nbOfItems: 3,
		},
		{
			input:     "f502 0831 98 01bf 0a1ebb43 022538e2 00",
			uap:       uap.Cat001V12,
			err:       nil,
			nbOfItems: 6,
		},
		{
			input:     "f502 0831 98 01bf 0a1ebb43 022538e2",
			uap:       uap.Cat001V12,
			err:       io.EOF,
			nbOfItems: 5,
		},
	}

	for _, row := range dataSet {
		// Arrange
		data, _ := HexStringToByte(row.input)
		rec := new(Record)

		// Act
		unRead, err := rec.Decode(data, row.uap)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: error: %s; Expected: %v", err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}
		if unRead != 0 {
			t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
		} else {
			t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
		}
		if row.nbOfItems != len(rec.Items) {
			t.Errorf("FAIL: nbOfItems = %v; Expected: %v", row.nbOfItems, len(rec.Items))
		} else {
			t.Logf("SUCCESS: nbOfItems = %v; Expected: %v", row.nbOfItems, len(rec.Items))
		}
	}
}

func TestRecordDecode_Empty(t *testing.T) {
	// Arrange
	input := ""
	var output []Item
	uap048 := uap.Cat048V127
	data, _ := HexStringToByte(input)
	rec := NewRecord()

	// Act
	unRead, err := rec.Decode(data, uap048)

	// Assert
	if err != io.EOF {
		t.Errorf("FAIL: error = %v; Expected: %v", err, io.EOF)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, io.EOF)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	if reflect.DeepEqual(rec.Items, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", rec.Items, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", rec.Items, output)
	}
}

// Decode Cat4Test
func TestRecordDecode_Cat4TestFullRecord(t *testing.T) {
	// Arrange
	input := "fd 40 ffff fffffe 03ffff 02ffffffff ab80 ff fffe 02ffffffff 04ffffff ffff 0101ffff 03ffff"
	output := []Item{
		{
			Meta: MetaItem{
				FRN:         1,
				DataItem:    "I026/001",
				Description: "Fixed type field for test",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff, 0xff}},
		},
		{
			Meta: MetaItem{
				FRN:         2,
				DataItem:    "I026/002",
				Description: "Extended type field for test",
				Type:        uap.Extended,
			},
			Extended: &Extended{
				Primary:   []byte{0xff},
				Secondary: []byte{0xff, 0xfe},
			},
		},
		{
			Meta: MetaItem{
				FRN:         3,
				DataItem:    "I026/003",
				Description: "Explicit type field for test",
				Type:        uap.Explicit,
			},
			Explicit: &Explicit{
				Len:  0x03,
				Data: []byte{0xff, 0xff},
			},
		},
		{
			Meta: MetaItem{
				FRN:         4,
				DataItem:    "I026/004",
				Description: "Repetitive type field for test",
				Type:        uap.Repetitive,
			},
			Repetitive: &Repetitive{
				Rep:  0x02,
				Data: []byte{0xff, 0xff, 0xff, 0xff},
			},
		},
		{
			Meta: MetaItem{
				FRN:         5,
				DataItem:    "I026/005",
				Description: "Compound type field for test",
				Type:        uap.Compound,
			},
			Compound: &Compound{
				Primary: []byte{0xab, 0x80},
				Secondary: []Item{
					{
						Meta: MetaItem{
							FRN:         1,
							DataItem:    "Compound/001",
							Description: "Compound Fixed type field for test",
							Type:        uap.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0xff}},
					},
					{
						Meta: MetaItem{
							FRN:         3,
							DataItem:    "Compound/003",
							Description: "Compound Extended type field for test",
							Type:        uap.Extended,
						},
						Extended: &Extended{
							Primary:   []byte{0xff},
							Secondary: []byte{0xfe},
						},
					},
					{
						Meta: MetaItem{
							FRN:         5,
							DataItem:    "Compound/005",
							Description: "Compound Repetitive type field for test",
							Type:        uap.Repetitive,
						},
						Repetitive: &Repetitive{
							Rep:  0x02,
							Data: []byte{0xff, 0xff, 0xff, 0xff},
						},
					},
					{
						Meta: MetaItem{
							FRN:         7,
							DataItem:    "Compound/007",
							Description: "Compound Explicit type field for test",
							Type:        uap.Explicit,
						},
						Explicit: &Explicit{
							Len:  0x04,
							Data: []byte{0xff, 0xff, 0xff},
						},
					},
					{
						Meta: MetaItem{
							FRN:         8,
							DataItem:    "Compound/008",
							Description: "Compound Fixed type field for test",
							Type:        uap.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0xff, 0xff}},
					},
				},
			},
		},
		{
			Meta: MetaItem{
				FRN:         6,
				DataItem:    "I026/006",
				Description: "RFS(Random Field Sequencing) type field for test",
				Type:        uap.RFS,
			},
			RFS: &RandomFieldSequencing{
				N: 0x01,
				Sequence: []RandomField{
					{
						FRN: 1,
						Field: Item{
							Meta: MetaItem{
								FRN:         1,
								DataItem:    "I026/001",
								Description: "Fixed type field for test",
								Type:        uap.Fixed,
							},
							Fixed: &Fixed{Data: []byte{0xff, 0xff}},
						},
					},
				},
			},
		},
		{
			Meta: MetaItem{
				FRN:         9,
				DataItem:    "SP",
				Description: "SP (Special Purpose field) type field for test",
				Type:        uap.SP,
			},
			SP: &SpecialPurpose{
				Len:  03,
				Data: []byte{0xff, 0xff},
			},
		},
	}
	uap4Test := uap.Cat4Test
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap4Test)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("FAIL: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", item, output[i])
		}
	}
}

func TestRecordDecode_Cat4TestTrackFullRecord(t *testing.T) {
	// Arrange
	input := "01 38 80ff ffff"
	output := []Item{
		{
			Meta: MetaItem{
				FRN:         10,
				DataItem:    "I026/010",
				Description: "Fixed type field for test",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x80}},
		},
		{
			Meta: MetaItem{
				FRN:         11,
				DataItem:    "I026/011",
				Description: "Fixed type field for test",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff}},
		},
		{
			Meta: MetaItem{
				FRN:         12,
				DataItem:    "I026/012",
				Description: "Fixed type field for test",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff, 0xff}},
		},
	}
	uap4Test := uap.Cat4Test
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap4Test)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("FAIL: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", item, output[i])
		}
	}
}

func TestRecordDecode_Cat4TestPlotFullRecord(t *testing.T) {
	// Arrange
	input := "01 38 00 ffffff ff"
	output := []Item{
		{
			Meta: MetaItem{
				FRN:         10,
				DataItem:    "I026/010",
				Description: "Fixed type field for test",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00}},
		},
		{
			Meta: MetaItem{
				FRN:         11,
				DataItem:    "I026/011",
				Description: "Fixed type field for test",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff, 0xff, 0xff}},
		},
		{
			Meta: MetaItem{
				FRN:         12,
				DataItem:    "I026/012",
				Description: "Fixed type field for test",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff}},
		},
	}
	uap4Test := uap.Cat4Test
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap4Test)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("FAIL: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", item, output[i])
		}
	}
}

func TestRecordDecode_Cat4TestError(t *testing.T) {
	// Setup
	type dataTest struct {
		TestCase string
		input    string
		output   []Item
		unRead   int
		err      error
	}
	dataSet := []dataTest{
		{
			TestCase: "testcase 1",
			input:    "02 FFFF",
			output:   nil,
			unRead:   2,
			err:      ErrDataFieldUnknown,
		},
		{
			TestCase: "testcase 2",
			// Repetitive FRN 4
			input:  "10 03FFFFFFFF",
			output: nil,
			unRead: 0,
			err:    io.ErrUnexpectedEOF,
		},
		{
			TestCase: "testcase 3",
			// Explicit FRN 3
			input:  "20 04FFFF",
			output: nil,
			unRead: 0,
			err:    io.ErrUnexpectedEOF,
		},
		{
			TestCase: "testcase 4",
			// RFS FRN 6
			input:  "04 0101",
			output: nil,
			unRead: 0,
			err:    io.EOF,
		},
		{
			TestCase: "testcase 5",
			// RE FRN 8
			input:  "0180 04FFFF",
			output: nil,
			unRead: 0,
			err:    io.ErrUnexpectedEOF,
		},
	}

	for _, row := range dataSet {
		// Arrange
		uap4Test := uap.Cat4Test
		data, _ := HexStringToByte(row.input)
		rec := NewRecord()

		// Act
		remaining, err := rec.Decode(data, uap4Test)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: %s - error = %v; Expected: %v", row.TestCase, err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}
		if remaining != row.unRead {
			t.Errorf("FAIL: %s - unRead = %v; Expected: %v", row.TestCase, remaining, row.unRead)
		} else {
			t.Logf("SUCCESS: unRead = %v; Expected: %v", remaining, row.unRead)
		}
		if reflect.DeepEqual(rec.Items, row.output) == false {
			t.Errorf("FAIL: %s - %v; Expected: %v", row.TestCase, rec.Items, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", rec.Items, row.output)
		}
	}
}

func TestRecordDecode_CAT048(t *testing.T) {
	// Arrange
	input := "fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5"
	output := []Item{
		{
			Meta: MetaItem{
				FRN:         1,
				DataItem:    "I048/010",
				Description: "Data Source Identifier",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x08, 0x36}},
		},
		{
			Meta: MetaItem{
				FRN:         2,
				DataItem:    "I048/140",
				Description: "Time-of-Day",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x42, 0x9b, 0x52}},
		},
		{
			Meta: MetaItem{
				FRN:         3,
				DataItem:    "I048/020",
				Description: "Target Report Descriptor",
				Type:        uap.Extended,
			},
			Extended: &Extended{
				Primary:   []byte{0xa0},
				Secondary: nil,
			},
		},
		{
			Meta: MetaItem{
				FRN:         4,
				DataItem:    "I048/040",
				Description: "Measured Position in Slant Polar Coordinates",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x94, 0xc7, 0x01, 0x81}},
		},
		{
			Meta: MetaItem{
				FRN:         5,
				DataItem:    "I048/070",
				Description: "Mode-3/A Code in Octal Representation",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x09, 0x13}},
		},
		{
			Meta: MetaItem{
				FRN:         6,
				DataItem:    "I048/090",
				Description: "Flight Level in Binary Representation",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x02, 0xd0}},
		},
		{
			Meta: MetaItem{
				FRN:         7,
				DataItem:    "I048/130",
				Description: "Radar Plot Characteristics",
				Type:        uap.Compound,
			},
			Compound: &Compound{
				Primary: []byte{0x60},
				Secondary: []Item{
					{
						Meta: MetaItem{
							FRN:         2,
							DataItem:    "SRR",
							Description: "Number of received replies",
							Type:        uap.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x02}},
					},
					{
						Meta: MetaItem{
							FRN:         3,
							DataItem:    "SAM",
							Description: "Amplitude of received replies for M(SSR)",
							Type:        uap.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0xb7}},
					},
				},
			},
		},
		{
			Meta: MetaItem{
				FRN:         8,
				DataItem:    "I048/220",
				Description: "Aircraft Address",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x49, 0x0d, 0x01}},
		},
		{
			Meta: MetaItem{
				FRN:         9,
				DataItem:    "I048/240",
				Description: "Aircraft Identification",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x38, 0xa1, 0x78, 0xcf, 0x42, 0x20}},
		},
		{
			Meta: MetaItem{
				FRN:         10,
				DataItem:    "I048/250",
				Description: "Mode S MB Data",
				Type:        uap.Repetitive,
			},
			Repetitive: &Repetitive{
				Rep:  0x02,
				Data: []byte{0xe7, 0x9a, 0x5d, 0x27, 0xa0, 0x0c, 0x00, 0x60, 0xa3, 0x28, 0x00, 0x30, 0xa4, 0x00, 0x00, 0x40},
			},
		},
		{
			Meta: MetaItem{
				FRN:         11,
				DataItem:    "I048/161",
				Description: "Track Number",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x06, 0x3a}},
		},
		{
			Meta: MetaItem{
				FRN:         13,
				DataItem:    "I048/200",
				Description: "Calculated Track Velocity in Polar Representation",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x07, 0x43, 0xce, 0x5b}},
		},
		{
			Meta: MetaItem{
				FRN:         14,
				DataItem:    "I048/170",
				Description: "Track Status",
				Type:        uap.Extended,
			},
			Extended: &Extended{
				Primary:   []byte{0x40},
				Secondary: nil,
			},
		},
		{
			Meta: MetaItem{
				FRN:         21,
				DataItem:    "I048/230",
				Description: "Communications / ACAS Capability and Flight Status",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x20, 0xf5}},
		},
	}

	uap048 := uap.Cat048V127
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap048)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("FAIL: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", item, output[i])
		}
	}
}

func TestRecordDecode_CAT063(t *testing.T) {
	// Arrange
	input := "bff0090c7cd2cc08294000000000000000000000000000000000"
	output := []Item{
		{
			Meta: MetaItem{
				FRN:         1,
				DataItem:    "I063/010",
				Description: "Data Source Identifier",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x09, 0x0c}},
		},
		{
			Meta: MetaItem{
				FRN:         3,
				DataItem:    "I063/030",
				Description: "Time of Message",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x7c, 0xd2, 0xcc}},
		},
		{
			Meta: MetaItem{
				FRN:         4,
				DataItem:    "I063/050",
				Description: "Sensor Identifier",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x08, 0x29}},
		},
		{
			Meta: MetaItem{
				FRN:         5,
				DataItem:    "I063/060",
				Description: "Sensor Configuration and Status",
				Type:        uap.Extended,
			},
			Extended: &Extended{
				Primary:   []byte{0x40},
				Secondary: nil,
			},
		},
		{
			Meta: MetaItem{
				FRN:         6,
				DataItem:    "I063/070",
				Description: "Time Stamping Bias",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00}},
		},
		{
			Meta: MetaItem{
				FRN:         7,
				DataItem:    "I063/080",
				Description: "SSR/Mode S Range Gain and Bias",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00, 0x00, 0x00}},
		},
		{
			Meta: MetaItem{
				FRN:         8,
				DataItem:    "I063/081",
				Description: "SSR/Mode S Azimuth Bias",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00}},
		},
		{
			Meta: MetaItem{
				FRN:         9,
				DataItem:    "I063/090",
				Description: "PSR Range Gain and Bias",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00, 0x00, 0x00}},
		},
		{
			Meta: MetaItem{
				FRN:         10,
				DataItem:    "I063/091",
				Description: "PSR Azimuth Bias",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00}},
		},
		{
			Meta: MetaItem{
				FRN:         11,
				DataItem:    "I063/092",
				Description: "PSR Elevation Bias",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00}},
		},
	}

	uap063 := uap.Cat063V16
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap063)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("FAIL: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", item, output[i])
		}
	}
}
func TestRecordDecode_CAT065(t *testing.T) {
	// Arrange
	input := "f8090c0203424cf30a"
	output := []Item{
		{
			Meta: MetaItem{
				FRN:         1,
				DataItem:    "I065/010",
				Description: "Data Source Identifier",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x09, 0x0c}},
		},
		{
			Meta: MetaItem{
				FRN:         2,
				DataItem:    "I065/000",
				Description: "Message Type",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x02}},
		},
		{
			Meta: MetaItem{
				FRN:         3,
				DataItem:    "I065/015",
				Description: "Service Identification",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x03}},
		},
		{
			Meta: MetaItem{
				FRN:         4,
				DataItem:    "I065/030",
				Description: "Time Of Message",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x42, 0x4c, 0xf3}},
		},
		{
			Meta: MetaItem{
				FRN:         5,
				DataItem:    "I065/020",
				Description: "Batch Number",
				Type:        uap.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x0a}},
		},
	}
	uap065 := uap.Cat065V15
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap065)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("FAIL: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", item, output[i])
		}
	}
}

/*
todo
func TestRecordDecode_CAT001Track(t *testing.T) {
	// Arrange
	input := "f502 0831 98 01bf 0a1ebb43 022538e2 00"
	output := [][]byte{
		{0x08, 0x31},
		{0x98},
		{0x01, 0xbf},
		{0x0a, 0x1e, 0xbb, 0x43},
		{0x02, 0x25, 0x38, 0xe2},
		{0x00},
	}

	uap001 := uap.Cat001V12
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap001)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		}
	}
}

func TestRecordDecode_CAT001Plot(t *testing.T) {
	// Arrange
	//input := "F0 08 31 08 0A 8A BB 2E 38 02"
	input := "f0 0831 00 0a8abb2e 3802"
	output := [][]byte{
		{0x08, 0x31},
		{0x00},
		{0x0a, 0x8a, 0xbb, 0x2e},
		{0x38, 0x02},
	}

	uap001 := uap.Cat001V12
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap001)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		}
	}
}

func TestRecordDecode_CAT002(t *testing.T) {
	// Arrange
	input := "f4 0839 02105fb35b02"
	output := [][]byte{
		{0x08, 0x39},
		{0x02},
		{0x10},
		{0x5f, 0xb3, 0x5b},
		{0x02},
	}

	uap002 := uap.Cat002V10
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap002)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		}
	}
}

func TestRecordDecode_CAT030STR(t *testing.T) {
	// Arrange
	input := "bfff0160 0885 5801b8 6092fc 010e 0200 0925f483 0c 04e6 04ea " +
		"fb5ff9c4 f8 fd9a 0d0174 48455b 2cc371cf1de0"
	output := [][]byte{
		{0x08, 0x85},
		{0x58, 0x01, 0xb8},
		{0x60, 0x92, 0xfc},
		{0x01, 0x0e},
		{0x02, 0x00},
		{0x09, 0x25, 0xf4, 0x83},
		{0x0c},
		{0x04, 0xe6},
		{0x04, 0xea},
		{0xfb, 0x5f, 0xf9, 0xc4},
		{0xf8},
		{0xfd, 0x9a},
		{0x0d, 0x01, 0x74},
		{0x48, 0x45, 0x5b},
		{0x2c, 0xc3, 0x71, 0xcf, 0x1d, 0xe0},
	}
	uap030 := uap.Cat030StrV51
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap030)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		}
	}
}

func TestRecordDecode_CAT032STR(t *testing.T) {
	// Arrange
	input := "d0 0884 3b5494 00130000008f002f008948006a007c"
	output := [][]byte{
		{0x08, 0x84},
		{0x3b, 0x54, 0x94},
		{0x00, 0x13, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x2f, 0x00, 0x89, 0x48, 0x00, 0x6a, 0x00, 0x7c},
	}

	uap030 := uap.Cat032StrV70
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap030)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		}
	}
}

func TestRecordDecode_CAT034(t *testing.T) {
	// Arrange
	input := "f6 0836 02 429b61 08 9400282000 94008000"
	output := [][]byte{
		{0x08, 0x36},
		{0x02},
		{0x42, 0x9b, 0x61},
		{0x08},
		{0x94, 0x00, 0x28, 0x20, 0x00},
		{0x94, 0x00, 0x80, 0x00},
	}
	uap048 := uap.Cat034V127
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap048)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		}
	}
}

func TestRecordDecode_CAT062(t *testing.T) {
	// Arrange
	input := "bf5ffd0304 090001532100008e6f3e0017d0961247f10b7086fed3019a0fc8e301010c87304a04e072c34820e300820800eb003104b2190301487fa0ff0614ffffffffffff0493110101c006061414141400e0045b00e00182dc622931a410a800e00fc84010e001622b05010d01622902fea60177"
	output := [][]byte{
		{0x09, 0x00},
		{0x01},
		{0x53, 0x21, 0x00},
		{0x00, 0x8e, 0x6f, 0x3e, 0x00, 0x17, 0xd0, 0x96},
		{0x12, 0x47, 0xf1, 0x0b, 0x70, 0x86},
		{0xfe, 0xd3, 0x01, 0x9a},
		{0x0f, 0xc8},
		{0xe3, 0x01, 0x01, 0x0c, 0x87, 0x30, 0x4a, 0x04, 0xe0, 0x72, 0xc3, 0x48, 0x20, 0xe3, 0x00, 0x82, 0x08, 0x00, 0xeb, 0x00, 0x31},
		{0x04, 0xb2},
		{0x19, 0x03, 0x01, 0x48},
		{0x7f, 0xa0, 0xff, 0x06, 0x14, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		{0x04},
		{0x93, 0x11, 0x01, 0x01, 0xc0, 0x06, 0x06, 0x14, 0x14, 0x14, 0x14},
		{0x00, 0xe0},
		{0x04, 0x5b},
		{0x00, 0xe0},
		{0x01, 0x82},
		{0xdc, 0x62, 0x29, 0x31, 0xa4, 0x10, 0xa8, 0x00, 0xe0, 0x0f, 0xc8, 0x40},
		{0x10, 0xe0, 0x01, 0x62, 0x2b, 0x05, 0x01, 0x0d, 0x01, 0x62, 0x29, 0x02, 0xfe, 0xa6, 0x01, 0x77},
	}

	uap062 := uap.Cat062V119
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap062)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		}
	}
}

func TestRecordDecode_CAT255STR(t *testing.T) {
	// Arrange
	input := "e0 08 83 7dfd9c 58"
	output := [][]byte{
		{0x08, 0x83},
		{0x7d, 0xfd, 0x9c},
		{0x58},
	}

	uap255 := uap.Cat255StrV51
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap255)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		}
	}
}

func TestRecordDecode_CAT030ARTAS(t *testing.T) {
	// Arrange
	input := "afbbf317f130 0883 04 0070 a8bcf3 ff070707 23f0a880 0713feb7 022b 0389 038b 14 07 04 012c 0808 " +
		"11580000001e7004f04aa004b001240054 4e494135313132 06c8 4c45424c 48454c58 4d413332300101a5389075c71ca0"

	output := [][]byte{
		{0x08, 0x83},
		{0x04},
		{0x00, 0x70},
		{0xa8, 0xbc, 0xf3},
		{0xff, 0x07, 0x07, 0x07},
		{0x23, 0xf0, 0xa8, 0x80},
		{0x07, 0x13, 0xfe, 0xb7},
		{0x02, 0x2b},
		{0x03, 0x89},
		{0x03, 0x8b},
		{0x14},
		{0x07},
		{0x04},
		{0x01, 0x2c},
		{0x08, 0x08},
		{0x11, 0x58, 0x00, 0x00, 0x00, 0x1e, 0x70, 0x04, 0xf0, 0x4a, 0xa0, 0x04, 0xb0, 0x01, 0x24, 0x00, 0x54},
		{0x4e, 0x49, 0x41, 0x35, 0x31, 0x31, 0x32},
		{0x06, 0xc8},
		{0x4c, 0x45, 0x42, 0x4c},
		{0x48, 0x45, 0x4c, 0x58},
		{0x4d},
		{0x41, 0x33, 0x32, 0x30},
		{0x01, 0x01, 0xa5},
		{0x38, 0x90, 0x75, 0xc7, 0x1c, 0xa0},
	}

	uap030 := uap.Cat030ArtasV62
	data, _ := HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap030)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.Items {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Data, output[i])
		}
	}
}
*/

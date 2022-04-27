package goasterix

import (
	"bytes"
	. "github.com/mokhtarimokhtar/goasterix/item"
	"github.com/mokhtarimokhtar/goasterix/util"
	"reflect"
	"testing"
)

/*func TestRecordPayload(t *testing.T) {
	// Arrange
	data, _ := util.HexStringToByte("ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020a0")
	nbOfBytes := 42
	rec := NewRecord()
	_, _ = rec.Decode(data, dataField.Cat048V127)

	// Act
	items := rec.Payload()

	// Assert
	if len(items) != nbOfBytes {
		t.Errorf("MsgFailInValue: len(items) = %v; Expected: %v", len(items), nbOfBytes)
	} else {
		t.Logf("MsgSuccessInValue: len(items) = %v; Expected: %v", len(items), nbOfBytes)
	}
}*/
/*
func TestRecord_String(t *testing.T) {
	// Arrange
	data, _ := util.HexStringToByte("ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020a0")
	nbOfItems := 15
	rec := NewRecord()
	_, _ = rec.Decode(data, _uap.Cat048V127)

	// Act
	items := rec.String()

	// Assert
	if len(items) != nbOfItems {
		t.Errorf("MsgFailInValue: len(items) = %v; Expected: %v", len(items), nbOfItems)
	} else {
		t.Logf("MsgSuccessInValue: len(items) = %v; Expected: %v", len(items), nbOfItems)
	}
}
*/

// Testing : Decode CatForTest
func TestDecodeCatForTest(t *testing.T) {
	type testCase struct {
		Name   string
		input  string // one record = fspec + items
		uap    StandardUAP
		output Record
		unRead int
		err    error
	}
	dataSet := []testCase{
		{
			Name:  "testcase 1",
			input: "f780 ffff 01 0302 0801020304050607 03aaaaaabbbbbbcccccc  b80101010202aaaabbbb0201 0201 04010203",
			uap:   CatForTest, // f780 1111-0111 1000-0000  // b80101010202aaaabbbb0201
			output: Record{
				Cat:   26,
				Fspec: []byte{0xf7, 0x80},
				DataItems: []DataItem{
					&Fixed{
						Base: Base{
							FRN:          1,
							DataItemName: "I026/010",
							Description:  "Fixed type field for test",
							Type:         FixedField,
						},
						Size: 2,
						//Data: []byte{0xff, 0xff},
						SubItems: []SubItem{
							&SubItemFromTo{
								Name: "SAC",
								Type: FromToField,
								Pos:  BitPosition{From: 16, To: 9},
								Data: []byte{0xff},
							},
							&SubItemFromTo{
								Name: "SIC",
								Type: FromToField,
								Pos:  BitPosition{From: 8, To: 1},
								Data: []byte{0xff},
							},
						},
					},
					&Extended{
						Base: Base{
							FRN:          2,
							DataItemName: "I026/020",
							Description:  "Extended type field for test",
							Type:         ExtendedField,
						},
						PrimaryItemSize:   1,
						SecondaryItemSize: 2,
						Primary:           []byte{0x01},
						Secondary:         []byte{0x03, 0x02},
						/*
							SubItems: []SubItem{
								&SubItemFromTo{
									Name: "TYP",
									Type: FromToField,
									From: 8,
									To:   6,
								},
								&SubItemBit{Name: "SIM", Type: BitField, Pos: 5},
								&SubItemBit{Name: "RDP", Type: BitField, Pos: 4},
								&SubItemBit{Name: "SPI", Type: BitField, Pos: 3},
								&SubItemBit{Name: "RAB", Type: BitField, Pos: 2},
								&SubItemBit{Name: "TST", Type: BitField, Pos: 8},
								&SubItemBit{Name: "ERR", Type: BitField, Pos: 7},
								&SubItemBit{Name: "XPP", Type: BitField, Pos: 6},
								&SubItemBit{Name: "ME", Type: BitField, Pos: 5},
								&SubItemBit{Name: "MI", Type: BitField, Pos: 4},
								&SubItemFromTo{
									Name: "FOE/FRI",
									Type: FromToField,
									From: 3,
									To:   2,
								},
							},
						*/
					},
					&Explicit{
						Base: Base{
							FRN:          3,
							DataItemName: "I026/030",
							Description:  "Explicit type field for test",
							Type:         ExplicitField,
						},
						Len:  0x08,
						Data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
					},
					&Repetitive{
						Base: Base{
							FRN:          4,
							DataItemName: "I026/040",
							Description:  "Repetitive type field for test",
							Type:         RepetitiveField,
						},
						SubItemSize: 3,
						Rep:         0x03,
						//Data:        []byte{0xaa, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xcc, 0xcc, 0xcc},
						Data: nil,
						SubItems: []SubItem{
							&SubItemFromTo{
								Name: "DOP",
								Type: FromToField,
								Pos:  BitPosition{From: 24, To: 17},
								Data: []byte{0xaa},
							},
							&SubItemFromTo{
								Pos:  BitPosition{From: 16, To: 9},
								Name: "AMB",
								Type: FromToField,
								Data: []byte{0xaa},
							},
							&SubItemFromTo{
								Name: "FRQ",
								Type: FromToField,
								Pos:  BitPosition{From: 8, To: 1},
								Data: []byte{0xaa},
							},
							&SubItemFromTo{
								Name: "DOP",
								Type: FromToField,
								Pos:  BitPosition{From: 24, To: 17},
								Data: []byte{0xbb},
							},
							&SubItemFromTo{
								Name: "AMB",
								Type: FromToField,
								Pos:  BitPosition{From: 16, To: 9},
								Data: []byte{0xbb},
							},
							&SubItemFromTo{
								Name: "FRQ",
								Type: FromToField,
								Pos:  BitPosition{From: 8, To: 1},
								Data: []byte{0xbb},
							},
							&SubItemFromTo{
								Name: "DOP",
								Type: FromToField,
								Pos:  BitPosition{From: 24, To: 17},
								Data: []byte{0xcc},
							},
							&SubItemFromTo{
								Name: "AMB",
								Type: FromToField,
								Pos:  BitPosition{From: 16, To: 9},
								Data: []byte{0xcc},
							},
							&SubItemFromTo{
								Name: "FRQ",
								Type: FromToField,
								Pos:  BitPosition{From: 8, To: 1},
								Data: []byte{0xcc},
							},
						},
					},
					&Compound{
						Base: Base{
							FRN:          6,
							DataItemName: "I026/060",
							Description:  "Compound type field for test",
							Type:         CompoundField,
						},
						Primary: []byte{0xb8},
						Secondary: []DataItem{
							&Fixed{
								Base: Base{
									FRN:          1,
									DataItemName: "Compound/001",
									Description:  "Compound Fixed type field for test",
									Type:         FixedField,
								},
								Size: 1,
								Data: []byte{0x01},
							},
							&Extended{
								Base: Base{
									FRN:          3,
									DataItemName: "Compound/003",
									Description:  "Compound Extended type field for test",
									Type:         ExtendedField,
								},
								PrimaryItemSize:   1,
								SecondaryItemSize: 1,
								Primary:           []byte{0x01},
								Secondary:         []byte{0x01, 0x02},
							},
							&Repetitive{
								Base: Base{
									FRN:          4,
									DataItemName: "Compound/004",
									Description:  "Compound Repetitive type field for test",
									Type:         RepetitiveField,
								},
								SubItemSize: 2,
								Rep:         0x02,
								Data:        []byte{0xaa, 0xaa, 0xbb, 0xbb},
							},
							&Explicit{
								Base: Base{
									FRN:          5,
									DataItemName: "Compound/005",
									Description:  "Compound Explicit type field for test",
									Type:         ExplicitField,
								},
								Len:  0x02,
								Data: []byte{0x01},
							},
						},
					},
					&ReservedExpansion{
						Base: Base{
							FRN:          7,
							DataItemName: "RE",
							Description:  "Reserved Expansion type field for test",
							Type:         REField,
						},
						Len:  0x02,
						Data: []byte{0x01},
					},
					&SpecialPurpose{
						Base: Base{
							FRN:          8,
							DataItemName: "SP",
							Description:  "Special Purpose type field for test",
							Type:         SPField,
						},
						Len:  0x04,
						Data: []byte{0x01, 0x02, 0x03},
					},
				},
			},
			unRead: 0,
			err:    nil,
		},
	}

	for _, tc := range dataSet {
		// Arrange
		data, _ := util.HexStringToByte(tc.input)
		rec := new(Record)

		// Act
		unRead, err := rec.Decode(data, tc.uap)

		// Assert
		if err != tc.err {
			t.Errorf(util.MsgFailInValue, tc.Name, err, tc.err)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, err, tc.err)
		}

		if unRead != tc.unRead {
			t.Errorf(util.MsgFailInValue, tc.Name, unRead, tc.unRead)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, unRead, tc.unRead)
		}

		if rec.Cat != tc.output.Cat {
			t.Errorf(util.MsgFailInValue, tc.Name, rec.Cat, tc.output.Cat)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, rec.Cat, tc.output.Cat)
		}
		if bytes.Equal(rec.Fspec, tc.output.Fspec) == false {
			t.Errorf(util.MsgFailInHex, tc.Name, rec.Fspec, tc.output.Fspec)
		} else {
			t.Logf(util.MsgSuccessInHex, tc.Name, rec.Fspec, tc.output.Fspec)
		}
		for i, item := range rec.DataItems {
			if reflect.DeepEqual(item, tc.output.DataItems[i]) == false {
				t.Errorf(util.MsgFailInValue, tc.Name, item, tc.output.DataItems[i])
			} else {
				t.Logf(util.MsgSuccessInValue, tc.Name, item, tc.output.DataItems[i])
			}
		}
	}
}

// Testing Integration : Decode by category
func TestRecordDecodeCAT048(t *testing.T) {
	// Arrange
	input := "fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5"
	output := []DataItem{
		&Fixed{
			Base: Base{
				FRN:          1,
				DataItemName: "I048/010",
				Description:  "Data Source Identifier",
				Type:         FixedField,
			},
			Size: 2,
			SubItems: []SubItem{
				&SubItemFromTo{
					Name: "SAC",
					Type: FromToField,
					Pos:  BitPosition{From: 16, To: 9},
					Data: []byte{0x08},
				},
				&SubItemFromTo{
					Name: "SIC",
					Type: FromToField,
					Pos:  BitPosition{From: 8, To: 1},
					Data: []byte{0x36},
				},
			},
		},
		&Fixed{
			Base: Base{
				FRN:          2,
				DataItemName: "I048/140",
				Description:  "Time-of-Day",
				Type:         FixedField,
			},
			Size: 3,
			Data: []byte{0x42, 0x9b, 0x52},
		},
		&Extended{
			Base: Base{
				FRN:          3,
				DataItemName: "I048/020",
				Description:  "Target Report Descriptor",
				Type:         ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 1,
			Primary:           []byte{0xa0},
			Secondary:         nil,
		},
		&Fixed{
			Base: Base{
				FRN:          4,
				DataItemName: "I048/040",
				Description:  "Measured Position in Slant Polar Coordinates",
				Type:         FixedField,
			},
			Size: 4,
			Data: []byte{0x94, 0xc7, 0x01, 0x81},
		},
		&Fixed{
			Base: Base{
				FRN:          5,
				DataItemName: "I048/070",
				Description:  "Mode-3/A Code in Octal Representation",
				Type:         FixedField,
			},
			Size: 2,
			Data: []byte{0x09, 0x13},
		},
		&Fixed{
			Base: Base{
				FRN:          6,
				DataItemName: "I048/090",
				Description:  "Flight Level in Binary Representation",
				Type:         FixedField,
			},
			Size: 2,
			Data: []byte{0x02, 0xd0},
		},
		&Compound{
			Base: Base{
				FRN:          7,
				DataItemName: "I048/130",
				Description:  "Radar Plot Characteristics",
				Type:         CompoundField,
			},
			Primary: []byte{0x60},
			Secondary: []DataItem{
				&Fixed{
					Base: Base{
						FRN:          2,
						DataItemName: "SRR",
						Description:  "Number of received replies",
						Type:         FixedField,
					},
					Size: 1,
					Data: []byte{0x02},
				},
				&Fixed{
					Base: Base{
						FRN:          3,
						DataItemName: "SAM",
						Description:  "Amplitude of received replies for M(SSR)",
						Type:         FixedField,
					},
					Size: 1,
					Data: []byte{0xb7},
				},
			},
		},
		&Fixed{
			Base: Base{
				FRN:          8,
				DataItemName: "I048/220",
				Description:  "Aircraft Address",
				Type:         FixedField,
			},
			Size: 3,
			Data: []byte{0x49, 0x0d, 0x01},
		},
		&Fixed{
			Base: Base{
				FRN:          9,
				DataItemName: "I048/240",
				Description:  "Aircraft Identification",
				Type:         FixedField,
			},
			Size: 6,
			Data: []byte{0x38, 0xa1, 0x78, 0xcf, 0x42, 0x20},
		},
		&Repetitive{
			Base: Base{
				FRN:          10,
				DataItemName: "I048/250",
				Description:  "Mode S MB Data",
				Type:         RepetitiveField,
			},
			SubItemSize: 8,
			Rep:         0x02,
			Data:        []byte{0xe7, 0x9a, 0x5d, 0x27, 0xa0, 0x0c, 0x00, 0x60, 0xa3, 0x28, 0x00, 0x30, 0xa4, 0x00, 0x00, 0x40},
		},
		&Fixed{
			Base: Base{
				FRN:          11,
				DataItemName: "I048/161",
				Description:  "Track Number",
				Type:         FixedField,
			},
			Size: 2,
			Data: []byte{0x06, 0x3a},
		},
		&Fixed{
			Base: Base{
				FRN:          13,
				DataItemName: "I048/200",
				Description:  "Calculated Track Velocity in Polar Representation",
				Type:         FixedField,
			},
			Size: 4,
			Data: []byte{0x07, 0x43, 0xce, 0x5b},
		},
		&Extended{
			Base: Base{
				FRN:          14,
				DataItemName: "I048/170",
				Description:  "Track Status",
				Type:         ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 1,
			Primary:           []byte{0x40},
			Secondary:         nil,
		},
		&Fixed{
			Base: Base{
				FRN:          21,
				DataItemName: "I048/230",
				Description:  "Communications / ACAS Capability and Flight Status",
				Type:         FixedField,
			},
			Size: 2,
			Data: []byte{0x20, 0xf5},
		},
	}

	uap048 := Cat048V127
	data, _ := util.HexStringToByte(input)
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
	for i, item := range rec.DataItems {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("FAIL: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", item, output[i])
		}
	}
}

/*
// RFSDataField
func TestRFSDataFieldReader(t *testing.T) {
	// Setup
	type testCase struct {
		TestCaseName string
		input        string
		item         []dataField.dataField
		output       RandomFieldSequencing
		err          error
	}
	dataSet := []testCase{
		{
			TestCaseName: "testCase 1",
			input:        "02 03 ffffffff 0a ff",
			item:         dataField.Cat001PlotV12,
			output: RandomFieldSequencing{
				N: 0x02,
				Sequence: []RandomField{
					{
						FRN: 0x03,
						Field: Item{
							Meta: Base{
								FRN:         3,
								DataItemName:    "I001/040",
								Description: "Measured Position in Polar Coordinates",
								Type:        dataField.Fixed,
							},
							Fixed: &Fixed{Data: []byte{0xff, 0xff, 0xff, 0xff}},
						},
					},
					{
						FRN: 0x0a,
						Field: Item{
							Meta: Base{
								FRN:         10,
								DataItemName:    "I001/131",
								Description: "Received Power",
								Type:        dataField.Fixed,
							},
							Fixed: &Fixed{Data: []byte{0xff}},
						},
					},
				},
			},
			err: nil,
		},
		{
			TestCaseName: "testCase 2",
			input:        "02",
			item:         dataField.Cat001PlotV12,
			output: RandomFieldSequencing{
				N: 0x02,
			},
			err: io.EOF,
		},
		{
			TestCaseName: "testCase 3",
			input:        "02 03 ffffffff 0a",
			item:         dataField.Cat001PlotV12,
			output: RandomFieldSequencing{
				N: 0x02,
				Sequence: []RandomField{
					{
						FRN: 0x03,
						Field: Item{
							Meta: Base{
								FRN:         3,
								DataItemName:    "I001/040",
								Description: "Measured Position in Polar Coordinates",
								Type:        dataField.Fixed,
							},
							Fixed: &Fixed{Data: []byte{0xff, 0xff, 0xff, 0xff}},
						},
					},
				},
			},
			err: io.EOF,
		},
		{
			TestCaseName: "testCase 4",
			input:        "",
			item:         dataField.Cat001PlotV12,
			output:       RandomFieldSequencing{},
			err:          io.EOF,
		},
	}

	for _, row := range dataSet {
		// Arrange
		input, _ := util.HexStringToByte(row.input)
		rb := bytes.NewReader(input)

		// Act
		cp, err := RFSDataFieldReader(rb, row.item)

		// Assert
		if err != row.err {
			t.Errorf("MsgFailInValue: error: %v; Expected: %v", err, row.err)
		} else {
			t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, row.err)
		}
		if reflect.DeepEqual(cp, row.output) == false {
			t.Errorf("MsgFailInValue: %s - \nCompound = %v;\nExpected: %v", row.TestCaseName, cp, row.output)
		} else {
			t.Logf("MsgSuccessInValue: Compound = %v; Expected: %v", cp, row.output)
		}
	}
}
*/
/*
// Testing integration record
func TestRecordDecodeNbOfItems(t *testing.T) {
	// setup
	type testCase struct {
		Name      string
		input     string           // data test one record = fspec + items
		uap       _uap.StandardUAP // DataItems of category corresponding to data test input
		nbOfItems int
		err       error // error expected
	}
	dataSet := []testCase{
		{
			Name:      "testcase 1",
			input:     "f6083602429b7110940028200094008000",
			uap:       _uap.Cat034V127,
			err:       nil,
			nbOfItems: 6,
		},
		{
			Name:      "testcase 2",
			input:     "f6083602429b71109400282000940080",
			uap:       _uap.Cat034V127,
			err:       io.EOF,
			nbOfItems: 5,
		},
		{
			Name:      "testcase 3",
			input:     "ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020a0",
			uap:       _uap.Cat048V127,
			err:       nil,
			nbOfItems: 14,
		},
		{
			Name:      "testcase 4", // 0xA0 last byte is removed
			input:     "ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020",
			uap:       _uap.Cat048V127,
			err:       io.ErrUnexpectedEOF,
			nbOfItems: 13,
		},
		{
			Name:      "testcase 5",
			input:     "f0 0831 00 0a8abb2e 3802",
			uap:       _uap.Cat001V12,
			err:       nil,
			nbOfItems: 4,
		},
		{
			Name:      "testcase 6",
			input:     "f0 0831 00 0a8abb2e 38",
			uap:       _uap.Cat001V12,
			err:       io.ErrUnexpectedEOF,
			nbOfItems: 3,
		},
		{
			Name:      "testcase 7",
			input:     "f502 0831 98 01bf 0a1ebb43 022538e2 00",
			uap:       _uap.Cat001V12,
			err:       nil,
			nbOfItems: 6,
		},
		{
			Name:      "testcase 8",
			input:     "f502 0831 98 01bf 0a1ebb43 022538e2",
			uap:       _uap.Cat001V12,
			err:       io.EOF,
			nbOfItems: 5,
		},
		{
			Name:      "testcase 9",
			input:     "",
			uap:       _uap.Cat048V127,
			nbOfItems: 0,
			err:       io.EOF,
		},
	}

	for _, row := range dataSet {
		// Arrange
		data, _ := util.HexStringToByte(row.input)
		rec := new(Record)

		// Act
		unRead, err := rec.Decode(data, row.uap)

		// Assert
		if err != row.err {
			t.Errorf(util.MsgFailInValue, row.Name, err, row.err)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, err, row.err)
		}

		if unRead != 0 {
			t.Errorf(util.MsgFailInValue, row.Name, unRead, 0)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, unRead, 0)
		}

		if row.nbOfItems != len(rec.DataItems) {
			t.Errorf(util.MsgFailInValue, row.Name, row.nbOfItems, len(rec.DataItems))
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, row.nbOfItems, len(rec.DataItems))
		}
	}
}
*/
/*
func TestRecordPayload(t *testing.T) {
	// setup
	type testCase struct {
		Name   string
		input  string           // data test one record = fspec + items
		uap    _uap.StandardUAP // DataItems of category corresponding to data test input
		output []byte
		err    error // error expected
	}
	dataSet := []testCase{
		{
			Name:  "testcase 1",
			input: "f6083602429b7110940028200094008000",
			uap:   _uap.Cat034V127,
			err:   nil,
			output: []byte{0xf6, 0x08, 0x36, 0x02, 0x42, 0x9b, 0x71, 0x10, 0x94, 0x00, 0x28, 0x20, 0x00, 0x94, 0x00,
				0x80, 0x00},
		},
		{
			Name:   "testcase 2", // 0x00 last byte is removed
			input:  "f6083602429b71109400282000940080",
			uap:    _uap.Cat034V127,
			err:    io.EOF,
			output: []byte{0xf6, 0x08, 0x36, 0x02, 0x42, 0x9b, 0x71, 0x10, 0x94, 0x00, 0x28, 0x20, 0x00},
		},
		{
			Name:  "testcase 3",
			input: "ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020a0",
			uap:   _uap.Cat048V127,
			err:   nil,
			output: []byte{0xff, 0xdf, 0x02, 0x93, 0x19, 0x37, 0x8d, 0x3d, 0xa2, 0x05, 0x6f, 0x13, 0x2d, 0x0f, 0xff,
				0x00, 0x94, 0x60, 0x02, 0xde, 0x50, 0x6f, 0x84, 0x4c, 0xc3, 0xc3, 0x51, 0x23, 0x31, 0x00, 0x17, 0x01,
				0x3b, 0x02, 0x6c, 0x00, 0x0c, 0x74, 0xa7, 0x40, 0x20, 0xa0},
		},
		{
			Name:  "testcase 4", // 0xa0 last byte is removed
			input: "ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020",
			uap:   _uap.Cat048V127,
			err:   io.ErrUnexpectedEOF,
			output: []byte{0xff, 0xdf, 0x02, 0x93, 0x19, 0x37, 0x8d, 0x3d, 0xa2, 0x05, 0x6f, 0x13, 0x2d, 0x0f, 0xff,
				0x00, 0x94, 0x60, 0x02, 0xde, 0x50, 0x6f, 0x84, 0x4c, 0xc3, 0xc3, 0x51, 0x23, 0x31, 0x00, 0x17, 0x01,
				0x3b, 0x02, 0x6c, 0x00, 0x0c, 0x74, 0xa7, 0x40},
		},
		{
			Name:   "testcase 5",
			input:  "f0 0831 00 0a8abb2e 3802",
			uap:    _uap.Cat001V12,
			err:    nil,
			output: []byte{0xf0, 0x08, 0x31, 0x00, 0x0a, 0x8a, 0xbb, 0x2e, 0x38, 0x02},
		},
		{
			Name:   "testcase 6",
			input:  "f0 0831 00 0a8abb2e 38",
			uap:    _uap.Cat001V12,
			err:    io.ErrUnexpectedEOF,
			output: []byte{0xf0, 0x08, 0x31, 0x00, 0x0a, 0x8a, 0xbb, 0x2e},
		},
		{
			Name:   "testcase 7",
			input:  "f502 0831 98 01bf 0a1ebb43 022538e2 00",
			uap:    _uap.Cat001V12,
			err:    nil,
			output: []byte{0xf5, 0x02, 0x08, 0x31, 0x98, 0x01, 0xbf, 0x0a, 0x1e, 0xbb, 0x43, 0x02, 0x25, 0x38, 0xe2, 0x00},
		},
		{
			Name:   "testcase 8",
			input:  "f502 0831 98 01bf 0a1ebb43 022538e2",
			uap:    _uap.Cat001V12,
			err:    io.EOF,
			output: []byte{0xf5, 0x02, 0x08, 0x31, 0x98, 0x01, 0xbf, 0x0a, 0x1e, 0xbb, 0x43, 0x02, 0x25, 0x38, 0xe2},
		},
		{
			Name:   "testcase 9",
			input:  "",
			uap:    _uap.Cat048V127,
			output: []byte{},
			err:    io.EOF,
		},
	}

	for _, row := range dataSet {
		// Arrange
		data, _ := util.HexStringToByte(row.input)
		rec := new(Record)
		_, err := rec.Decode(data, row.uap)

		// Act
		res := rec.Payload()

		// Assert
		if err != row.err {
			t.Errorf(util.MsgFailInValue, row.Name, err, row.err)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, err, row.err)
		}

		if bytes.Equal(res, row.output) == false {
			t.Errorf(util.MsgFailInHex, row.Name, res, row.output)
		} else {
			t.Logf(util.MsgSuccessInHex, row.Name, res, row.output)
		}
	}
}
*/
/*
// Testing : Decode CatForTest
func TestRecordDecode_Cat4TestFullRecord(t *testing.T) {
	// Arrange
	input := "fd 40 ffff fffffe 03ffff 02ffffffff ab80 ff fffe 02ffffffff 04ffffff ffff 0101ffff 03ffff"
	output := []Item{
		{
			Meta: Base{
				FRN:         1,
				DataItemName:    "I026/001",
				Description: "Fixed type dataField for test",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff, 0xff}},
		},
		{
			Meta: Base{
				FRN:         2,
				DataItemName:    "I026/002",
				Description: "Extended type dataField for test",
				Type:        dataField.Extended,
			},
			Extended: &Extended{
				Primary:   []byte{0xff},
				Secondary: []byte{0xff, 0xfe},
			},
		},
		{
			Meta: Base{
				FRN:         3,
				DataItemName:    "I026/003",
				Description: "Explicit type dataField for test",
				Type:        dataField.Explicit,
			},
			Explicit: &Explicit{
				Len:  0x03,
				Data: []byte{0xff, 0xff},
			},
		},
		{
			Meta: Base{
				FRN:         4,
				DataItemName:    "I026/004",
				Description: "Repetitive type dataField for test",
				Type:        dataField.Repetitive,
			},
			Repetitive: &Repetitive{
				Rep:  0x02,
				Data: []byte{0xff, 0xff, 0xff, 0xff},
			},
		},
		{
			Meta: Base{
				FRN:         5,
				DataItemName:    "I026/005",
				Description: "Compound type dataField for test",
				Type:        dataField.Compound,
			},
			Compound: &Compound{
				Primary: []byte{0xab, 0x80},
				Secondary: []Item{
					{
						Meta: Base{
							FRN:         1,
							DataItemName:    "Compound/001",
							Description: "Compound Fixed type dataField for test",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0xff}},
					},
					{
						Meta: Base{
							FRN:         3,
							DataItemName:    "Compound/003",
							Description: "Compound Extended type dataField for test",
							Type:        dataField.Extended,
						},
						Extended: &Extended{
							Primary:   []byte{0xff},
							Secondary: []byte{0xfe},
						},
					},
					{
						Meta: Base{
							FRN:         5,
							DataItemName:    "Compound/005",
							Description: "Compound Repetitive type dataField for test",
							Type:        dataField.Repetitive,
						},
						Repetitive: &Repetitive{
							Rep:  0x02,
							Data: []byte{0xff, 0xff, 0xff, 0xff},
						},
					},
					{
						Meta: Base{
							FRN:         7,
							DataItemName:    "Compound/007",
							Description: "Compound Explicit type dataField for test",
							Type:        dataField.Explicit,
						},
						Explicit: &Explicit{
							Len:  0x04,
							Data: []byte{0xff, 0xff, 0xff},
						},
					},
					{
						Meta: Base{
							FRN:         8,
							DataItemName:    "Compound/008",
							Description: "Compound Fixed type dataField for test",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0xff, 0xff}},
					},
				},
			},
		},
		{
			Meta: Base{
				FRN:         6,
				DataItemName:    "I026/006",
				Description: "RFS(Random Field Sequencing) type dataField for test",
				Type:        dataField.RFS,
			},
			RFS: &RandomFieldSequencing{
				N: 0x01,
				Sequence: []RandomField{
					{
						FRN: 1,
						Field: Item{
							Meta: Base{
								FRN:         1,
								DataItemName:    "I026/001",
								Description: "Fixed type dataField for test",
								Type:        dataField.Fixed,
							},
							Fixed: &Fixed{Data: []byte{0xff, 0xff}},
						},
					},
				},
			},
		},
		{
			Meta: Base{
				FRN:         9,
				DataItemName:    "SP",
				Description: "SP (Special Purpose dataField) type dataField for test",
				Type:        dataField.SP,
			},
			SP: &SpecialPurpose{
				Len:  03,
				Data: []byte{0xff, 0xff},
			},
		},
	}
	uap4Test := dataField.CatForTest
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap4Test)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("MsgFailInValue: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %v; Expected: %v", item, output[i])
		}
	}
}

func TestRecordDecode_Cat4TestTrackFullRecord(t *testing.T) {
	// Arrange
	input := "01 38 80ff ffff"
	output := []Item{
		{
			Meta: Base{
				FRN:         10,
				DataItemName:    "I026/010",
				Description: "Fixed type dataField for test",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x80}},
		},
		{
			Meta: Base{
				FRN:         11,
				DataItemName:    "I026/011",
				Description: "Fixed type dataField for test",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff}},
		},
		{
			Meta: Base{
				FRN:         12,
				DataItemName:    "I026/012",
				Description: "Fixed type dataField for test",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff, 0xff}},
		},
	}
	uap4Test := dataField.CatForTest
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap4Test)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("MsgFailInValue: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %v; Expected: %v", item, output[i])
		}
	}
}

func TestRecordDecode_Cat4TestPlotFullRecord(t *testing.T) {
	// Arrange
	input := "01 38 00 ffffff ff"
	output := []Item{
		{
			Meta: Base{
				FRN:         10,
				DataItemName:    "I026/010",
				Description: "Fixed type dataField for test",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00}},
		},
		{
			Meta: Base{
				FRN:         11,
				DataItemName:    "I026/011",
				Description: "Fixed type dataField for test",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff, 0xff, 0xff}},
		},
		{
			Meta: Base{
				FRN:         12,
				DataItemName:    "I026/012",
				Description: "Fixed type dataField for test",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff}},
		},
	}
	uap4Test := dataField.CatForTest
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap4Test)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("MsgFailInValue: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %v; Expected: %v", item, output[i])
		}
	}
}

func TestRecordDecode_Cat4TestError(t *testing.T) {
	// Setup
	type testCase struct {
		TestCase string
		input    string
		output   []Item
		unRead   int
		err      error
	}
	dataSet := []testCase{
		{
			TestCase: "testCase 1",
			input:    "02 FFFF",
			output:   nil,
			unRead:   2,
			err:      ErrDataFieldUnknown,
		},
		{
			TestCase: "testCase 2",
			// Repetitive FRN 4
			input:  "10 03FFFFFFFF",
			output: nil,
			unRead: 0,
			err:    io.ErrUnexpectedEOF,
		},
		{
			TestCase: "testCase 3",
			// Explicit FRN 3
			input:  "20 04FFFF",
			output: nil,
			unRead: 0,
			err:    io.ErrUnexpectedEOF,
		},
		{
			TestCase: "testCase 4",
			// RFS FRN 6
			input:  "04 0101",
			output: nil,
			unRead: 0,
			err:    io.EOF,
		},
		{
			TestCase: "testCase 5",
			// RE FRN 8
			input:  "0180 04FFFF",
			output: nil,
			unRead: 0,
			err:    io.ErrUnexpectedEOF,
		},
	}

	for _, row := range dataSet {
		// Arrange
		uap4Test := dataField.CatForTest
		data, _ := util.HexStringToByte(row.input)
		rec := NewRecord()

		// Act
		remaining, err := rec.Decode(data, uap4Test)

		// Assert
		if err != row.err {
			t.Errorf("MsgFailInValue: %s - error = %v; Expected: %v", row.TestCase, err, row.err)
		} else {
			t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, row.err)
		}
		if remaining != row.unRead {
			t.Errorf("MsgFailInValue: %s - unRead = %v; Expected: %v", row.TestCase, remaining, row.unRead)
		} else {
			t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", remaining, row.unRead)
		}
		if reflect.DeepEqual(rec.DataItems, row.output) == false {
			t.Errorf("MsgFailInValue: %s - %v; Expected: %v", row.TestCase, rec.DataItems, row.output)
		} else {
			t.Logf("MsgSuccessInValue: %v; Expected: %v", rec.DataItems, row.output)
		}
	}
}
*/
/*
// Testing : Decode by category
func TestRecordDecodeCAT048(t *testing.T) {
	// Arrange
	input := "fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5"
	output := []Item{
		&Fixed{
			Base: Base{
				FRN:         1,
				DataItemName:    "I048/010",
				Description: "Data Source Identifier",
				Type:        _uap.Fixed,
			},
			Size: 2,
			Data: []byte{0x08, 0x36},
		},
		&Fixed{
			Base: Base{
				FRN:         2,
				DataItemName:    "I048/140",
				Description: "Time-of-Day",
				Type:        _uap.Fixed,
			},
			Size: 3,
			Data: []byte{0x42, 0x9b, 0x52},
		},
		&Extended{
			Base: Base{
				FRN:         3,
				DataItemName:    "I048/020",
				Description: "Target Report Descriptor",
				Type:        _uap.Extended,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 1,
			Primary:           []byte{0xa0},
			Secondary:         nil,
		},
		&Fixed{
			Base: Base{
				FRN:         4,
				DataItemName:    "I048/040",
				Description: "Measured Position in Slant Polar Coordinates",
				Type:        _uap.Fixed,
			},
			Size: 4,
			Data: []byte{0x94, 0xc7, 0x01, 0x81},
		},
		&Fixed{
			Base: Base{
				FRN:         5,
				DataItemName:    "I048/070",
				Description: "Mode-3/A Code in Octal Representation",
				Type:        _uap.Fixed,
			},
			Size: 2,
			Data: []byte{0x09, 0x13},
		},
		&Fixed{
			Base: Base{
				FRN:         6,
				DataItemName:    "I048/090",
				Description: "Flight Level in Binary Representation",
				Type:        _uap.Fixed,
			},
			Size: 2,
			Data: []byte{0x02, 0xd0},
		},
		&Compound{
			Base: Base{
				FRN:         7,
				DataItemName:    "I048/130",
				Description: "Radar Plot Characteristics",
				Type:        _uap.Compound,
			},
			Fields:  _uap.Cat048V127.DataItems[6].Compound,
			Primary: []byte{0x60},
			Secondary: []Item{
				&Fixed{
					Base: Base{
						FRN:         2,
						DataItemName:    "SRR",
						Description: "Number of received replies",
						Type:        _uap.Fixed,
					},
					Size: 1,
					Data: []byte{0x02},
				},
				&Fixed{
					Base: Base{
						FRN:         3,
						DataItemName:    "SAM",
						Description: "Amplitude of received replies for M(SSR)",
						Type:        _uap.Fixed,
					},
					Size: 1,
					Data: []byte{0xb7},
				},
			},
		},
		&Fixed{
			Base: Base{
				FRN:         8,
				DataItemName:    "I048/220",
				Description: "Aircraft Address",
				Type:        _uap.Fixed,
			},
			Size: 3,
			Data: []byte{0x49, 0x0d, 0x01},
		},
		&Fixed{
			Base: Base{
				FRN:         9,
				DataItemName:    "I048/240",
				Description: "Aircraft Identification",
				Type:        _uap.Fixed,
			},
			Size: 6,
			Data: []byte{0x38, 0xa1, 0x78, 0xcf, 0x42, 0x20},
		},
		&Repetitive{
			Base: Base{
				FRN:         10,
				DataItemName:    "I048/250",
				Description: "Mode S MB Data",
				Type:        _uap.Repetitive,
			},
			SubItemSize: 8,
			Rep:         0x02,
			Data:        []byte{0xe7, 0x9a, 0x5d, 0x27, 0xa0, 0x0c, 0x00, 0x60, 0xa3, 0x28, 0x00, 0x30, 0xa4, 0x00, 0x00, 0x40},
		},
		&Fixed{
			Base: Base{
				FRN:         11,
				DataItemName:    "I048/161",
				Description: "Track Number",
				Type:        _uap.Fixed,
			},
			Size: 2,
			Data: []byte{0x06, 0x3a},
		},
		&Fixed{
			Base: Base{
				FRN:         13,
				DataItemName:    "I048/200",
				Description: "Calculated Track Velocity in Polar Representation",
				Type:        _uap.Fixed,
			},
			Size: 4,
			Data: []byte{0x07, 0x43, 0xce, 0x5b},
		},
		&Extended{
			Base: Base{
				FRN:         14,
				DataItemName:    "I048/170",
				Description: "Track Status",
				Type:        _uap.Extended,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 1,
			Primary:           []byte{0x40},
			Secondary:         nil,
		},
		&Fixed{
			Base: Base{
				FRN:         21,
				DataItemName:    "I048/230",
				Description: "Communications / ACAS Capability and Flight Status",
				Type:        _uap.Fixed,
			},
			Size: 2,
			Data: []byte{0x20, 0xf5},
		},
	}

	uap048 := _uap.Cat048V127
	data, _ := util.HexStringToByte(input)
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
	for i, item := range rec.DataItems {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("FAIL: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", item, output[i])
		}
	}
}
*/
/*
func TestRecordDecode_CAT034(t *testing.T) {
	// Arrange
	input := "f6 0836 02 429b61 08 9400282000 94008000"
	output := []Item{
		{
			Meta: Base{
				FRN:         1,
				DataItemName:    "I034/010",
				Description: "Data Source Identifier",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x08, 0x36}},
		},
		{
			Meta: Base{
				FRN:         2,
				DataItemName:    "I034/000",
				Description: "Message Type",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x02}},
		},
		{
			Meta: Base{
				FRN:         3,
				DataItemName:    "I034/030",
				Description: "Time-of-Day",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x42, 0x9b, 0x61}},
		},
		{
			Meta: Base{
				FRN:         4,
				DataItemName:    "I034/020",
				Description: "Sector Number",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x08}},
		},
		{
			Meta: Base{
				FRN:         6,
				DataItemName:    "I034/050",
				Description: "System Configuration and Status",
				Type:        dataField.Compound,
			},
			Compound: &Compound{
				Primary: []byte{0x94},
				Secondary: []Item{
					{
						Meta: Base{
							FRN:         1,
							DataItemName:    "COM",
							Description: "Common Part",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x00}},
					},
					{
						Meta: Base{
							FRN:         4,
							DataItemName:    "PSR",
							Description: "Specific Status for PSR Sensor",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x28}},
					},
					{
						Meta: Base{
							FRN:         6,
							DataItemName:    "MDS",
							Description: "Specific Status for Mode S Sensor",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x20, 0x00}},
					},
				},
			},
		},
		{
			Meta: Base{
				FRN:         7,
				DataItemName:    "I034/060",
				Description: "System Processing Mode",
				Type:        dataField.Compound,
			},
			Compound: &Compound{
				Primary: []byte{0x94}, //1001-0100 94 00 80 00
				Secondary: []Item{
					{
						Meta: Base{
							FRN:         1,
							DataItemName:    "COM",
							Description: "Common Part",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x00}},
					},
					{
						Meta: Base{
							FRN:         4,
							DataItemName:    "PSR",
							Description: "Specific Processing Mode information for PSR Sensor",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x80}},
					},
					{
						Meta: Base{
							FRN:         6,
							DataItemName:    "MDS",
							Description: "Specific Processing Mode information for Mode S Sensor",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x00}},
					},
				},
			},
		},
	}

	uap034 := dataField.Cat034V127
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap034)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("MsgFailInValue: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %v; Expected: %v", item, output[i])
		}
	}
}

func TestRecordDecode_CAT063(t *testing.T) {
	// Arrange
	input := "bff0090c7cd2cc08294000000000000000000000000000000000"
	output := []Item{
		{
			Meta: Base{
				FRN:         1,
				DataItemName:    "I063/010",
				Description: "Data Source Identifier",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x09, 0x0c}},
		},
		{
			Meta: Base{
				FRN:         3,
				DataItemName:    "I063/030",
				Description: "Time of Message",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x7c, 0xd2, 0xcc}},
		},
		{
			Meta: Base{
				FRN:         4,
				DataItemName:    "I063/050",
				Description: "Sensor Identifier",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x08, 0x29}},
		},
		{
			Meta: Base{
				FRN:         5,
				DataItemName:    "I063/060",
				Description: "Sensor Configuration and Status",
				Type:        dataField.Extended,
			},
			Extended: &Extended{
				Primary:   []byte{0x40},
				Secondary: nil,
			},
		},
		{
			Meta: Base{
				FRN:         6,
				DataItemName:    "I063/070",
				Description: "Time Stamping Bias",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00}},
		},
		{
			Meta: Base{
				FRN:         7,
				DataItemName:    "I063/080",
				Description: "SSR/Mode S Range Gain and Bias",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00, 0x00, 0x00}},
		},
		{
			Meta: Base{
				FRN:         8,
				DataItemName:    "I063/081",
				Description: "SSR/Mode S Azimuth Bias",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00}},
		},
		{
			Meta: Base{
				FRN:         9,
				DataItemName:    "I063/090",
				Description: "PSR Range Gain and Bias",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00, 0x00, 0x00}},
		},
		{
			Meta: Base{
				FRN:         10,
				DataItemName:    "I063/091",
				Description: "PSR Azimuth Bias",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00}},
		},
		{
			Meta: Base{
				FRN:         11,
				DataItemName:    "I063/092",
				Description: "PSR Elevation Bias",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00}},
		},
	}

	uap063 := dataField.Cat063V16
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap063)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("MsgFailInValue: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %v; Expected: %v", item, output[i])
		}
	}
}

func TestRecordDecode_CAT065(t *testing.T) {
	// Arrange
	input := "f8090c0203424cf30a"
	output := []Item{
		{
			Meta: Base{
				FRN:         1,
				DataItemName:    "I065/010",
				Description: "Data Source Identifier",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x09, 0x0c}},
		},
		{
			Meta: Base{
				FRN:         2,
				DataItemName:    "I065/000",
				Description: "Message Type",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x02}},
		},
		{
			Meta: Base{
				FRN:         3,
				DataItemName:    "I065/015",
				Description: "Service Identification",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x03}},
		},
		{
			Meta: Base{
				FRN:         4,
				DataItemName:    "I065/030",
				Description: "Time Of Message",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x42, 0x4c, 0xf3}},
		},
		{
			Meta: Base{
				FRN:         5,
				DataItemName:    "I065/020",
				Description: "Batch Number",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x0a}},
		},
	}
	uap065 := dataField.Cat065V15
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap065)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("MsgFailInValue: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %v; Expected: %v", item, output[i])
		}
	}
}

func TestRecordDecode_CAT004(t *testing.T) {
	// Arrange
	input := "fdcb80 08a2 08 010882 6ae180 0000 08 0001 d1c0 41504d30303031 0001 0bc51ef7a55900f5 050370c30c40 00003039 ff50 ffd8a8 80 404cb3820820"
	output := []Item{
		{
			Meta: Base{
				FRN:         1,
				DataItemName:    "I004/010",
				Description: "Data Source Identifier",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x08, 0xa2}},
		},
		{
			Meta: Base{
				FRN:         2,
				DataItemName:    "I004/000",
				Description: "Message Type",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x08}},
		},
		{
			Meta: Base{
				FRN:         3,
				DataItemName:    "I004/015",
				Description: "SDPS Identifier",
				Type:        dataField.Repetitive,
			},
			Repetitive: &Repetitive{
				Rep:  0x01,
				Data: []byte{0x08, 0x82},
			},
		},
		{
			Meta: Base{
				FRN:         4,
				DataItemName:    "I004/020",
				Description: "Time Of Message",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x6a, 0xe1, 0x80}},
		},
		{
			Meta: Base{
				FRN:         5,
				DataItemName:    "I004/040",
				Description: "Alert Identifier",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x00}},
		},
		{
			Meta: Base{
				FRN:         6,
				DataItemName:    "I004/045",
				Description: "Alert Status",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x08}},
		},
		{
			Meta: Base{
				FRN:         8,
				DataItemName:    "I004/030",
				Description: "Track Number 1",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0x00, 0x01}},
		},
		{
			Meta: Base{
				FRN:         9,
				DataItemName:    "I004/170",
				Description: "Aircraft Identification & Characteristics 1",
				Type:        dataField.Compound,
			},
			Compound: &Compound{
				Primary: []byte{0xd1, 0xc0}, //1101-0001 1100-0000
				Secondary: []Item{
					{
						Meta: Base{
							FRN:         1,
							DataItemName:    "AI1",
							Description: "Aircraft Identifier 1",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x41, 0x50, 0x4d, 0x30, 0x30, 0x30, 0x31}},
					},
					{
						Meta: Base{
							FRN:         2,
							DataItemName:    "M31",
							Description: "Mode 3/A Code Aircraft 1",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x00, 0x01}},
					},
					{
						Meta: Base{
							FRN:         4,
							DataItemName:    "CPC",
							Description: "Predicted Conflict Position 1 (Cartesian Coordinates)",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x0b, 0xc5, 0x1e, 0xf7, 0xa5, 0x59, 0x00, 0xf5}},
					},
					{
						Meta: Base{
							FRN:         8,
							DataItemName:    "MS1",
							Description: "Mode S Identifier Aircraft 1",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x05, 0x03, 0x70, 0xc3, 0x0c, 0x40}},
					},
					{
						Meta: Base{
							FRN:         9,
							DataItemName:    "FP1",
							Description: "Flight Plan Number Aircraft 1",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x00, 0x00, 0x30, 0x39}},
					},
				},
			},
		},
		{
			Meta: Base{
				FRN:         12,
				DataItemName:    "I004/076",
				Description: "Vertical Deviation",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff, 0x50}},
		},
		{
			Meta: Base{
				FRN:         14,
				DataItemName:    "I004/075",
				Description: "Transversal Distance Deviation",
				Type:        dataField.Fixed,
			},
			Fixed: &Fixed{Data: []byte{0xff, 0xd8, 0xa8}},
		},
		{
			Meta: Base{
				FRN:         15,
				DataItemName:    "I004/100",
				Description: "Area Definitions",
				Type:        dataField.Compound,
			},
			Compound: &Compound{
				Primary: []byte{0x80},
				Secondary: []Item{
					{
						Meta: Base{
							FRN:         1,
							DataItemName:    "AN",
							Description: "Area Name",
							Type:        dataField.Fixed,
						},
						Fixed: &Fixed{Data: []byte{0x40, 0x4c, 0xb3, 0x82, 0x08, 0x20}},
					},
				},
			},
		},
	}

	uap004 := dataField.Cat004V112
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap004)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if reflect.DeepEqual(item, output[i]) == false {
			t.Errorf("MsgFailInValue: %v; \nExpected: %v", item, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %v; Expected: %v", item, output[i])
		}
	}
}
*/
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

	uap001 := dataField.Cat001V12
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap001)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("MsgFailInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
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

	uap001 := dataField.Cat001V12
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap001)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("MsgFailInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
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

	uap002 := dataField.Cat002V10
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap002)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("MsgFailInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
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
	uap030 := dataField.Cat030StrV51
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap030)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("MsgFailInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
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

	uap030 := dataField.Cat032StrV70
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap030)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("MsgFailInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
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

	uap062 := dataField.Cat062V119
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap062)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("MsgFailInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
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

	uap255 := dataField.Cat255StrV51
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap255)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("MsgFailInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
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

	uap030 := dataField.Cat030ArtasV62
	data, _ := util.HexStringToByte(input)
	rec := new(Record)

	// Act
	unRead, err := rec.Decode(data, uap030)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}
	if unRead != 0 {
		t.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
	} else {
		t.Logf("MsgSuccessInValue: unRead = %v; Expected: %v", unRead, 0)
	}
	for i, item := range rec.DataItems {
		if bytes.Equal(item.Data, output[i]) == false {
			t.Errorf("MsgFailInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
		} else {
			t.Logf("MsgSuccessInValue: %s = % X; Expected: % X", item.DataItemName, item.Data, output[i])
		}
	}
}
*/

package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"io"
	"testing"
)

func Test_Record_Payload(t *testing.T) {
	// Arrange
	data := StringToHex("ff df 02 93 19 37 8d 3d a2 05 6f 13 2d 0f ff 00 94 60 02 de 50 6f 84 4c c3 c3 51 23 31 00 17 01 3b 02 6c 00 0c 74 a7 40 20 a0")
	nbOfBytes := 42
	rec := new(Record)
	_, _ = rec.Decode(data, uap.Cat048V127.Items)

	// Act
	items := rec.Payload()

	// Assert
	if len(items) != nbOfBytes {
		t.Errorf("FAIL: len(items) = %v; Expected: %v", len(items), nbOfBytes)
	} else {
		t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(items), nbOfBytes)
	}
}

func Test_Record_String(t *testing.T) {
	// Arrange
	data := StringToHex("ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020a0")
	nbOfItems := 15
	rec := new(Record)
	_, _ = rec.Decode(data, uap.Cat048V127.Items)

	// Act
	items := rec.String()

	// Assert
	if len(items) != nbOfItems {
		t.Errorf("FAIL: len(items) = %v; Expected: %v", len(items), nbOfItems)
	} else {
		t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(items), nbOfItems)
	}

}

func Test_FspecReader_valid(t *testing.T) {
	// Arrange
	input := []byte{0xFF, 0x01, 0xF2, 0xFF}
	output := []byte{0xFF, 0x01, 0xF2}
	rb := bytes.NewReader(input)

	// Act
	fspec, err := FspecReader(rb, 1)

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

func Test_FspecReader_invalid(t *testing.T) {
	// Arrange
	input := []byte{0xFF, 0x01}
	var output []byte
	rb := bytes.NewReader(input)

	// Act
	fspec, err := FspecReader(rb, 1)

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

func Test_DataFieldSPAndREReader_valid(t *testing.T) {
	// Arrange
	input := StringToHex("03 FF FF")
	output := []byte{0x03, 0xFF, 0xFF}
	rb := bytes.NewReader(input)

	// Act
	sp, err := SPAndREDataFieldReader(rb)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if bytes.Equal(sp, output) == false {
		t.Errorf("FAIL: sp = % X; Expected: % X", sp, output)
	} else {
		t.Logf("SUCCESS: sp = % X; Expected: % X", sp, output)
	}
}

func Test_DataFieldSPAndREReader_invalid(t *testing.T) {
	// Arrange
	input := StringToHex("03 FF")
	var output []byte
	rb := bytes.NewReader(input)

	// Act
	sp, err := SPAndREDataFieldReader(rb)

	// Assert
	if err != io.ErrUnexpectedEOF {
		t.Errorf("FAIL: error: %s; Expected: %v", err, io.ErrUnexpectedEOF)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, io.ErrUnexpectedEOF)
	}

	if bytes.Equal(sp, output) == false {
		t.Errorf("FAIL: sp = % X; Expected: % X", sp, output)
	} else {
		t.Logf("SUCCESS: sp = % X; Expected: % X", sp, output)
	}
}

func Test_DataFieldRepetitiveReader_valid(t *testing.T) {
	// Arrange
	input := StringToHex("03 01 02 03 04 05 06 07 08 09")
	nb := uint8(3)
	rb := bytes.NewReader(input)
	output := []byte{0x03, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}

	// Act
	item, err := RepetitiveDataFieldReader(rb, nb)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if bytes.Equal(item, output) == false {
		t.Errorf("FAIL: item = % X; Expected: % X", item, output)
	} else {
		t.Logf("SUCCESS: item = % X; Expected: % X", item, output)
	}
}

func Test_DataFieldRepetitiveReader_invalid(t *testing.T) {
	// Arrange
	input := StringToHex("04 01 02 03 04 05 06 07 08 09")
	nb := uint8(3)
	rb := bytes.NewReader(input)

	// Act
	item, err := RepetitiveDataFieldReader(rb, nb)

	// Assert
	if err != io.ErrUnexpectedEOF {
		t.Errorf("FAIL: error: %s; Expected: %v", err, io.ErrUnexpectedEOF)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, io.ErrUnexpectedEOF)
	}
	if item != nil {
		t.Errorf("FAIL: item = %v; Expected: %v", item, nil)
	} else {
		t.Logf("SUCCESS: item = %v; Expected: %v", item, nil)
	}
}

func Test_DataFieldFixedReader_valid(t *testing.T) {
	// Arrange
	input := StringToHex("FF FE FD BF 00 01 02 03")
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
	if bytes.Equal(item, output) == false {
		t.Errorf("FAIL: sp = % X; Expected: % X", item, output)
	} else {
		t.Logf("SUCCESS: sp = % X; Expected: % X", item, output)
	}
}

func Test_DataFieldFixedReader_invalid(t *testing.T) {
	// Arrange
	input := StringToHex("FF FE BF 00 01 02")
	nb := uint8(7)
	rb := bytes.NewReader(input)

	// Act
	item, err := FixedDataFieldReader(rb, nb)

	// Assert
	if err != io.EOF {
		t.Errorf("FAIL: error: %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if item != nil {
		t.Errorf("FAIL: item = %v; Expected: %v", item, nil)
	} else {
		t.Logf("SUCCESS: item = %v; Expected: %v", item, nil)
	}
}

func Test_DataFieldCompoundReader(t *testing.T) {
	// Arrange
	// todo: make a tab for different cat
	input := StringToHex("94 00 80 00")
	output := []byte{0x94, 0x00, 0x80, 0x00}
	item034060 := uap.MetaField{
		8: {Name: "Fixed", Size: 1},
		7: {Name: "Spare"},
		6: {Name: "Spare"},
		5: {Name: "Fixed", Size: 1},
		4: {Name: "Fixed", Size: 1},
		3: {Name: "Fixed", Size: 1},
		2: {Name: "Spare"},
	}

	rb := bytes.NewReader(input)

	// Act
	item, err := CompoundDataFieldReader(rb, item034060)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if bytes.Equal(item, output) == false {
		t.Errorf("FAIL: sp = % X; Expected: % X", item, output)
	} else {
		t.Logf("SUCCESS: sp = % X; Expected: % X", item, output)
	}

}

func Test_DataFieldRFSReader_valid(t *testing.T) {
	// Arrange
	// N = 2, FRN = 3, FRN = 17
	input := StringToHex("02 03 FFFF 11 FFFFFFFF")
	uap001 := uap.Cat001TrackV12.Items
	rb := bytes.NewReader(input)
	output := []byte{0x02, 0x03, 0xFF, 0xFF, 0x11, 0xFF, 0xFF, 0xFF, 0xFF}

	// Act
	item, err := RFSDataFieldReader(rb, uap001)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if bytes.Equal(item, output) == false {
		t.Errorf("FAIL: rfs = % X; Expected: % X", item, output)
	} else {
		t.Logf("SUCCESS: rfs = % X; Expected: % X", item, output)
	}
}

func Test_DataFieldExtendedReader_valid(t *testing.T) {
	// Arrange
	input := StringToHex("01 03 07 09 0B 0D 0F 0E")
	rb := bytes.NewReader(input)
	output := []byte{0x01, 0x03, 0x07, 0x09, 0x0B, 0x0D, 0x0F, 0x0E}

	// Act
	item, err := ExtendedDataFieldReader(rb, 1)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if bytes.Equal(item, output) == false {
		t.Errorf("FAIL: item = % X; Expected: % X", item, output)
	} else {
		t.Logf("SUCCESS: item = % X; Expected: % X", item, output)
	}
}

func Test_DataFieldExtendedReader_invalid(t *testing.T) {
	// Arrange
	input := StringToHex("")
	rb := bytes.NewReader(input)

	// Act
	item, err := ExtendedDataFieldReader(rb, 1)

	// Assert
	if err != io.EOF {
		t.Errorf("FAIL: error: %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if item != nil {
		t.Errorf("FAIL: item = %v; Expected: %v", item, nil)
	} else {
		t.Logf("SUCCESS: item = %v; Expected: %v", item, nil)
	}
}

func Test_DataFieldExtendedReader_valid_size3(t *testing.T) {
	// Arrange
	input := StringToHex("FFFFFE")
	rb := bytes.NewReader(input)
	output := []byte{0xFF, 0xFF, 0xFE}

	// Act
	item, err := ExtendedDataFieldReader(rb, 3)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if bytes.Equal(item, output) == false {
		t.Errorf("FAIL: item = % X; Expected: % X", item, output)
	} else {
		t.Logf("SUCCESS: item = % X; Expected: % X", item, output)
	}
}

/**
Testing by record
*/
type DataRecordTest struct {
	input     string          // data test one record = fspec + items
	uap       []uap.DataField // Items of category corresponding to data test input
	nbOfItems int
	err       error // error expected
}

func Test_Record_Decode_nbOfItems(t *testing.T) {
	// setup
	dataSetRecordTests := []DataRecordTest{
		{
			input:     "f6083602429b7110940028200094008000",
			uap:       uap.Cat034V127.Items,
			err:       nil,
			nbOfItems: 6,
		},
		{
			input:     "f6083602429b71109400282000940080",
			uap:       uap.Cat034V127.Items,
			err:       io.EOF,
			nbOfItems: 5,
		},
		{
			input:     "ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020a0",
			uap:       uap.Cat048V127.Items,
			err:       nil,
			nbOfItems: 14,
		},
		{
			// 0xA0 last byte is removed
			input:     "ffdf029319378d3da2056f132d0fff00946002de506f844cc3c35123310017013b026c000c74a74020",
			uap:       uap.Cat048V127.Items,
			err:       io.EOF,
			nbOfItems: 13,
		},
		{
			input:     "f0 0831 00 0a8abb2e 3802",
			uap:       uap.CatT001PlotV12.Items,
			err:       nil,
			nbOfItems: 4,
		},
		{
			input:     "f0 0831 00 0a8abb2e 38",
			uap:       uap.CatT001PlotV12.Items,
			err:       io.EOF,
			nbOfItems: 3,
		},
		{
			input:     "f502 0831 98 01bf 0a1ebb43 022538e2 00",
			uap:       uap.Cat001TrackV12.Items,
			err:       nil,
			nbOfItems: 6,
		},
		{
			input:     "f502 0831 98 01bf 0a1ebb43 022538e2",
			uap:       uap.Cat001TrackV12.Items,
			err:       io.EOF,
			nbOfItems: 5,
		},
	}

	for _, row := range dataSetRecordTests {
		// Arrange
		data := StringToHex(row.input)
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

func Test_Record_Decode_CAT048_record(t *testing.T) {
	// Arrange
	input := "fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5"
	output := [][]byte{
		{0x08, 0x36},
		{0x42, 0x9b, 0x52},
		{0xa0},
		{0x94, 0xc7, 0x01, 0x81},
		{0x09, 0x13},
		{0x02, 0xd0},
		{0x60, 0x02, 0xb7},
		{0x49, 0x0d, 0x01},
		{0x38, 0xa1, 0x78, 0xcf, 0x42, 0x20},
		{0x02, 0xe7, 0x9a, 0x5d, 0x27, 0xa0, 0x0c, 0x00, 0x60, 0xa3, 0x28, 0x00, 0x30, 0xa4, 0x00, 0x00, 0x40},
		{0x06, 0x3a},
		{0x07, 0x43, 0xce, 0x5b},
		{0x40},
		{0x20, 0xf5},
	}
	uap048 := uap.Cat048V127.Items
	data := StringToHex(input)
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
		if bytes.Equal(item.Payload, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		}
	}
}

func Test_Record_Decode_CAT001_Track_record(t *testing.T) {
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

	uap001 := uap.Cat001TrackV12.Items
	data := StringToHex(input)
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
		if bytes.Equal(item.Payload, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		}
	}
}

func Test_Record_Decode_CAT001_Plot_record(t *testing.T) {
	// Arrange
	//input := "F0 08 31 08 0A 8A BB 2E 38 02"
	input := "f0 0831 00 0a8abb2e 3802"
	output := [][]byte{
		{0x08, 0x31},
		{0x00},
		{0x0a, 0x8a, 0xbb, 0x2e},
		{0x38, 0x02},
	}

	uap001 := uap.CatT001PlotV12.Items
	data := StringToHex(input)
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
		if bytes.Equal(item.Payload, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		}
	}
}

func Test_Record_Decode_CAT002_record(t *testing.T) {
	// Arrange
	input := "f4 0839 02105fb35b02"
	output := [][]byte{
		{0x08, 0x39},
		{0x02},
		{0x10},
		{0x5f, 0xb3, 0x5b},
		{0x02},
	}

	uap002 := uap.Cat002V10.Items
	data := StringToHex(input)
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
		if bytes.Equal(item.Payload, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		}
	}
}

func Test_Record_Decode_CAT030_STR_record(t *testing.T) {
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
	uap030 := uap.Cat030StrV51.Items
	data := StringToHex(input)
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
		if bytes.Equal(item.Payload, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		}
	}
}

func Test_Record_Decode_CAT032_STR_record_valid(t *testing.T) {
	// Arrange
	input := "d0 0884 3b5494 00130000008f002f008948006a007c"
	output := [][]byte{
		{0x08, 0x84},
		{0x3b, 0x54, 0x94},
		{0x00, 0x13, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x2f, 0x00, 0x89, 0x48, 0x00, 0x6a, 0x00, 0x7c},
	}

	uap030 := uap.Cat032StrV70.Items
	data := StringToHex(input)
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
		if bytes.Equal(item.Payload, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		}
	}
}

func Test_Record_Decode_CAT034_record(t *testing.T) {
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
	uap048 := uap.Cat034V127.Items
	data := StringToHex(input)
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
		if bytes.Equal(item.Payload, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		}
	}
}

func Test_Record_Decode_CAT255_STR_record(t *testing.T) {
	// Arrange
	input := "e0 08 83 7dfd9c 58"
	output := [][]byte{
		{0x08, 0x83},
		{0x7d, 0xfd, 0x9c},
		{0x58},
	}

	uap255 := uap.Cat255StrV51.Items
	data := StringToHex(input)
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
		if bytes.Equal(item.Payload, output[i]) == false {
			t.Errorf("FAIL: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		} else {
			t.Logf("SUCCESS: %s = % X; Expected: % X", item.DataItem, item.Payload, output[i])
		}
	}
}

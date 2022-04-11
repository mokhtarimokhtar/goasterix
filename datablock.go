// Package goasterix parses ASTERIX binary data Format.
package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"math"
	"strings"
)

var (
	// ErrUndersized reports that the data byte source is too small.
	ErrUndersized = errors.New("[ASTERIX] undersized packet")

	// ErrCategoryUnknown reports which Category Unknown or not processed.
	ErrCategoryUnknown = errors.New("[ASTERIX] category unknown or not processed")
)

// WrapperDataBlock
// a Wrapper DataBlock correspond to one or more category and contains one or more Records.
// DataBlock = CAT + LEN + [FSPEC + items...] + [...] + ...
// WrapperDataBlock = [CAT + LEN + RECORD + ...] + [ DATABLOCK ] + [...]
type WrapperDataBlock struct {
	DataBlocks []*DataBlock
}

func NewWrapperDataBlock() (*WrapperDataBlock, error) {
	w := &WrapperDataBlock{}
	return w, nil
}

func (w *WrapperDataBlock) Decode(data []byte) (unRead int, err error) {
	offset := uint16(0)
	for {
		db := NewDataBlock()
		unRead, err := db.Decode(data[offset:])
		offset += db.Len
		if err != nil {
			return unRead, err
		}

		w.DataBlocks = append(w.DataBlocks, db)
		if unRead == 0 {
			break
		}
	}
	return unRead, err
}

// DataBlock
// a DataBlock corresponds to one (only) category and contains one or more Records.
// DataBlock = CAT + LEN + [FSPEC + items...] + [...] + ...
type DataBlock struct {
	Category uint8
	Len      uint16
	Records  []*Record
}

func NewDataBlock() *DataBlock {
	return &DataBlock{}
}

// Decode extracts an asterix data block: CAT + LEN + N * RECORD(S).
// An asterix data block can contain a or more records.
// It returns the number of bytes unRead and fills the DataBlock Struct(Category, Len, Records array) in byte.
func (db *DataBlock) Decode(data []byte) (int, error) {
	var unRead int
	var err error
	rb := bytes.NewReader(data)

	// retrieve category field
	err = binary.Read(rb, binary.BigEndian, &db.Category)
	if err != nil {
		unRead = rb.Len()
		return unRead, err // err = io.EOF
	}

	// retrieve length field
	err = binary.Read(rb, binary.BigEndian, &db.Len)
	if err != nil {
		unRead = rb.Len()
		return unRead, err
	}
	// check if the rest is big enough
	rbSize := uint16(rb.Size())
	if rbSize < db.Len {
		db.Records = nil
		err = ErrUndersized
		unRead = rb.Len()
		return unRead, err
	}

	// retrieve records
	tmp := make([]byte, db.Len-3)
	_ = binary.Read(rb, binary.BigEndian, tmp) // ignore explicit error because it's detect before
	unRead = rb.Len()

	// decode N * records
	offset := 0
	lenData := len(tmp)

	// selection of the appropriate UAP
	uapSelected, found := uap.DefaultProfiles[db.Category]
	if !found {
		err = ErrCategoryUnknown
		return unRead, err
	}

LoopRecords:
	for {
		rec := NewRecord()
		unRead, err := rec.Decode(tmp[offset:], uapSelected)
		db.Records = append(db.Records, rec)
		offset = lenData - unRead

		if err != nil {
			return unRead, err
		}
		// offset == lenData is for the case payload is oversize of LEN field asterix
		// if unRead == 0 || offset == lenData {
		if unRead == 0 {
			break LoopRecords
		}
	}
	return unRead, nil
}

func (db DataBlock) String() [][]string {
	var records [][]string
	for _, record := range db.Records {
		records = append(records, record.String())
	}
	return records
}

func (db DataBlock) Payload() [][]byte {
	var pd [][]byte
	var catPd []byte
	var lenPd []byte

	catPd = append(catPd, db.Category)
	pd = append(pd, catPd)

	var h, l = byte(db.Len >> 8), byte(db.Len & 0xff)
	lenPd = append(lenPd, h, l)
	pd = append(pd, lenPd)

	for _, record := range db.Records {
		pd = append(pd, record.Payload())
	}
	return pd
}

// TwoComplement16 returns an int16 (signed).
// sizeBits is the number of bit complement.
func TwoComplement16(sizeBits uint8, data uint16) (v int16) {
	n := float64(sizeBits - 1)
	p := math.Pow(2, n) // 2^(N-1)
	mask := uint16(p)

	tmp1 := -int16(data & mask)
	tmp2 := int16(data & ^mask)

	v = tmp1 + tmp2
	return v
}

// TwoComplement32 returns an int32 (signed).
// sizeBits is the number of bit complement.
func TwoComplement32(sizeBits uint8, data uint32) (v int32) {
	n := float64(sizeBits - 1)
	p := math.Pow(2, n) // 2^(N-1)
	mask := uint32(p)

	tmp1 := -int32(data & mask)
	tmp2 := int32(data & ^mask)

	v = tmp1 + tmp2

	// checking example
	// mask := uint32(0x007F)
	/*signed := data[2] & 0x80 >> 7
	if signed == 1 {
		complement := ^tmpLatitude 	// one complement
		v := complement & mask	// apply mask 2^(N-1)
		v = v + 1
		latitude = -float64(v) * 0.000021458
	} else {
		latitude = float64(tmpLatitude) * 0.000021458
	}*/

	return v
}

// HexStringToByte converts a hexadecimal string format to an array of byte.
// It is used to facilitate the testing.
func HexStringToByte(s string) ([]byte, error) {
	s = strings.ReplaceAll(s, " ", "")
	data, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CleanStringMultiline cleans the escapes characters.
// It is used to facilitate the testing.
func CleanStringMultiline(s string) string {
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, " ", "")
	return s
}

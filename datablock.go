// Package goasterix parses ASTERIX binary data Format.
package goasterix

import (
	//"bytes"
	//"encoding/binary"
	//"errors"
	//"github.com/mokhtarimokhtar/goasterix/_uap"
	"bytes"
	"encoding/binary"
	"errors"
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
	DataBlocks []DataBlock
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

		w.DataBlocks = append(w.DataBlocks, *db)
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
	Records  []Record
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

	// retrieve category dataField
	db.Category, err = rb.ReadByte()
	if err != nil {
		unRead = rb.Len()
		return unRead, err // err = io.EOF
	}

	// retrieve length dataField
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
	uapSelected, found := ProfileRegistry[db.Category]
	if !found {
		err = ErrCategoryUnknown
		return unRead, err
	}

LoopRecords:
	for {
		rec := NewRecord()
		unRead, err := rec.Decode(tmp[offset:], uapSelected)
		db.Records = append(db.Records, *rec)
		offset = lenData - unRead

		if err != nil {
			return unRead, err
		}
		// offset == lenData is for the case payload is oversize of LEN dataField asterix
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

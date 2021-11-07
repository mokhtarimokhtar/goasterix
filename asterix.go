// Package goasterix parses ASTERIX binary data Format,
// (All Purpose Structured EUROCONTROL Surveillance Information Exchange)
// For information about ASTERIX, see https://www.eurocontrol.int/asterix

package goasterix

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

const (
	track string = "track"
	plot  string = "plot"
)

var (
	// ErrUndersized reports that the data byte source is too small.
	ErrUndersized = errors.New("[ASTERIX Error] undersized packet")

	// ErrCategoryUnknown reports which Category Unknown or not processed.
	ErrCategoryUnknown = errors.New("[ASTERIX Error] Category Unknown or not processed")
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
		db, _ := NewDataBlock()
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
// a DataBlock correspond to one (only) category and contains one or more Records.
// DataBlock = CAT + LEN + [FSPEC + items...] + [...] + ...
type DataBlock struct {
	Category uint8
	Len      uint16
	Records  []*Record
}

func NewDataBlock() (*DataBlock, error) {
	db := &DataBlock{}
	return db, nil
}

func (db *DataBlock) Decode(data []byte) (unRead int, err error) {
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
	_ = binary.Read(rb, binary.BigEndian, tmp) // ignore explicity error because it's detect before
	unRead = rb.Len()

	// decode N * records
	offset := 0
	lenData := len(tmp)

	// retrieve its Items
	var uapSelected uap.StandardUAP

	switch db.Category {
	case 1:
		//uap.Profiles[1].Items
		uapSelected = uap.Cat001V12
	case 2:
		uapSelected = uap.Cat002V10
	case 30:
		uapSelected = uap.Cat030StrV51
	case 32:
		uapSelected = uap.Cat032StrV70
	case 34:
		uapSelected = uap.Cat034V127
	case 48:
		uapSelected = uap.Cat048V127
	case 255:
		uapSelected = uap.Cat255StrV51
	default:
		err = ErrCategoryUnknown
		return unRead, err
	}

LoopRecords:
	for {
		rec := new(Record)
		unRead, err := rec.Decode(tmp[offset:], uapSelected)
		db.Records = append(db.Records, rec)
		offset = lenData - unRead

		if err != nil {
			return unRead, err
		}
		if unRead == 0 {
			break LoopRecords
		}
	}
	return unRead, nil
}

func (db *DataBlock) String() (records [][]string) {
	for _, record := range db.Records {
		records = append(records, record.String())
	}
	return records
}

func (db *DataBlock) Payload() (b [][]byte) {
	for _, record := range db.Records {
		b = append(b, record.Payload())
	}
	return b
}

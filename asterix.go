// Copyright 2019 DSNA-DTI, Mokhtar Mokhtari. All rights reserved.

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

func (w *WrapperDataBlock) Decode(data []byte) (unRead int, err error) {
	offset := uint16(0)

	for {
		db := new(DataBlock)
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
	var uapSelected []uap.DataField

	switch db.Category {
	case 1:
		uapSelected = uap.CatT001PlotV12.Items
	case 2:
		uapSelected = uap.Cat002V10.Items
	case 30:
		uapSelected = uap.Cat030StrV51.Items
	case 32:
		uapSelected = uap.Cat032StrV70.Items
	case 34:
		uapSelected = uap.Cat034V127.Items
	case 48:
		uapSelected = uap.Cat048V127.Items
	case 255:
		uapSelected = uap.Cat255StrV51.Items
	default:
		err = ErrCategoryUnknown
		return unRead, err
	}

	if db.Category != 1 {
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
	} else {
	LoopRecords2:
		// special use case for CAT001
		// uap can change in different records
		for {
			rec := new(Record)
			typeTarget, _ := SelectUAPCat001(tmp[offset:])
			if typeTarget == track {
				unRead, err := rec.Decode(tmp[offset:], uap.Cat001TrackV12.Items)
				db.Records = append(db.Records, rec)
				offset = lenData - unRead

				if err != nil {
					return unRead, err
				}
				if unRead == 0 {
					break LoopRecords2
				}
			} else if typeTarget == plot {
				unRead, err := rec.Decode(tmp[offset:], uap.CatT001PlotV12.Items)
				db.Records = append(db.Records, rec)
				offset = lenData - unRead

				if err != nil {
					return unRead, err
				}
				if unRead == 0 {
					break LoopRecords2
				}
			}
		}
		return unRead, nil
	}
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

func SelectUAPCat001(data []byte) (uap string, err error) {
	rb := bytes.NewReader(data)

	fspec, err := FspecReader(rb, 1)
	if err != nil {
		return "", err
	}

	// Item010
	if fspec[0]&0x80 != 0 {
		tmp := make([]byte, 2)
		err = binary.Read(rb, binary.BigEndian, &tmp)
		if err != nil {
			return "", err
		}
	}

	// Item020
	if fspec[0]&0x40 != 0 {
		b := make([]byte, 1)
		err = binary.Read(rb, binary.BigEndian, &b)
		if err != nil {
			return "", err
		}
		// retrieve type of Items (plot or Track information)
		typeTarget := b[0] & 0x80 >> 7
		if typeTarget == 1 {
			uap = track
		} else {
			uap = plot
		}
	}

	return uap, nil
}

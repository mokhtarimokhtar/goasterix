package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"io"
	"math/bits"

	"github.com/mokhtarimokhtar/goasterix/uap"
)

var (
	// ErrDataFieldUnknown reports which ErrDatafield Unknown.
	ErrDataFieldUnknown = errors.New("type of datafield not found")
)

type Record struct {
	Cat   uint8
	Fspec []byte
	Items []Item
}

func NewRecord() *Record {
	return &Record{}
}

// Decode extracts a Record of asterix data block (only one record).
// An asterix data block can contain a or more records.
// It returns the number of bytes unread and fills the Record Struct(Fspec, Items array) in byte.
func (rec *Record) Decode(data []byte, stdUAP uap.StandardUAP) (unRead int, err error) {
	rec.Cat = stdUAP.Category

	rb := bytes.NewReader(data)
	rec.Fspec, err = FspecReader(rb)
	unRead = rb.Len()
	if err != nil {
		return unRead, err
	}

	frnIndex := FspecIndex(rec.Fspec)
	offset := uint8(0) // offset shifts the index for a conditional UAP

	for _, frn := range frnIndex {
		uapItem := stdUAP.Items[frn-1-offset] // here the index corresponds to the FRN
		var item Item
		item, err = GetItem(uapItem)
		if err != nil {
			unRead = rb.Len()
			return unRead, err
		}
		//err = Readers(item, rb, uapItem)
		err = Readers(item, rb)
		if err != nil {
			unRead = rb.Len()
			return unRead, err
		}
		/*
			contextType := new(ContextType)
			err = contextType.setReader(uapItem.Type)
			err = contextType.Reader(rb, uapItem)
		*/

		unRead = rb.Len()
		//rec.Items = append(rec.Items, contextType.Item)
		rec.Items = append(rec.Items, item)

		if uapItem.Conditional {
			switch uapItem.Type {
			case uap.Fixed:
				stdUAP.Items = selectUAPConditional(stdUAP.Category, item.Payload())
				//stdUAP.Items = selectUAPConditional(stdUAP.Category, contextType.Item.Payload())
			case uap.Extended:
				stdUAP.Items = selectUAPConditional(stdUAP.Category, item.Payload())
				//stdUAP.Items = selectUAPConditional(stdUAP.Category, contextType.Item.Payload())
			}
			offset = frn
		}
	}
	return unRead, nil
}

// String returns a string(hex) representation of one asterix record (only existing items).
func (rec Record) String() []string {
	var items []string
	tmp := "FSPEC: " + hex.EncodeToString(rec.Fspec)
	items = append(items, tmp)

	for _, item := range rec.Items {
		tmp := item.String()
		items = append(items, tmp)
	}
	return items
}

// Payload returns a slice of byte for one asterix record.
func (rec Record) Payload() []byte {
	var pd []byte
	pd = append(pd, rec.Fspec...)
	for _, item := range rec.Items {
		pd = append(pd, item.Payload()...)
	}
	return pd
}

func selectUAPConditional(category uint8, field []byte) []uap.DataField {
	var selectedUAP []uap.DataField
	switch category {
	case 1:
		tmp := field[0] & 0x80 >> 7
		if tmp == 1 {
			selectedUAP = uap.Cat001TrackV12
		} else {
			selectedUAP = uap.Cat001PlotV12
		}
	case 26:
		tmp := field[0] & 0x80 >> 7
		if tmp == 1 {
			selectedUAP = uap.Cat4TestTrack
		} else {
			selectedUAP = uap.Cat4TestPlot
		}
	}
	return selectedUAP
}

// FspecReader returns a slice of FSPEC data record asterix.
func FspecReader(reader io.Reader) ([]byte, error) {
	var fspec []byte
	var err error
	for {
		var tmp uint8
		err = binary.Read(reader, binary.BigEndian, &tmp)
		if err != nil {
			return nil, err
		}
		fspec = append(fspec, tmp)
		if tmp&0x01 == 0 {
			break
		}
	}
	return fspec, err
}

// FspecIndex returns an array of uint8 corresponding to number FRN(Field Reference Number of Items).
// In other words, it transposes a fspec bits to an array FRNs.
// e.g. fspec = 1010 1010 => frnIndex = []uint8{1, 3, 5, 7}
func FspecIndex(fspec []byte) []uint8 {
	var frnIndex []uint8
	for j, val := range fspec {
		for i := 0; i < 7; i++ {
			frn := 7*j + i + 1
			tmp := bits.RotateLeft8(val, i)
			if tmp&0x80 != 0 {
				frnIndex = append(frnIndex, uint8(frn))
			}
		}
	}
	return frnIndex
}

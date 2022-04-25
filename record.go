package goasterix

import (
	"bytes"
	"encoding/hex"
	"errors"
	"github.com/mokhtarimokhtar/goasterix/_uap"
)

var (
	// ErrDataFieldUnknown reports which ErrDatafield Unknown.
	ErrDataFieldUnknown = errors.New("type of datafield not found")
)

type StandardUAP struct {
	Name      string
	Category  uint8
	Version   float64
	DataItems []Item
}

type IRecord interface {
	GetItems() []Item
}

type Record struct {
	Cat   uint8
	Fspec []byte
	Items []Item
}

func (rec Record) GetItems() []Item {
	return rec.Items
}

func NewRecord() *Record {
	return &Record{}
}

// Decode extracts a Record of asterix data block (only one record).
// An asterix data block can contain a or more records.
// It returns the number of bytes unread and fills the Record Struct(Fspec, Items array) in byte.
//func (rec *Record) Decode(data []byte, stdUAP _uap.StandardUAP) (unRead int, err error) {
func (rec *Record) Decode(data []byte, stdUAP StandardUAP) (unRead int, err error) {
	rec.Cat = stdUAP.Category

	rb := bytes.NewReader(data)
	rec.Fspec, err = FspecReader(rb)
	unRead = rb.Len()
	if err != nil {
		return unRead, err
	}

	frnIndex := FspecIndex(rec.Fspec)
	offset := uint8(0) // offset shifts the index for a conditional UAP
	rec.Items = make([]Item, 0, len(frnIndex))

	for _, frn := range frnIndex {
		uapItem := stdUAP.DataItems[frn-1-offset] // here the index corresponds to the FRN
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
		//rec.DataItems = append(rec.DataItems, contextType.Item)
		rec.Items = append(rec.Items, item)
		/*
			if uapItem.Conditional {
				switch uapItem.Type {
				case _uap.Fixed:
					stdUAP.DataItems = selectUAPConditional(stdUAP.Category, item.Payload())
					//stdUAP.DataItems = selectUAPConditional(stdUAP.Category, contextType.Item.Payload())
				case _uap.Extended:
					stdUAP.DataItems = selectUAPConditional(stdUAP.Category, item.Payload())
					//stdUAP.DataItems = selectUAPConditional(stdUAP.Category, contextType.Item.Payload())
				}
				offset = frn
			}*/
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

func selectUAPConditional(category uint8, field []byte) []_uap.DataField {
	var selectedUAP []_uap.DataField
	switch category {
	case 1:
		tmp := field[0] & 0x80 >> 7
		if tmp == 1 {
			selectedUAP = _uap.Cat001TrackV12
		} else {
			selectedUAP = _uap.Cat001PlotV12
		}
	case 26:
		tmp := field[0] & 0x80 >> 7
		if tmp == 1 {
			selectedUAP = _uap.Cat4TestTrack
		} else {
			selectedUAP = _uap.Cat4TestPlot
		}
	}
	return selectedUAP
}

// FspecReader returns a slice of FSPEC data record asterix.
func FspecReader(rb *bytes.Reader) ([]byte, error) {
	var fspec []byte
	var err error
	var tmp byte
	for {
		tmp, err = rb.ReadByte()
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

// FspecIndex returns an array of uint8 corresponding to number FRN(Field Reference Number of DataItems).
// In other words, it transposes a fspec pos to an array FRNs.
// e.g. fspec = 1010 1010 => frnIndex = []uint8{1, 3, 5, 7}
func FspecIndex(fspec []byte) []uint8 {
	l := uint8(len(fspec))
	var frnIndex = make([]uint8, 0, l*7)
	var tmp byte
	for j := uint8(0); j < l; j++ {
		for i := uint8(0); i < 7; i++ {
			tmp = fspec[j] << i
			if tmp&0x80 != 0 {
				frnIndex = append(frnIndex, 7*j+i+1)
			}
		}
	}
	return frnIndex
}

package goasterix

import (
	"bytes"
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/item"
)

/*var (
	// ErrDataFieldUnknown reports which ErrDatafield Unknown.
	ErrDataFieldUnknown = errors.New("type of datafield not found")
)*/

/*type IRecord interface {
	GetItems() []DataItemName
}*/

type Record struct {
	Cat       uint8
	Fspec     []byte
	DataItems []item.DataItem
}

func (rec Record) GetItems() []item.DataItem {
	return rec.DataItems
}

func NewRecord() *Record {
	return &Record{}
}

// Decode extracts a Record of asterix data block (only one record).
// An asterix data block can contain a or more records.
// It returns the number of bytes unread and fills the Record Struct(Fspec, DataItems array) in byte.
//func (rec *Record) Decode(data []byte, stdUAP _uap.StandardUAP) (unRead int, err error) {
func (rec *Record) Decode(data []byte, uap item.StandardUAP) (unRead int, err error) {
	rec.Cat = uap.Category

	rb := bytes.NewReader(data)
	rec.Fspec, err = item.FspecReader(rb)
	unRead = rb.Len()
	if err != nil {
		return unRead, err
	}

	frnIndex := item.FspecIndex(rec.Fspec)
	offset := uint8(0) // offset shifts the index for a conditional UAP
	rec.DataItems = make([]item.DataItem, 0, len(frnIndex))

	for _, frn := range frnIndex {
		uapItem := uap.DataItems[frn-1-offset] // here the index corresponds to the FRN
		var dataItem item.DataItem
		dataItem, err = item.GetItem(uapItem)
		if err != nil {
			unRead = rb.Len()
			return unRead, err
		}
		//err = Readers(dataItem, rb, uapItem)
		err = item.Readers(dataItem, rb)
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
		//rec.DataItems = append(rec.DataItems, contextType.DataItemName)
		rec.DataItems = append(rec.DataItems, dataItem)
		/*
			if uapItem.Conditional {
				switch uapItem.Type {
				case _uap.Fixed:
					uap.DataItems = selectUAPConditional(uap.Category, dataItem.Payload())
					//uap.DataItems = selectUAPConditional(uap.Category, contextType.DataItemName.Payload())
				case _uap.Extended:
					uap.DataItems = selectUAPConditional(uap.Category, dataItem.Payload())
					//uap.DataItems = selectUAPConditional(uap.Category, contextType.DataItemName.Payload())
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

	for _, dataItem := range rec.DataItems {
		tmp := dataItem.String()
		items = append(items, tmp)
	}
	return items
}

// Payload returns a slice of byte for one asterix record.
func (rec Record) Payload() []byte {
	var pd []byte
	pd = append(pd, rec.Fspec...)
	for _, dataItem := range rec.DataItems {
		pd = append(pd, dataItem.Payload()...)
	}
	return pd
}

/*
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
*/

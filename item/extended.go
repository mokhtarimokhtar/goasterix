package item

import (
	"bytes"
	"encoding/binary"
)

// Extended length Data Fields, being of a variable length, shall contain a primary part of predetermined length,
// immediately followed by a number of secondary parts, each of predetermined length.
// The presence of the next following secondary part shall be indicated by the setting to one of the
// Least Significant Bit (LSB) of the last octet of the preceding part (either the primary part or a secondary part).
// This bit which is reserved for that purpose is called the Field Extension Indicator (FX).
type Extended struct {
	Base
	PrimaryItemSize   uint8
	SecondaryItemSize uint8
	SubItems          []SubItem
}

func (e *Extended) Clone() DataItem {
	return &Extended{
		Base:              e.Base,
		PrimaryItemSize:   e.PrimaryItemSize,
		SecondaryItemSize: e.SecondaryItemSize,
		SubItems:          e.SubItems,
	}
}

func (e Extended) GetSubItems() []SubItem {
	return e.SubItems
}

// Reader extracts data item type Extended (FX: last bit = 1).
// primarySize parameter defines the Primary Subitem of extended dataField.
// secondarySize parameter defines the Secondary Subitem of extended dataField.
func (e *Extended) Reader(rb *bytes.Reader) error {
	var err error
	/*
		tmp := make([]byte, e.PrimaryItemSize)
		err = binary.Read(rb, binary.BigEndian, &tmp)
		if err != nil {
			return err
		}
		e.Primary = tmp

		if tmp[e.PrimaryItemSize-1]&0x01 != 0 {
			for {
				tmp := make([]byte, e.SecondaryItemSize)
				err = binary.Read(rb, binary.BigEndian, &tmp)
				if err != nil {
					return err
				}
				e.Secondary = append(e.Secondary, tmp...)
				if tmp[e.SecondaryItemSize-1]&0x01 == 0 {
					break
				}
			}
		}
	*/

	tmp := make([]byte, e.PrimaryItemSize)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		e.SubItems = nil
		return err
	}
	offset := 0
	tmpSubItems := e.SubItems

	e.SubItems = make([]SubItem, 0, len(e.SubItems))
	for i, subItem := range tmpSubItems {
		if subItem.Bit == 1 {
			offset = i + 1
			break
		}
		sub := subItem.Clone()
		err = sub.Reader(tmp)
		if err != nil {
			return err
		}
		e.SubItems = append(e.SubItems, *sub)
	}

	if tmp[e.PrimaryItemSize-1]&0x01 != 0 {
		for {
			tmpSec := make([]byte, e.SecondaryItemSize)
			err = binary.Read(rb, binary.BigEndian, &tmpSec)
			if err != nil {
				return err
			}

			for i, subItem := range tmpSubItems[offset:] {
				if subItem.Bit == 1 {
					offset = offset + i + 1
					break
				}
				sub := subItem.Clone()
				err = sub.Reader(tmpSec)
				if err != nil {
					return err
				}
				e.SubItems = append(e.SubItems, *sub)
			}

			if tmpSec[e.SecondaryItemSize-1]&0x01 == 0 {
				break
			}
		}
	}
	return err
}

/*
// Payload returns this dataField as bytes.
func (e Extended) Payload() []byte {
	var p []byte
	p = append(p, e.Primary...)
	p = append(p, e.Secondary...)
	return p
}
*/

// String implements fmt.Stringer in hexadecimal
func (e Extended) String() string {
	var buf bytes.Buffer
	buf.Reset()
	buf.WriteString(e.Base.DataItemName)
	buf.WriteByte(':')

	for _, subItem := range e.SubItems {
		buf.WriteByte('[')
		buf.WriteString(subItem.String())
		buf.WriteByte(']')
	}

	return buf.String()
}

package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

// Extended length Data Fields, being of a variable length, shall contain a primary part of predetermined length,
// immediately followed by a number of secondary parts, each of predetermined length.
// The presence of the next following secondary part shall be indicated by the setting to one of the
// Least Significant Bit (LSB) of the last octet of the preceding part (either the primary part or a secondary part).
// This bit which is reserved for that purpose is called the Field Extension Indicator (FX).
type Extended struct {
	MetaItem
	Primary   []byte
	Secondary []byte
}
// Reader extracts data item type Extended (FX: last bit = 1).
// primarySize parameter defines the Primary Subitem of extended field.
// secondarySize parameter defines the Secondary Subitem of extended field.
func (e *Extended) Reader(rb *bytes.Reader, field uap.DataField) error {
	var err error
	e.MetaItem.NewMetaItem(field)
	primarySize := field.Extended.PrimarySize
	secondarySize := field.Extended.SecondarySize
	tmp := make([]byte, primarySize)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return err
	}
	e.Primary = tmp

	if tmp[primarySize-1]&0x01 != 0 {
		for {
			tmp := make([]byte, secondarySize)
			err = binary.Read(rb, binary.BigEndian, &tmp)
			if err != nil {
				return err
			}
			e.Secondary = append(e.Secondary, tmp...)
			if tmp[secondarySize-1]&0x01 == 0 {
				break
			}
		}
	}

	return err
}

func (e Extended) Payload() []byte {
	var p []byte
	p = append(p, e.Primary...)
	p = append(p, e.Secondary...)
	return p
}

func (e Extended) String() string {
	return e.MetaItem.DataItem + ": " + hex.EncodeToString(e.Primary) + hex.EncodeToString(e.Secondary)
}

func (e Extended) Frn() uint8 {
	return e.MetaItem.FRN
}

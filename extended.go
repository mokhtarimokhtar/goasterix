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
	Base
	PrimaryItemSize   uint8
	SecondaryItemSize uint8
	Primary           []byte
	Secondary         []byte
}

func NewExtended(field uap.IDataField) Item {
	f := &Extended{}
	f.Base.NewBase(field)
	//f.PrimaryItemSize = field.Extended.PrimarySize
	f.PrimaryItemSize = field.GetSize().ForExtendedPrimary
	//f.SecondaryItemSize = field.Extended.SecondarySize
	f.SecondaryItemSize = field.GetSize().ForExtendedSecondary
	return f
}

// Reader extracts data item type Extended (FX: last bit = 1).
// primarySize parameter defines the Primary Subitem of extended dataField.
// secondarySize parameter defines the Secondary Subitem of extended dataField.
func (e *Extended) Reader(rb *bytes.Reader) error {
	var err error
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

	return err
}

// Payload returns this dataField as bytes.
func (e Extended) Payload() []byte {
	var p []byte
	p = append(p, e.Primary...)
	p = append(p, e.Secondary...)
	return p
}

// String implements fmt.Stringer in hexadecimal
func (e Extended) String() string {
	var buf bytes.Buffer
	buf.Reset()
	buf.WriteString(e.Base.DataItem)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(e.Primary))
	buf.WriteString(hex.EncodeToString(e.Secondary))
	return buf.String()
}
